package models

type Machine struct {
	Machine MachineConfiguration `json:"machine"`
}

type MachineConfiguration struct {
	Outlets          OutletsConfig             `json:"outlets"`
	QuantitiesConfig map[string]int            `json:"total_items_quantity"`
	Beverages        map[string]map[string]int `json:"beverages"`
}

type OutletsConfig struct {
	Count int `json:"count_n"`
}
