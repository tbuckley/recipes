package scrape

import (
	"expvar"
	"log"
	"net/url"
	"regexp"

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

func AddPattern(s goscrape.WebScraper, pattern string, handler goscrape.Handler) {
	p, err := regexp.CompilePOSIX(pattern)
	if err != nil {
		log.Fatal(err)
	}
	s.AddHandler(p, handler)
}
func AddPatternPriority(s goscrape.WebScraper, pattern string, handler goscrape.Handler, priority uint) {
	p, err := regexp.CompilePOSIX(pattern)
	if err != nil {
		log.Fatal(err)
	}
	s.AddHandlerPriority(p, handler, priority)
}

func UrlOrDie(page string) *url.URL {
	pageURL, err := url.Parse(page)
	if err != nil {
		log.Fatal("Could not parse:", page)
	}
	return pageURL
}

func Run() {
	// Start MongoDB
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	scraper := goscrape.NewScraper()

	// Set scraper rules
	AddPatternPriority(scraper, "^http://allrecipes\\.com/recipe/.*?/detail\\.aspx.*", HandleRecipe, goscrape.HighPriority)
	AddPattern(scraper, "^http://allrecipes\\.com/menu/.*", handlers.Null)
	AddPattern(scraper, "^http://allrecipes\\.com/video.*", handlers.Null)
	AddPattern(scraper, "^http://allrecipes\\.com/my/.*", handlers.Null)
	AddPattern(scraper, "^http://allrecipes\\.com.*/membership/.*", handlers.Null)
	AddPattern(scraper, "^http://allrecipes\\.com.*", handlers.Default)

	scraper.Enqueue(UrlOrDie("http://allrecipes.com"))
	scraper.Start()
}
