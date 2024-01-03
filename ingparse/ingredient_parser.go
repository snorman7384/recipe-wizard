package ingparse

import (
	"fmt"
	"strings"

	"github.com/schollz/ingredients"
)

type StandardUnit int

const (
	FluidOunce StandardUnit = iota // for standard volume measure
	Ounce                          // for things measured in weight, not volume
	Each                           // for things like "2 eggs"
)

func (u StandardUnit) String() string {
	if u == FluidOunce {
		return "fl. oz."
	}
	if u == Ounce {
		return "oz"
	}
	if u == Each {
		return "whole"
	}
	return "<error>"
}

type Measure struct {
	OriginalAmount float64
	OriginalUnits  string
	StandardAmount float64
	StandardUnits  StandardUnit
}

type Ingredient struct {
	Line        string  // original string
	Name        string  // core item name
	Description *string // secondary info about the item, non-grouping
	Measure     Measure
}

type IngredientParser interface {
	ParseIngredientLine(line string) (Ingredient, error)
	ParseIngredients(lines []string) ([]Ingredient, error)
}

type SchollzParser struct{}

func (p SchollzParser) ParseIngredients(lines []string) ([]Ingredient, error) {
	ings, err := ingredients.ParseTextIngredients(strings.Join(lines, "\n"))

	if err != nil {
		return nil, err
	}

	ingredients := make([]Ingredient, 0, len(ings.Ingredients))

	for _, ing := range ings.Ingredients {
		ingredients = append(ingredients, p.convertIngredient(ing))
	}

	return ingredients, nil
}

func (p SchollzParser) ParseIngredientLine(line string) (Ingredient, error) {
	ings, err := ingredients.ParseTextIngredients(line)
	if err != nil {
		return Ingredient{}, err
	}

	if l := len(ings.Ingredients); l != 1 {
		return Ingredient{}, fmt.Errorf("incorrect number of ingredients: %d", l)
	}

	ing := ings.Ingredients[0]

	return p.convertIngredient(ing), nil
}

func (p SchollzParser) convertIngredient(ing ingredients.Ingredient) Ingredient {
	var desc *string = nil
	if ing.Comment != "" {
		desc = &ing.Comment
	}
	return Ingredient{
		Line:        ing.Line,
		Name:        ing.Name,
		Description: desc,
		Measure:     p.convertMeasure(ing),
	}
}

func (p SchollzParser) convertMeasure(ing ingredients.Ingredient) Measure {
	m := Measure{
		OriginalAmount: ing.Measure.Amount,
		OriginalUnits:  ing.Measure.Name,
		StandardAmount: -1,   // TODO not implemented yet
		StandardUnits:  Each, // TODO not implemented yet
	}
	return m
}