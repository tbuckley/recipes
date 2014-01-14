package sample

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"

	"github.com/tbuckley/recipes/db"
)

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Run(arguments map[string]interface{}) {
	outputPath := arguments["OUTPUT"].(string)
	finfo, err := os.Lstat(outputPath)
	if err != nil {
		panic(err)
	}
	if !finfo.IsDir() {
		fmt.Printf("Invalid path: %s\n", outputPath)
	}

	size, err := strconv.ParseInt(arguments["SIZE"].(string), 10, 64)
	if err != nil {
		panic(err)
	}

	iter := db.SomeRecipes(200, 0.25)
	recipe := db.Recipe{}
	for i := 0; i < int(size) && iter.Next(&recipe); i++ {
		filename := path.Join(outputPath, fmt.Sprintf("recipe_%04d.html", i))
		err = ioutil.WriteFile(filename, []byte(recipe.Content), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Close(); err != nil {
		panic(err)
	}
}
