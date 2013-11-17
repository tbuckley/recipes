package scrape

import (
	"expvar"
	"log"
	"net/url"

	"github.com/tbuckley/goscrape"
	"github.com/tbuckley/goscrape/handlers"
	"github.com/tbuckley/htmlutils"
	"github.com/tbuckley/recipes/db"
)

var varRecipes = expvar.NewInt("recipes")

func HandleRecipe(s goscrape.WebScraper, page *url.URL) {
	log.Printf("RECIPE: %s\n", page.String())
	varRecipes.Add(1)

	body, err := htmlutils.FetchPage(page)
	if err != nil {
		log.Printf("Error w/ page fetch: %s", err)
		return
	}

	r := db.Recipe{
		Page:    page.String(),
		Content: string(body),
	}
	err = r.Save()
	if err != nil {
		panic(err)
	}

	handlers.DefaultPage(s, page, body)
}

func Run(arguments map[string]interface{}) {
	scraper := goscrape.NewScraper()

	config, err := ParseConfiguration(arguments["CONFIG"].(string))
	if err != nil {
		panic(err)
	}

	err = config.InitializeScraper(scraper)
	if err != nil {
		panic(err)
	}

	scraper.Start()
}
