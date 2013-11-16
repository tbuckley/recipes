package sample

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/tbuckley/recipes/db"
)

var (
	sampleSize = flag.Int("size", 10, "")
	samplePath = flag.String("path", ".", "")
)

func Run() {
	finfo, err := os.Lstat(*samplePath)
	if err != nil {
		panic(err)
	}
	if !finfo.IsDir() {
		fmt.Printf("Invalid path: %s\n", *samplePath)
	}

	err = db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	recipes, err := db.AllRecipes()
	if err != nil {
		panic(err)
	}

	var num int = 0
	if *sampleSize < len(recipes) {
		num = *sampleSize
	} else {
		num = len(recipes)
	}
	for i := 0; i < num; i++ {
		filename := path.Join(*samplePath, fmt.Sprintf("recipe_%04d.html", i))
		ioutil.WriteFile(filename, []byte(recipes[i].Content), os.ModePerm)
	}
}
