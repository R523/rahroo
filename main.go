package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pterm/pterm"
	"github.com/stianeikeland/go-rpio/v4"
)

const (
	// Channel indicates MCP3008 input channel.
	Channel uint8 = 0
	// Command sends over spi interface
	// nolint: gomnd
	Command uint8 = (8 + Channel) << 4

	Threshold = 20

	LEDPin rpio.Pin = 22
)

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Rah", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("roo", pterm.NewStyle(pterm.FgLightRed)),
	).Render(); err != nil {
		_ = err
	}

	if err := rpio.Open(); err != nil {
		pterm.Error.Printf("falied to open rpio %s", err)

		return
	}
	defer rpio.Close()

	rpio.PinMode(LEDPin, rpio.Output)

	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		pterm.Error.Printf("falied to open spi0 %s", err)

		return
	}
	defer rpio.SpiEnd(rpio.Spi0)

	rpio.SpiChipSelect(0)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			data := []byte{0x01, Command, 0x00}

			rpio.SpiExchange(data)

			// nolint: gomnd
			value := uint16(data[1]&0x03)<<8 | uint16(data[2])

			pterm.Info.Println(value)

			if value < Threshold {
				rpio.WritePin(LEDPin, rpio.High)
			} else {
				rpio.WritePin(LEDPin, rpio.Low)
			}
		}
	}
}
