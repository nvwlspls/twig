class Twig < Formula
  desc "Minimal static site generator written in Go"
  homepage "https://github.com/yourusername/twig"
  version "1.0.0"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/yourusername/twig/releases/download/v1.0.0/twig-1.0.0-darwin-arm64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_ARM64"
    else
      url "https://github.com/yourusername/twig/releases/download/v1.0.0/twig-1.0.0-darwin-amd64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_AMD64"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/yourusername/twig/releases/download/v1.0.0/twig-1.0.0-linux-arm64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_LINUX_ARM64"
    else
      url "https://github.com/yourusername/twig/releases/download/v1.0.0/twig-1.0.0-linux-amd64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_LINUX_AMD64"
    end
  end

  def install
    bin.install "twig"
  end

  test do
    system "#{bin}/twig", "--version"
  end
end 