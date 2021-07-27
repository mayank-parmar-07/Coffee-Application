package services

import (
	"coffee-app/pkg/models"
	"fmt"
	"log"
)

type BeverageManager struct {
	Beverage map[string]models.Beverage
	Logger   *log.Logger
}

func (svc *BeverageManager) FetchIngredient(beverage string) (*models.Beverage, error) {
	if _, ok := svc.Beverage[beverage]; ok {
		result := svc.Beverage[beverage]
		return &result, nil
	} else {
		svc.Logger.Printf("%s cannot be prepared because there is no such beverage\n", beverage)
		return nil, fmt.Errorf("%s cannot be prepared because there is no such beverage", beverage)
	}
}
