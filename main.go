package main

import (
	"coffee-app/pkg/config"
	"coffee-app/pkg/models"
	"coffee-app/pkg/services"
	"encoding/json"
	"io/ioutil"
	"os"
)

/*
Entry point
*/
func main() {
	logger := config.SetupLogging()
	Machine := models.Machine{}
	jsonFile, err := os.Open("config.json") //Parsing the JSON config file
	if err != nil {
		logger.Fatalf("Error while loading config file : %s", err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Machine)
	CoffeeMachine := services.NewCoffeMachine(Machine, logger)
	beverageRequest := []string{"green_tea", "hot_tea", "hot_coffee", "black_tea"}
	for _, beverage := range beverageRequest {
		CoffeeMachine.AddRequest(beverage)
	}

}
