package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
)

type NewRecipe struct {
	name         string
	portions     string
	ingredients  string
	instructions string
	tags         []string
}

func validatePortionVal() func(s string) error {
	return func(s string) error {
		n, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("input är inte en siffra")
		} else if n > 10 {
			return fmt.Errorf("för högt max input är 10")
		} else if n < 1 {
			return fmt.Errorf("för lågt minimum input är 1")
		}
		return nil
	}
}

func createRecipe() NewRecipe {
	newRecipe := NewRecipe{}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Fyll i namn på receptet").
				Value(&newRecipe.name),

			huh.NewInput().
				Title("Hur många portioner? (1-10)").
				Validate(validatePortionVal()).
				Value(&newRecipe.portions),

			huh.NewText().
				Title("Fyll i ingredienser och mått separerat med --, t.ex 'soya 1 msk -- ris 2dl'").
				Value(&newRecipe.ingredients).CharLimit(50000),

			huh.NewText().
				Title("Fyll i instruktioner separerat med med --, t.ex 'koka ris -- häll på soya'").
				Value(&newRecipe.instructions).CharLimit(50000),

			huh.NewMultiSelect[string]().
				Options(
					huh.NewOption("Nötkött", "beef"),
					huh.NewOption("Kyckling", "chicken"),
					huh.NewOption("Fisk/skaldjur", "fish"),
					huh.NewOption("Fläsk", "pork"),
					huh.NewOption("Vegetarianskt (lakto-ovo)", "vegetarian"),
					huh.NewOption("Linser", "lentils"),
					huh.NewOption("Grönsaker", "vegetables"),
					huh.NewOption("Svamp", "mushroom"),
					huh.NewOption("Vilt", "wild"),
					huh.NewOption("Lamm", "lamb"),
					huh.NewOption("Anka", "duck"),
					huh.NewOption("Rotfrukt", "roots"),
					huh.NewOption("Ris", "rice"),
					huh.NewOption("Pasta", "pasta"),
					huh.NewOption("Deg (Pizza/Paj)", "pastry"),
					huh.NewOption("Nudlar", "noodles"),
					huh.NewOption("Gryta", "stew"),
					huh.NewOption("Pastasås", "pastaSauce"),
					huh.NewOption("Soppa", "soup"),
					huh.NewOption("Sallad", "salad"),
					huh.NewOption("Torr rätt", "dry"),
					huh.NewOption("Gratäng", "gratin"),
					huh.NewOption("Macka", "sandwich"),
					huh.NewOption("Gräddbaserad", "creamBased"),
					huh.NewOption("Tomatbaserad", "tomatoBased"),
					huh.NewOption("Buljong", "broth"),
				).
				Title("Lägg till taggar för mat preferenser").
				Value(&newRecipe.tags),
		),
	)

	fmt.Println()
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return newRecipe
}
