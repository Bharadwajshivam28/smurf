class Smurf < Formula
    desc "CloudNative CI/CD Management Tool"
    homepage "https://github.com/clouddrove/smurf"
    license "Apache-2.0"
    version "v1"
  
    on_macos do
  
      if Hardware::CPU.intel?
        url "https://github.com/clouddrove-sandbox/smurf-custon-github-action-test/releases/download/v1/smurf-darwin-amd64.zip"
        sha256 ""
      end
  
      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/clouddrove-sandbox/smurf-custon-github-action-test/releases/download/v1/smurf-darwin-arm64.zip"
        sha256 ""
      end
    end
  
    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/clouddrove-sandbox/smurf-custon-github-action-test/releases/download/v1/smurf-linux-amd64.zip"
        sha256 ""
      end
      if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
        url "https://github.com/clouddrove-sandbox/smurf-custon-github-action-test/releases/download/v1/smurf-linux-arm64.zip"
        sha256 ""
      end
    end
  
    def install
      bin.install "smurf"
    end
  end