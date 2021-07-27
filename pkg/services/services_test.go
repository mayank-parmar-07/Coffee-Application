package services

import (
	"coffee-app/pkg/config"
	"coffee-app/pkg/models"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var CoffeeMachineObj = CoffeeMachine{}

func init() {
	logger := config.SetupLogging()
	Machine := models.Machine{}
	jsonFile, err := os.Open("./test_config.json") //Parsing the JSON config file
	if err != nil {
		logger.Fatalf("Error while loading config file : %s", err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Machine)
	CoffeeMachineObj = NewCoffeMachine(Machine, logger)
}

func TestTakeIngredients(t *testing.T) {
	CoffeeMachineObj.IngredientManager.TakeIngredients([]models.IngredientRequirement{{Name: "sugar_syrup", Quantity: 1}})
}

func TestCheckIngredient(t *testing.T) {
	error := CoffeeMachineObj.IngredientManager.CheckIngredient("sugar_syrup", 50, "hot tea")
	assert.NoError(t, error)
}

func TestRefill(t *testing.T) {
	error := CoffeeMachineObj.IngredientManager.Refill("sugar_syrup", 50)
	assert.NoError(t, error)
}

func TestFetchIngredient(t *testing.T) {
	_, error := CoffeeMachineObj.BeverageManager.FetchIngredient("hot_tea")
	assert.NoError(t, error)
}

// Below are negative test cases
func TestNegCheckIngredient(t *testing.T) {
	error := CoffeeMachineObj.IngredientManager.CheckIngredient("sugar_syrup", 1000, "hot_tea")
	assert.Error(t, error)
}

func TestNegRefill(t *testing.T) {
	error := CoffeeMachineObj.IngredientManager.Refill("sugar_syrup", 10000)
	assert.Error(t, error)
}

func TestNegFetchIngredient(t *testing.T) {
	_, error := CoffeeMachineObj.BeverageManager.FetchIngredient("spa")
	assert.Error(t, error)
}
