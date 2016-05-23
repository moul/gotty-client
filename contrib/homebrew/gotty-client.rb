require "language/go"

class GottyClient < Formula
  desc "gotty-client: GoTTY client for your terminal"
  homepage "https://github.com/moul/gotty-client"
  url "https://github.com/moul/gotty-client/archive/v1.6.0.tar.gz"
  sha256 "5f8b67fd47c1586a06aead9a89706d49c52cd4aaaad9181e38722d4be6a78d43"

  head "https://github.com/moul/gotty-client.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["GOBIN"] = buildpath
    ENV["GO15VENDOREXPERIMENT"] = "1"
    (buildpath/"src/github.com/moul/gotty-client").install Dir["*"]

    system "go", "build", "-o", "#{bin}/gotty-client", "github.com/moul/gotty-client/cmd/gotty-client/"

    bash_completion.install "#{buildpath}/src/github.com/moul/gotty-client/contrib/completion/bash_autocomplete"
    zsh_completion.install "#{buildpath}/src/github.com/moul/gotty-client/contrib/completion/zsh_autocomplete"
  end

  test do
    output = shell_output(bin/"gotty-client --version")
    assert output.include? "gotty-client version"
  end
end
