package services

import (
	"coffee-app/pkg/models"
	"fmt"
	"log"
)

type IngredientManager struct {
	Ingredients map[string]models.Ingredient
	Logger      *log.Logger
}

func (svc *IngredientManager) Refill(ingredient string, refill_quantity int) error {
	if _, ok := svc.Ingredients[ingredient]; !ok {
		svc.Logger.Printf("Creating a new ingredient : %s\n", ingredient)
		svc.Ingredients[ingredient] = models.Ingredient{Quantity: refill_quantity, MaxQuantity: refill_quantity}
		return nil
	} else {
		ingredientObj := svc.Ingredients[ingredient]
		if ingredientObj.Quantity+refill_quantity > ingredientObj.MaxQuantity {
			return fmt.Errorf("Excess amount added")
		} else {
			ingredientObj.Quantity = svc.Ingredients[ingredient].Quantity + refill_quantity
			svc.Ingredients[ingredient] = ingredientObj
			return nil
		}
	}
}

func (svc *IngredientManager) CheckIngredient(ingredient string, quantity_req int, beverage string) error {
	if _, ok := svc.Ingredients[ingredient]; !ok {
		svc.Logger.Printf("%s cannot be prepared because %s is not available\n", beverage, ingredient)
		return fmt.Errorf("%s cannot be prepared because %s is not available", beverage, ingredient)
	} else {
		ingredientObj := svc.Ingredients[ingredient]
		if ingredientObj.Quantity < quantity_req {
			svc.Logger.Printf("%s cannot be prepared because %s is not sufficient\n", beverage, ingredient)
			return fmt.Errorf("%s cannot be prepared because %s is not sufficient", beverage, ingredient)
		}
		return nil
	}
}

func (svc *IngredientManager) TakeIngredients(requirement []models.IngredientRequirement) {
	for _, ingredient := range requirement {
		ingredientObj := svc.Ingredients[ingredient.Name]
		ingredientObj.Quantity = ingredientObj.Quantity - ingredient.Quantity
		svc.Ingredients[ingredient.Name] = ingredientObj
	}
}
