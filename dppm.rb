# Homebrew Formula for DPPM (Dropbox Project Manager)
class Dppm < Formula
  desc "AI-first CLI tool for project, phase, and task management using Dropbox"
  homepage "https://github.com/AI-S-Tools/dppm"
  version "1.0.0"

  if OS.mac?
    if Hardware::CPU.arm?
      url "https://github.com/AI-S-Tools/dppm/releases/download/v1.0.0/dppm-macos-arm64"
      sha256 "TBD" # This will need to be updated with actual SHA256
    else
      url "https://github.com/AI-S-Tools/dppm/releases/download/v1.0.0/dppm-macos-amd64"
      sha256 "TBD" # This will need to be updated with actual SHA256
    end
  elsif OS.linux?
    if Hardware::CPU.arm?
      url "https://github.com/AI-S-Tools/dppm/releases/download/v1.0.0/dppm-linux-arm64"
      sha256 "TBD" # This will need to be updated with actual SHA256
    else
      url "https://github.com/AI-S-Tools/dppm/releases/download/v1.0.0/dppm-linux-amd64"
      sha256 "TBD" # This will need to be updated with actual SHA256
    end
  end

  def install
    bin.install Dir["dppm*"].first => "dppm"
  end

  test do
    system "#{bin}/dppm", "--version"
    system "#{bin}/dppm", "--help"
  end

  def caveats
    <<~EOS
      🚀 DPPM (Dropbox Project Manager) - AI-First Project Management

      Quick Start:
        dppm                         # Smart startup guide
        dppm wiki "getting started"   # Complete quick start guide
        dppm wiki list               # See all available topics

      AI-Optimized Features:
        - Built-in comprehensive wiki system
        - Self-documenting commands
        - Structured, verbose output
        - AI collaboration system with DSL markers

      Storage Location: ~/Dropbox/project-management/

      Learn More:
        Repository: https://github.com/AI-S-Tools/dppm
        Use the built-in wiki system for complete documentation
    EOS
  end
end