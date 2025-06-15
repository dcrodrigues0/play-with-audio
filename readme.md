# Play with Audio (Go + FFmpeg)
A small Golang project to experiment with audio recording using FFmpeg and ALSA devices on Linux.

This project allows you to:
* Start and stop audio recording via CLI.
* Choose your input audio device (microphone).
* Cancel the recording gracefully

# Requirements
✅ Linux (tested on Fedora)
✅ FFmpeg installed and available in your PATH
✅ ALSA audio devices (run arecord -l to check available devices)
✅ Go 1.20 or higher

# How to run
`go run ./cmd/main.go`
