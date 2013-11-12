package main

import (
	"expvar"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/tbuckley/goscrape"
	"github.com/tbuckley/goscrape/handlers"
	"labix.org/v2/mgo"
)

var (
	varRecipes = expvar.NewInt("recipes")

	recipes *mgo.Collection
)

type (
	Recipe struct {
		Page string
	}
)

func HandleRecipe(s goscrape.WebScraper, page *url.URL) {
	log.Printf("RECIPE: %s\n", page.String())
	varRecipes.Add(1)
	err := recipes.Insert(&Recipe{Page: page.String()})
	if err != nil {
		panic(err)
	}
	handlers.Default(s, page)
}

func AddPattern(s goscrape.WebScraper, pattern string, handler goscrape.Handler) {
	p, err := regexp.CompilePOSIX(pattern)
	if err != nil {
		log.Fatal(err)
	}
	s.AddHandler(p, handler)
}
func AddPatternPriority(s goscrape.WebScraper, pattern string, handler goscrape.Handler, priority int) {
	p, err := regexp.CompilePOSIX(pattern)
	if err != nil {
		log.Fatal(err)
	}
	s.AddHandlerPriority(p, handler, priority)
}

func init() {
	log.SetPrefix("[LOG] ")
	log.SetOutput(os.Stdout)
}

func UrlOrDie(page string) *url.URL {
	pageURL, err := url.Parse(page)
	if err != nil {
		log.Fatal("Could not parse:", page)
	}
	return pageURL
}

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	recipes = session.DB("recipes").C("scraper")

	scraper := goscrape.NewScraper()
	AddPatternPriority(scraper, "^http://allrecipes\\.com/recipe/.*?/detail\\.aspx.*", HandleRecipe, goscrape.HighPriority)
	AddPattern(scraper, "^http://allrecipes\\.com/menu/.*", handlers.Null)
	AddPattern(scraper, "^http://allrecipes\\.com/video.*", handlers.Null)
	AddPattern(scraper, "^http://allrecipes\\.com/my/.*", handlers.Null)
	AddPattern(scraper, "^http://allrecipes\\.com.*/membership/.*", handlers.Null)
	AddPattern(scraper, "^http://allrecipes\\.com.*", handlers.Default)

	scraper.Enqueue(UrlOrDie("http://allrecipes.com"))
	scraper.Start()
}
