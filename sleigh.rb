require "formula"

class Sleigh < Formula
  homepage "https://github.com/hackm/sleigh"
  url "https://github.com/hackm/sleigh/releases/download/v0.0.1/sleigh_v0.0.1_darwin_amd64.tar.gz"
  sha256 ""
  head "https://github.com/hackm/sleigh.git"
  version "0.0.1"

  def install
    bin.install "sleigh"
  end
end
