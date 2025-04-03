package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/charmbracelet/huh"
)

func validateNVal(maxCriteria int) func(s string) error {
	return func(s string) error {
		n, err := strconv.Atoi(s)
		maxS := strconv.Itoa(maxCriteria)
		if err != nil {
			return fmt.Errorf("input är inte en siffra")
		} else if n > maxCriteria {
			return fmt.Errorf("för högt max input är " + maxS)
		} else if n < 1 {
			return fmt.Errorf("för lågt minimum input är 1")
		}
		return nil
	}
}

func contains(value map[string]string, collection []map[string]string) bool {
	for _, curr := range collection {
		if value["id"] == curr["id"] {
			return true
		}
	}
	return false
}

func containsNTimes(value map[string]string, collection []map[string]string, n int) bool {
	counter := 0
	for _, curr := range collection {
		if counter == n-1 {
			return true
		} else if value["id"] == curr["id"] {
			counter++
		}
	}
	return false
}

func getRecipes(foodChoice FoodChoice) []map[string]string {

	recipes := []map[string]string{}
	reflFoodChioce := reflect.ValueOf(foodChoice)
	var n string
	maxCriteria := 0
	for i := range reflFoodChioce.NumField() {
		value := reflFoodChioce.Field(i).String()
		if value != "" {
			maxCriteria++
		}
	}
	maxS := strconv.Itoa(maxCriteria)
	startForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Hur många kriterier måste uppfyllas? (1-" + maxS + ")").
				Validate(validateNVal(maxCriteria)).
				Value(&n),
		),
	)

	err := startForm.Run()
	if err != nil {
		log.Fatal(err)
	}

	nInt, _ := strconv.Atoi(n)

	collection := []map[string]string{}
	for i := range reflFoodChioce.NumField() {
		value := reflFoodChioce.Field(i).String()
		if value != "" {
			filename := "tags/" + reflFoodChioce.Field(i).String() + ".json"
			file, _ := os.ReadFile(filename)
			contentFile, _ := content.ReadFile(filename)
			var localData map[string][]map[string]string
			var data map[string][]map[string]string

			json.NewDecoder(bytes.NewBuffer(contentFile)).Decode(&localData)
			collection = append(collection, localData["data"]...)

			json.NewDecoder(bytes.NewBuffer(file)).Decode(&data)
			collection = append(collection, data["data"]...)
		}
	}

	for i := range collection {
		currValue := collection[i]

		if containsNTimes(currValue, collection[i+1:], nInt) && !contains(currValue, recipes) {
			recipes = append(recipes, currValue)
		}
	}
	return recipes
}
