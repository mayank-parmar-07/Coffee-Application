package cron

import (
	"coffee-app/pkg/services"
	"fmt"

	"github.com/robfig/cron"
)

func SetUP(CoffeeMachine services.CoffeeMachine) {
	c := cron.New()
	c.AddFunc("@every 10s", func() { pool_verification(CoffeeMachine) })
	c.Start()
	fmt.Println("Cron SetUp Done")

}

func pool_verification(CoffeeMachine services.CoffeeMachine) {
	IngredientsManager := CoffeeMachine.IngredientManager
	for ingredientName, ingredientObj := range IngredientsManager.Ingredients {
		if ingredientObj.Quantity == 0 {
			fmt.Println("Please refill", ingredientName)
		}
	}
}
