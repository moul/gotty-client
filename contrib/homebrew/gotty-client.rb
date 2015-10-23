require "language/go"

class GottyClient < Formula
  desc "gotty-client: GoTTY client for your terminal"
  homepage "https://github.com/moul/gotty-client"
  url "https://github.com/moul/gotty-client/archive/v1.2.0.tar.gz"
  sha256 "cddd14a551d76ae3cf79ccd5ee4beb39178cc3641a0ddbe69bc27e2a686a0cb5"

  head "https://github.com/moul/gotty-client.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["CGO_ENABLED"] = "0"
    ENV.prepend_create_path "PATH", buildpath/"bin"

    mkdir_p buildpath/"src/github.com/moul"
    ln_s buildpath, buildpath/"src/github.com/moul/gotty-client"
    Language::Go.stage_deps resources, buildpath/"src"

    # FIXME: update version variable
    system "go", "build", "-o", "gotty-client", "./cmd/gotty-client/"
    bin.install "gotty-client"

    # FIXME: add autocompletion
  end

  test do
    output = shell_output(bin/"gotty-client --version")
    assert output.include? "gotty-client version"
  end
end
