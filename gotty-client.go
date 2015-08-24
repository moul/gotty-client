package gottyclient

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
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

	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			return err
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
	return nil
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
