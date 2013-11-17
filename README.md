# Recipes

Recipes is currently just a scraper for recipe sites written in Go, but it hopes to one day be oh-so-much more.

### Requirements
* go get github.com/tbuckley/goscrape
* go get github.com/tbuckley/htmlutils
* go get labix.org/v2/mgo
* go get github.com/docopt/docopt.go

### Steps

1. Start a Mongo server on your local machine.
2. go install github.com/tbuckley/recipes
3. ./bin/recipes scrape path/to/config

## Configuration

Configuration is stored as a JSON file. Here is a sample:

```json
{
  "handlers": [
    {
      "pattern": "^http://allrecipes\\.com/recipe/.*?/detail\\.aspx.*",
      "handler": "recipe",
      "priority": 31
    },
    {
      "pattern": "^http://allrecipes\\.com.*",
      "handler": "default",
      "priority": 15
    }
  ],
  "initial": ["http://allrecipes.com"]
}
```

### Issues
* May crash instead of completing gracefully (haven't run it to end...)
* Probably many more...
