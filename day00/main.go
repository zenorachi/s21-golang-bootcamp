package main

import (
	"flag"
	"fmt"
	"stats/statistics"
)

type Mode struct {
	mean   bool
	median bool
	mode   bool
	sd     bool
}

func main() {
	var sm statistics.StatMetrics
	var m Mode
	m.initialize()
	sm.Init()
	if m.mean {
		fmt.Printf("Mean: %.2f\n", sm.Mean())
	}
	if m.median {
		fmt.Printf("Median: %.2f\n", sm.Median())
	}
	if m.mode {
		fmt.Printf("Mode: %d\n", sm.Mode())
	}
	if m.sd {
		fmt.Printf("SD: %.2f\n", sm.StandardDeviation())
	}
	if !m.mean && !m.median && !m.mode && !m.sd {
		fmt.Printf("Mean: %.2f\n", sm.Mean())
		fmt.Printf("Median: %.2f\n", sm.Median())
		fmt.Printf("Mode: %d\n", sm.Mode())
		fmt.Printf("SD: %.2f\n", sm.StandardDeviation())
	}
}

func (m *Mode) initialize() {
	flag.BoolVar(&m.mean, "mean", false, "mean output")
	flag.BoolVar(&m.median, "median", false, "median output")
	flag.BoolVar(&m.mode, "mode", false, "mode output")
	flag.BoolVar(&m.sd, "sd", false, "standard deviation output")
	flag.Parse()
}
