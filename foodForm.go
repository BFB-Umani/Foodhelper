package main

import (
	"fmt"
	"log"

	"slices"

	"github.com/charmbracelet/huh"
)

type FoodChoice struct {
	protein     string
	accessories string
	carb        string
	dishType    string
	pastaType   string
}

func fetchOptionsForCarb(carb string) ([]string, []string) {
	options := []string{"Inget", "Gryta", "Pastasås", "Soppa", "Sallad", "Torr rätt", "Gratäng", "Macka"}
	values := []string{"", "stew", "pastaSauce", "soup", "salad", "dry", "gratin", "sandwich"}

	if carb == "pasta" {
		return options[1:], values[1:]
	} else if carb == "" {
		return options, values
	}
	options = slices.Delete(options, 1, 2)
	values = slices.Delete(values, 1, 2)

	return options, values
}

func foodForm() FoodChoice {
	foodChoice := FoodChoice{}
	form := huh.NewForm(
		huh.NewGroup(
			// Ask the user for a base to build the dish.
			huh.NewSelect[string]().
				Title("Välj ett protein").
				Options(
					huh.NewOption("Inget", ""),
					huh.NewOption("Nötkött", "beef"),
					huh.NewOption("Kyckling", "chicken"),
					huh.NewOption("Fisk/skaldjur", "fish"),
					huh.NewOption("Fläsk", "pork"),
					huh.NewOption("Vegetarianskt (lakto-ovo)", "vegetarian"),
					huh.NewOption("Vilt", "wild"),
					huh.NewOption("Lamm", "lamb"),
					huh.NewOption("Anka", "duck"),
				).
				Value(&foodChoice.protein), // store the chosen option in the "protein" variable

			// Ask the user for accessories to the dish.
			huh.NewSelect[string]().
				Title("Välj tillbehör").
				Options(
					huh.NewOption("Inget", ""),
					huh.NewOption("Linser", "lentils"),
					huh.NewOption("Grönsaker", "vegetables"),
					huh.NewOption("Svamp", "mushroom"),
				).
				Value(&foodChoice.accessories),

			// Ask the user for carbs to build the dish.
			huh.NewSelect[string]().
				Title("Välj typ av kolhydrat").
				Options(
					huh.NewOption("Inget", ""),
					huh.NewOption("Rotfrukt", "roots"),
					huh.NewOption("Ris", "rice"),
					huh.NewOption("Pasta", "pasta"),
					huh.NewOption("Deg (Pizza/Paj)", "pastry"),
					huh.NewOption("Nudlar", "noodles"),
				).
				Value(&foodChoice.carb),

			// Ask the user for what kind of dish they want.
			huh.NewSelect[string]().
				Title("Välj typ av rätt").
				Value(&foodChoice.dishType).
				OptionsFunc(func() []huh.Option[string] {
					optionsList := []huh.Option[string]{}
					options, values := fetchOptionsForCarb(foodChoice.carb)
					for i, dish := range options {
						optionsList = append(optionsList, huh.NewOption(dish, values[i]))
					}
					return optionsList
				}, &foodChoice.carb),

			// Ask the user what type of sauce they want for the dish.
			huh.NewSelect[string]().
				Title("Vilken typ av sås?").
				Options(
					huh.NewOption("Inget", ""),
					huh.NewOption("Gräddbaserad", "creamBased"),
					huh.NewOption("Tomatbaserad", "tomatoBased"),
					huh.NewOption("Buljong", "broth"),
				).
				Value(&foodChoice.pastaType), // store the chosen option in the "protein" variable
		),
	)
	fmt.Println()
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return foodChoice
}
