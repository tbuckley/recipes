package empty

import (
	"github.com/tbuckley/recipes/db"
)

func Run(arguments map[string]interface{}) {
	err := db.DeleteRecipes()
	if err != nil {
		panic(err)
	}
}
