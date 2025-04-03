package main

import (
	"embed"
	"log"

	"github.com/charmbracelet/huh"
)

var (
	//go:embed recipes tags
	content embed.FS
)

var registerNewRecipe bool

func main() {

	startForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[bool]().
				Title("Vill du lägga till ett recept eller få förslag på maträtter?").
				Options(
					huh.NewOption("Lägg till nytt recept", true),
					huh.NewOption("Få förslag på recept", false),
				).
				Value(&registerNewRecipe),
		),
	)

	err := startForm.Run()
	if err != nil {
		log.Fatal(err)
	}

	if registerNewRecipe {
		newRecipe := createRecipe()
		createMarkdown(newRecipe)
	} else {
		foodChoice := foodForm()
		recipes := getRecipes(foodChoice)
		pickRecipe(recipes)
	}

}
