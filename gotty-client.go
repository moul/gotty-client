package gottyclient

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"unicode/utf8"
	"unsafe"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh/terminal"
)

// GetWebsocketURL transforms a GoTTY http URL to its WebSocket URL
func GetWebsocketURL(httpURL string) (*url.URL, error) {
	target, err := url.Parse(httpURL)
	if err != nil {
		return nil, err
	}

	if target.Scheme == "https" {
		target.Scheme = "wss"
	} else {
		target.Scheme = "ws"
	}

	target.Path = strings.TrimLeft(target.Path+"ws", "/")

	return target, nil
}

type Client struct {
	Dialer    *websocket.Dialer
	Conn      *websocket.Conn
	Headers   http.Header
	Target    string
	Connected bool
}

// Connect tries to dial a websocket server
func (c *Client) Connect() error {
	conn, _, err := c.Dialer.Dial(c.Target, c.Headers)
	if err != nil {
		return err
	}
	c.Conn = conn
	return nil
}

// Close will nicely close the dialer
func (c *Client) Close() {
	c.Conn.Close()
}

// Loop will look indefinitely for new messages
func (c *Client) Loop() error {
	if !c.Connected {
		err := c.Connect()
		if err != nil {
			return err
		}
	}

	done := make(chan bool)
	go c.readLoop(done)
	go c.writeLoop(done)
	go c.termsizeLoop(done)
	<-done
	return nil
}

type winsize struct {
	Rows    uint16 `json:"rows"`
	Columns uint16 `json:"columns"`
	// unused
	x uint16
	y uint16
}

func (c *Client) termsizeLoop(done chan bool) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	ws := winsize{}

	for {
		syscall.Syscall(syscall.SYS_IOCTL,
			uintptr(0), uintptr(syscall.TIOCGWINSZ),
			uintptr(unsafe.Pointer(&ws)))

		b, err := json.Marshal(ws)
		if err != nil {
			logrus.Warnf("json.Marshal error: %v", err)
		}

		err = c.Conn.WriteMessage(websocket.TextMessage, append([]byte("1"), b...))
		if err != nil {
			logrus.Warnf("ws.WriteMessage failed: %v", err)
		}

		<-ch
	}
}

func (c *Client) writeLoop(done chan bool) {
	oldState, err := terminal.MakeRaw(0)
	if err == nil {
		defer terminal.Restore(0, oldState)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		x, size, err := reader.ReadRune()
		if size <= 0 || err != nil {
			done <- true
			return
		}

		p := make([]byte, size)
		utf8.EncodeRune(p, x)

		err = c.Conn.WriteMessage(websocket.TextMessage, append([]byte("0"), p...))
		if err != nil {
			done <- true
			return
		}
	}
}

func (c *Client) readLoop(done chan bool) {
	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			done <- true
			logrus.Warnf("c.Conn.ReadMessage: %v", err)
			return
		}

		switch data[0] {
		case '0': // data
			fmt.Print(string(data[1:]))
		case '1': // new title
			newTitle := string(data[1:])
			fmt.Printf("\033]0;%s\007", newTitle)
		case '2': // json prefs
			logrus.Warnf("Unhandled protocol message: json pref: %s", string(data))
		default:
			logrus.Warnf("Unhandled protocol message: %s", string(data))
		}
	}
}

// NewClient returns a GoTTY client object
func NewClient(httpURL string) (*Client, error) {
	target, err := GetWebsocketURL(httpURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		Dialer: &websocket.Dialer{},
		Target: target.String(),
	}, nil
}
