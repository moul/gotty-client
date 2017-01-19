require "language/go"

class GottyClient < Formula
  desc "gotty-client: GoTTY client for your terminal"
  homepage "https://github.com/moul/gotty-client"
  url "https://github.com/moul/gotty-client/archive/v1.6.1.tar.gz"
  sha256 "322aba97334073eaac833f1b40ba83c7861ba32247d3160504cc78f64b5a2ef6"

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
