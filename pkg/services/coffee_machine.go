package services

import (
	"coffee-app/pkg/models"
	"fmt"
	"log"
	"sync"
)

type CoffeeMachine struct {
	Contents          []models.Ingredient
	BeverageManager   BeverageManager
	Outlets           int
	Logger            *log.Logger
	IngredientManager IngredientManager
	Buffer            chan bool //Buffer channel for the outlets to not keep other waiting.
}

func (svc *CoffeeMachine) MakeCoffee(beverage string) {
	err := svc.CheckIngredient(beverage)
	if err == nil {
		svc.Display("Your " + beverage + " is ready")
		return
	}
	svc.Buffer <- true

}

func (svc *CoffeeMachine) AddRequest(beverage string) {
	<-svc.Buffer
	svc.Logger.Printf("Request for %s \n", beverage)
	svc.MakeCoffee(beverage)
	return
}

func (svc *CoffeeMachine) CheckIngredient(beverage string) error {
	mutex := sync.Mutex{} //Synchronizing the execution of this function by mutex
	mutex.Lock()
	defer mutex.Unlock()
	ingredients, err := svc.BeverageManager.FetchIngredient(beverage)
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		ingredientsSlice := []models.IngredientRequirement{}
		for ingredientName, ingredientRequired := range ingredients.Recipe {
			err = svc.IngredientManager.CheckIngredient(ingredientName, ingredientRequired, beverage)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			ingredientsSlice = append(ingredientsSlice, models.IngredientRequirement{Name: ingredientName, Quantity: ingredientRequired})
		}
		svc.IngredientManager.TakeIngredients(ingredientsSlice)
		return nil
	}
}

func (svc *CoffeeMachine) Display(message string) {
	fmt.Println(message)
}

func NewCoffeMachine(config models.Machine, logger *log.Logger) CoffeeMachine {
	NewMachine := CoffeeMachine{}
	machine_configuration := config.Machine
	NewMachine.Outlets = machine_configuration.Outlets.Count
	ingredients := []models.Ingredient{}
	NewMachine.Contents = ingredients
	NewMachine.IngredientManager = IngredientManager{}
	NewMachine.IngredientManager.Ingredients = make(map[string]models.Ingredient)
	NewMachine.IngredientManager.Logger = logger
	for ingredientName, quantity := range machine_configuration.QuantitiesConfig {
		ingredientQuantity := models.Ingredient{Quantity: quantity, MaxQuantity: 400} //Adding default max quantity
		NewMachine.IngredientManager.Ingredients[ingredientName] = ingredientQuantity
	}
	NewMachine.BeverageManager = BeverageManager{}
	NewMachine.BeverageManager.Beverage = make(map[string]models.Beverage)
	NewMachine.BeverageManager.Logger = logger
	NewMachine.Logger = logger
	for beverageName, recipe := range machine_configuration.Beverages {
		beverageObj := models.Beverage{Name: beverageName, Recipe: recipe}
		NewMachine.BeverageManager.Beverage[beverageName] = beverageObj
	}
	NewBuffer := make(chan bool, NewMachine.Outlets)
	for outlets := 0; outlets < NewMachine.Outlets; outlets++ {
		NewBuffer <- true
	}
	NewMachine.Buffer = NewBuffer
	NewMachine.Logger.Printf("New Coffee machine.........")
	return NewMachine
}
