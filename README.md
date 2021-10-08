# Project currently in alpha state, public API not guaranteed to stay the same across minor/patch versions

# Youtube-DL Wrapper

[![Drone Build Status](https://cloud.drone.io/api/badges/SomethingBot/youtubedlwrapper/status.svg)](https://cloud.drone.io/SomethingBot/youtubedlwrapper)
[![Go Reference](https://pkg.go.dev/badge/github.com/SomethingBot/youtubedlwrapper.svg)](https://pkg.go.dev/github.com/SomethingBot/youtubedlwrapper)
[![Go Report Card](https://goreportcard.com/badge/github.com/SomethingBot/youtubedlwrapper)](https://goreportcard.com/report/github.com/SomethingBot/youtubedlwrapper)

Type-safe Go (golang) wrapper for youtube-dl, to access youtube-dl without having to worry about directly interacting
with youtube-dl using os/exec yourself

## Installation

### youtube-dl

Note this project uses youtube-dl using exec, so it needs to know how to access and where youtube-dl is. as there are
many ways to install youtube-dl, look up instructions for your specific distro/operating system

### youtubedlwrapper

Run

```go get github.com/SomethingBot/youtubedlwrapper```

## Code Quick Start

Simple CLI application that uses youtubedlwrapper to grab a video's Metadata

```go
package main

import (
	"flag"
	"fmt"
	"github.com/SomethingBot/youtubedlwrapper"
)

func main() {
	url := flag.String("url", "", "URL of youtube video")
	flag.Parse()
	if *url == "" {
		flag.PrintDefaults()
		return
	}
	wrapper, err := youtubedlwrapper.New(youtubedlwrapper.WrapperOptions{YoutubeDLBinary: "youtube-dl"})
	if err != nil {
		fmt.Printf("Could not setup youtubedlwrapper, error (%v).\n", err)
		return
	}
	metadata, err := wrapper.GetVideoMetadata(*url)
	if err != nil {
		fmt.Printf("Could not get video metadata, error (%v)\n", err)
		return
	}
	fmt.Printf("ID (%v), Title (%v), Views (%v), Uploader (%v)\n", metadata.ID, metadata.Title, metadata.ViewCount, metadata.Uploader)
}
```
