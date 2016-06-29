package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/timercrack/task/collector"
	"github.com/timercrack/task/g"
	"github.com/timercrack/task/http"
	"github.com/timercrack/task/index"
	"github.com/timercrack/task/proc"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	versionGit := flag.Bool("vg", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}
	if *versionGit {
		fmt.Println(g.VERSION, g.COMMIT)
		os.Exit(0)
	}

	// global config
	g.ParseConfig(*cfg)
	// proc
	proc.Start()

	// graph index
	index.Start()
	// collector
	collector.Start()

	// http
	http.Start()

	select {}
}
