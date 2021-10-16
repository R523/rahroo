package main

import (
	"github.com/pterm/pterm"
	"golang.org/x/exp/io/spi"
)

// Channel indicates MCP3008 input channel.
const Channel = 0

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Ra", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("hroo", pterm.NewStyle(pterm.FgLightRed)),
	).Render(); err != nil {
		_ = err
	}

	dev, err := spi.Open(&spi.Devfs{
		Dev:      "/dev/spidev0.0",
		Mode:     spi.Mode0,
		MaxSpeed: 3600000,
	})
	if err != nil {
		pterm.Fatal.Printf("cannot open spi device 0.0 %s\n", err)
	}
	defer dev.Close()

	txbuf := []byte{0x01, (8 + Channel) << 4, 0x00}
	rxbuf := make([]byte, 3)

	if err := dev.Tx(txbuf, rxbuf); err != nil {
		pterm.Fatal.Printf("cannot read from spi %s\n", err)
	}

	value := uint16(rxbuf[1]&0x03)<<8 | uint16(rxbuf[2])

	pterm.Success.Printf("%d\n", value)
}
