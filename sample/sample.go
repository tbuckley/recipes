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

	recipes, err := db.AllRecipes()
	if err != nil {
		panic(err)
	}

	size, err := strconv.ParseInt(arguments["SIZE"].(string), 10, 64)
	if err != nil {
		panic(err)
	}
	num := MinInt(int(size), len(recipes))
	for i := 0; i < num; i++ {
		filename := path.Join(outputPath, fmt.Sprintf("recipe_%04d.html", i))
		ioutil.WriteFile(filename, []byte(recipes[i].Content), os.ModePerm)
	}
}
