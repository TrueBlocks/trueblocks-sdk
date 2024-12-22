// Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were auto generated. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

package sdk

import (
	// EXISTING_CODE
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	// EXISTING_CODE
)

type DaemonOptions struct {
	Url       string            `json:"url,omitempty"`
	Api       DaemonApi         `json:"api,omitempty"`
	Scrape    DaemonScrape      `json:"scrape,omitempty"`
	Monitor   bool              `json:"monitor,omitempty"`
	Silent    bool              `json:"silent,omitempty"`
	RenderCtx *output.RenderCtx `json:"-"`
	Globals
}

// String implements the stringer interface
func (opts DaemonOptions) String() string {
	bytes, _ := json.Marshal(opts)
	return string(bytes)
}

type DaemonApi int

const (
	NoDA  DaemonApi = 0
	DAOff           = 1 << iota
	DAOn
)

func (v DaemonApi) String() string {
	switch v {
	case NoDA:
		return "none"
	}

	var m = map[DaemonApi]string{
		DAOff: "off",
		DAOn:  "on",
	}

	var ret []string
	for _, val := range []DaemonApi{DAOff, DAOn} {
		if v&val != 0 {
			ret = append(ret, m[val])
		}
	}

	return strings.Join(ret, ",")
}

func enumFromDaemonApi(values []string) (DaemonApi, error) {
	if len(values) == 0 {
		return NoDA, fmt.Errorf("no value provided for api option")
	}

	var result DaemonApi
	for _, val := range values {
		switch val {
		case "off":
			result |= DAOff
		case "on":
			result |= DAOn
		default:
			return NoDA, fmt.Errorf("unknown api: %s", val)
		}
	}

	return result, nil
}

type DaemonScrape int

const (
	NoDS  DaemonScrape = 0
	DSOff              = 1 << iota
	DSBlooms
	DSIndex
)

func (v DaemonScrape) String() string {
	switch v {
	case NoDS:
		return "none"
	}

	var m = map[DaemonScrape]string{
		DSOff:    "off",
		DSBlooms: "blooms",
		DSIndex:  "index",
	}

	var ret []string
	for _, val := range []DaemonScrape{DSOff, DSBlooms, DSIndex} {
		if v&val != 0 {
			ret = append(ret, m[val])
		}
	}

	return strings.Join(ret, ",")
}

func enumFromDaemonScrape(values []string) (DaemonScrape, error) {
	if len(values) == 0 {
		return NoDS, fmt.Errorf("no value provided for scrape option")
	}

	var result DaemonScrape
	for _, val := range values {
		switch val {
		case "off":
			result |= DSOff
		case "blooms":
			result |= DSBlooms
		case "index":
			result |= DSIndex
		default:
			return NoDS, fmt.Errorf("unknown scrape: %s", val)
		}
	}

	return result, nil
}

// EXISTING_CODE
func StartApiServer(urlOut *string) {
	go func() {
		*urlOut = getApiUrl()
		ready := make(chan bool)

		opts := DaemonOptions{
			Silent: true,
			Url:    *urlOut,
		}

		// Start the daemon process
		go func() {
			in := opts.toInternal()
			buffer := bytes.Buffer{}
			if err := in.DaemonBytes(&buffer); err != nil {
				fmt.Fprintf(os.Stderr, "Error starting daemon: %s\n", err)
				return
			}
			ready <- true
		}()

		<-ready
		fmt.Printf("API server started at %s\n", *urlOut)

		// Handle signals for graceful shutdown
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt) // Portable `os.Interrupt` for Ctrl+C

		cleanupDone := make(chan bool, 1) // Tracks cleanup completion
		firstSignal := true               // Tracks whether it's the first signal

		for {
			sig := <-sigChan
			if firstSignal {
				fmt.Printf("\nReceived signal: %s. Initiating cleanup...\n", sig)
				firstSignal = false

				// Start cleanup in a separate goroutine
				go func() {
					// Perform cleanup tasks here
					time.Sleep(3 * time.Second) // Simulate cleanup work
					fmt.Println("Cleanup completed.")
					cleanupDone <- true
				}()

				// Wait for cleanup to complete or proceed with shutdown
				<-cleanupDone
				fmt.Println("Exiting gracefully after cleanup.")
				os.Exit(0)
			} else {
				fmt.Printf("\nReceived signal: %s. Forcing shutdown...\n", sig)
				os.Exit(1) // Force quit
			}
		}
	}()
}

// getApiUrl returns the URL (including port) where the API server is running (or will run).
func getApiUrl() string {
	apiPort := strings.ReplaceAll(os.Getenv("TB_API_PORT"), ":", "")
	if apiPort == "" {
		preferred := []string{"8080", "8088", "9090", "9099"}
		apiPort = findAvailablePort(preferred)
	}

	return "localhost:" + apiPort
}

// findAvailablePort returns a port number that is available for listening.
func findAvailablePort(preferred []string) string {
	for _, port := range preferred {
		if listener, err := net.Listen("tcp", port); err == nil {
			defer listener.Close()
			return port
		}
	}

	if listener, err := net.Listen("tcp", ":0"); err == nil {
		defer listener.Close()
		addr := listener.Addr().(*net.TCPAddr)
		return fmt.Sprintf("%d", addr.Port)
	}

	return "0"
}

// EXISTING_CODE
