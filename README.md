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
3. ./bin/recipes
4. Monitor progress at localhost:8080

### Issues
* May crash instead of completing gracefully (haven't run it to end...)
* Requires Mongo running on local machine
* Uses hard-coded URLs instead of config file
* Probably many more...
