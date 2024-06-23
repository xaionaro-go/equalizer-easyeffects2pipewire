package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/xaionaro-go/equalizer-easyeffects2pipewire/pkg/easyeffects"
	"github.com/xaionaro-go/equalizer-easyeffects2pipewire/pkg/pipewire"
)

func assertNoError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Fprintf(os.Stderr, "syntax: %s <easyeffects-equalizer-preset> <pipewire-filter-config-output>\n", os.Args[0])
		os.Exit(1)
	}

	eeFilePath := flag.Arg(0)
	pwFilePath := flag.Arg(1)

	eeFile, err := os.Open(eeFilePath)
	assertNoError(err)

	preset, err := easyeffects.ParseEqualizerPreset(eeFile)
	eeFile.Close()
	assertNoError(err)

	pwFile, err := os.OpenFile(pwFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	assertNoError(err)

	err = pipewire.WriteEqualizerPreset(pwFile, preset)
	pwFile.Close()
	assertNoError(err)
}
