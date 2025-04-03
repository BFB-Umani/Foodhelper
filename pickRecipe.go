package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/huh"
)

func fetchOptions(recipes []map[string]string) []string {
	receipeNames := []string{}
	for _, curr := range recipes {
		receipeNames = append(receipeNames, curr["name"])
	}
	return receipeNames
}

func findRecipe(choice string, recipes []map[string]string) map[string]string {
	for _, recipe := range recipes {
		if recipe["name"] == choice {
			return recipe
		}
	}
	return nil
}

func pickRecipe(recipes []map[string]string) {
	var choice string
	startForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Välj vilket recept du vill läsa").
				Value(&choice).
				OptionsFunc(func() []huh.Option[string] {
					optionsList := []huh.Option[string]{}
					options := fetchOptions(recipes)
					for _, dish := range options {
						optionsList = append(optionsList, huh.NewOption(dish, dish))
					}
					return optionsList
				}, ""),
		),
	)

	err := startForm.Run()
	if err != nil {
		log.Fatal(err)
	}

	chosenRecipe := findRecipe(choice, recipes)

	input, err := os.Open("recipes/" + chosenRecipe["id"] + "/" + chosenRecipe["name"] + ".md")
	if err != nil {
		contentInput, _ := content.Open("recipes/" + chosenRecipe["id"] + "/" + chosenRecipe["name"] + ".md")
		markdown, err := io.ReadAll(contentInput)
		if err != nil {
			log.Fatal(err)
		}
		out, _ := glamour.Render(string(markdown), "dark")
		fmt.Print(out)
	} else {
		markdown, err := io.ReadAll(input)
		if err != nil {
			log.Fatal(err)
		}
		out, _ := glamour.Render(string(markdown), "dark")
		fmt.Print(out)
	}
	defer input.Close()
}
