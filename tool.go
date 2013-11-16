package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tbuckley/recipes/sample"
	"github.com/tbuckley/recipes/scrape"
)

var (
	expvarServer = flag.Bool("expvar", false, "")
	serverPort   = flag.String("port", ":8080", "")
)

const usage = `Usage:
  recipes [options] [--expvar] [--port=PORT] scrape
  recipes [options] [--size=N] [--path=PATH] sample
`

// Scrape will scrape the appropriate websites
func Scrape() {
	scrape.Run()
}

func PrintUsage() {
	fmt.Println(usage)
}

func StartServer() {
	go func() {
		log.Fatal(http.ListenAndServe(*serverPort, nil))
	}()
}

func init() {
	log.SetPrefix("[LOG] ")
	log.SetOutput(os.Stdout)
}

func main() {
	flag.Parse()

	if *expvarServer {
		StartServer()
	}

	if flag.NArg() < 1 {
		PrintUsage()
		return
	}
	command := flag.Arg(0)

	switch command {
	case "scrape":
		Scrape()
	case "sample":
		sample.Run()
	default:
		PrintUsage()
	}
}
