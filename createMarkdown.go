package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"unicode"

	"github.com/google/uuid"
	md "github.com/nao1215/markdown"
)

func makeData(id, name string, prevData []map[string]string) any {
	newData := map[string]string{
		"name": name,
		"id":   id,
	}
	prevData = append(prevData, newData)
	data := map[string][]map[string]string{
		"data": prevData,
	}

	return data
}

func capitalize(s string) string {
	// Need to convert to rune slice and capitalize because of weird interactions using string.ToUpper() on ÅÄÖ
	runeArr := []rune(s)
	capitalizedRunes := []rune{unicode.ToUpper(runeArr[0])}
	capitalizedRunes = append(capitalizedRunes, runeArr[1:]...)
	return string(capitalizedRunes)
}

func createMarkdown(newRecipe NewRecipe) {
	ID, _ := uuid.NewV7()

	os.Mkdir("recipes", os.ModePerm)

	os.Mkdir("recipes/"+ID.String(), os.ModePerm)

	for _, tag := range newRecipe.tags {
		filename := "tags/" + tag + ".json"
		file, err := os.ReadFile(filename)
		if err != nil {
			var data map[string][]map[string]string
			jsonString, _ := json.Marshal(makeData(ID.String(), capitalize(newRecipe.name), data["data"]))
			os.Mkdir("tags", os.ModePerm)
			jsonFile, err := os.Create(filename)
			if err != nil {
				panic(err)
			}
			defer jsonFile.Close()

			os.WriteFile(filename, jsonString, os.ModePerm)

		} else {
			var data map[string][]map[string]string
			json.NewDecoder(bytes.NewBuffer(file)).Decode(&data)
			jsonString, _ := json.Marshal(makeData(ID.String(), capitalize(newRecipe.name), data["data"]))
			os.WriteFile(filename, jsonString, os.ModePerm)
		}

	}

	file, err := os.Create("recipes/" + ID.String() + "/" + capitalize(newRecipe.name) + ".md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ingredients := []string{}
	instructions := []string{}

	for _, ingr := range strings.Split(newRecipe.ingredients, "--") {
		ingredients = append(ingredients, capitalize(strings.Trim(ingr, " ")))
	}
	for _, inst := range strings.Split(newRecipe.instructions, "--") {
		instructions = append(instructions, capitalize(strings.Trim(inst, " ")))
	}
	md.NewMarkdown(file).H1("För " + newRecipe.portions + " person(er)").H1("Ingredienser").BulletList(ingredients...).H1("Instruktioner").OrderedList(instructions...).Build()
}
