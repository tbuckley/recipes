package main

import (
	"github.com/tbuckley/recipes/db"
	"log"
	"net/http"
	"os"

	"github.com/docopt/docopt.go"
	"github.com/tbuckley/recipes/empty"
	"github.com/tbuckley/recipes/sample"
	"github.com/tbuckley/recipes/scrape"
)

const usage = `Recipes.

Usage:
  recipes [options] sample SIZE OUTPUT
  recipes [options] scrape CONFIG
  recipes [options] empty

Options:
  --mongodb=DB    Mongo DB to use [default: recipes]
  --mongourl=URL  Mongo server to connect to [default: 127.0.0.1]
  --expvar        Start expvar server if no server running
  --addr=ADDR     Address on which to run server [default: :8080]
`

func StartServer(addr string) {
	go func() {
		log.Fatal(http.ListenAndServe(addr, nil))
	}()
}

func init() {
	log.SetPrefix("[LOG] ")
	log.SetOutput(os.Stdout)
}

func main() {
	arguments, _ := docopt.Parse(usage, nil, true, "Recipes 0.1", false)

	err := db.Connect(arguments["--mongourl"].(string), arguments["--mongodb"].(string))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if arguments["--expvar"].(bool) {
		StartServer(arguments["--addr"].(string))
	}

	if arguments["scrape"].(bool) {
		scrape.Run(arguments)
	} else if arguments["sample"].(bool) {
		sample.Run(arguments)
	} else if arguments["empty"].(bool) {
		empty.Run(arguments)
	}
}
