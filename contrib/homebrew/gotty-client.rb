require "language/go"

class GottyClient < Formula
  desc "gotty-client: GoTTY client for your terminal"
  homepage "https://github.com/moul/gotty-client"
  url "https://github.com/moul/gotty-client/archive/v1.0.1.tar.gz"
  sha256 "b74a2501218863b3853dc871f53dab1893357dfcd395be853c0f36b3d427928c"

  head "https://github.com/moul/gotty-client.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["CGO_ENABLED"] = "0"
    ENV.prepend_create_path "PATH", buildpath/"bin"

    mkdir_p buildpath/"src/github.com/moul"
    ln_s buildpath, buildpath/"src/github.com/moul/gotty-client"
    Language::Go.stage_deps resources, buildpath/"src"

    # FIXME: update version
    system "go", "build", "-o", "gotty-client", "./cmd/gotty-client/"
    bin.install "gotty-client"

    # FIXME: add autocompletion
  end

  test do
    # FIXME: add test
    #output = shell_output(bin/"gotty-client --version")
    #assert output.include? "gotty-client version XXX"
  end
end
