package scrape

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"regexp"

	"github.com/tbuckley/goscrape"
	"github.com/tbuckley/goscrape/handlers"
)

type HandlerConfiguration struct {
	Pattern  string `json:"pattern"`
	Handler  string `json:"handler"`
	Priority int    `json:"priority"`
}

type Configuration struct {
	InitialURLs []string               `json:"initial"`
	Handlers    []HandlerConfiguration `json:"handlers"`
}

func getNamedHandler(name string) goscrape.Handler {
	switch name {
	case "recipe":
		return HandleRecipe
	case "default":
		return handlers.Default
	case "null":
		return handlers.Null
	default:
		return nil
	}
}

func (c *Configuration) InitializeScraper(scraper goscrape.WebScraper) error {
	var err error

	err = c.AddHandlers(scraper)
	if err != nil {
		return err
	}
	err = c.AddURLs(scraper)
	if err != nil {
		return err
	}
	return nil
}

func (c *Configuration) AddURLs(scraper goscrape.WebScraper) error {
	for _, rawurl := range c.InitialURLs {
		parsedurl, err := url.Parse(rawurl)
		if err != nil {
			return err
		}
		scraper.Enqueue(parsedurl)
	}
	return nil
}

func (c *Configuration) AddHandlers(scraper goscrape.WebScraper) error {
	for _, h := range c.Handlers {

		pattern, err := regexp.CompilePOSIX(h.Pattern)
		if err != nil {
			return err
		}
		handler := getNamedHandler(h.Handler)
		scraper.AddHandlerPriority(pattern, handler, uint(h.Priority))
	}
	return nil
}

func ParseConfiguration(filename string) (*Configuration, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := new(Configuration)
	json.Unmarshal(contents, config)
	return config, nil
}
