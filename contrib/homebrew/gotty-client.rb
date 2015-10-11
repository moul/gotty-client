require "language/go"

class GottyClient < Formula
  desc "gotty-client: GoTTY client for your terminal"
  homepage "https://github.com/moul/gotty-client"
  url "https://github.com/moul/gotty-client/archive/v1.1.0.tar.gz"
  sha256 "f080e685d6eaff9dfe627d8a7b5b485e2d023fa1ec0b1db7db37775177008708"

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
