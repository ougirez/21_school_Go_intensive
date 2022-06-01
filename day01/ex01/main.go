package main

import (
	"fmt"
	"log"
	"os"
)

func findCake(cake Cake, recipes Recipes) Cake {
	for _, c := range recipes.Cakes {
		if cake.Name == c.Name {
			return c
		}
	}
	return Cake{}
}

func findIngredient(ingredient Ingredient, cake Cake) Ingredient {
	for _, ing := range cake.Ingredients {
		if ingredient.Name == ing.Name {
			return ing
		}
	}
	return Ingredient{}
}

func compare(oldRecipes, newRecipes Recipes) {
	for _, newCake := range newRecipes.Cakes {
		foundedCake := findCake(newCake, oldRecipes)
		if foundedCake.Name == "" {
			fmt.Printf("ADDED cake \"%s\"\n", newCake.Name)
			continue 
		}
	}
	for _, oldCake := range oldRecipes.Cakes {
		foundedCake := findCake(oldCake, newRecipes)
		if foundedCake.Name == "" {
			fmt.Printf("REMOVED cake \"%s\"\n", oldCake.Name)
		}
	}

	for _, newCake := range newRecipes.Cakes {
		foundedCake := findCake(newCake, oldRecipes)
		if foundedCake.Name == "" {
			continue 
		}
		if foundedCake.Time != newCake.Time {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", foundedCake.Name, newCake.Time, foundedCake.Time)
		}
		for _, newIngredient := range newCake.Ingredients {
			foundedIngredient := findIngredient(newIngredient, foundedCake)
			if foundedIngredient.Name == "" {
				fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIngredient.Name, newCake.Name)
				continue 
			}
			if foundedIngredient.Count != newIngredient.Count {
				fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", foundedIngredient.Name, newCake.Name, newIngredient.Count, foundedIngredient.Count)
			}
			if newIngredient.Unit != "" && foundedIngredient.Unit == "" {
				fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", newIngredient.Unit, newIngredient.Name, newCake.Name)
			}
			if newIngredient.Unit == "" && foundedIngredient.Unit != "" {
				fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", foundedIngredient.Unit, newIngredient.Name, newCake.Name)
			}
			if newIngredient.Unit != "" && foundedIngredient.Unit != "" && newIngredient.Unit != foundedIngredient.Unit {
				fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", newIngredient.Name, newCake.Name, newIngredient.Unit, foundedIngredient.Unit)
			}
		}
		
		for _, oldIngredient := range foundedCake.Ingredients {
			foundedIngredient := findIngredient(oldIngredient, newCake)
			if foundedIngredient.Name == "" {
				fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", oldIngredient.Name, newCake.Name)
				continue 
			}
		}
		
	}
}

func main() {
	var oldDB, newDB string
	var oldRecipes, newRecipes Recipes

	if len(os.Args) != 5 {
		log.Fatal("Invalid number of arguments")
	}
	for i := 1; i < len(os.Args); i += 2 {
		if os.Args[i] == "--old" {
			oldDB = os.Args[i+1]
		} else if os.Args[i] == "--new" {
			newDB = os.Args[i+1]
		} else {
			log.Fatal("Use these flags: --old, --new")
		}
	}
	var oldDBReader = getReader(oldDB)
	var newDBReader = getReader(newDB)
	oldRecipes = oldDBReader.parseFile()
	newRecipes = newDBReader.parseFile()
	// contains(oldRecipes.Cakes[0], oldRecipes, newRecipes)
	compare(oldRecipes, newRecipes)
}
