package main

type FoodItem struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Calories       int    `json:"calories"`
	ExpirationDate string `json:"expiration_date"`
}

var foodItems = []*FoodItem{
	{
		ID:             1,
		Name:           "Apple",
		Description:    "A fruit",
		Calories:       52,
		ExpirationDate: "2021-12-31",
	},
	{
		ID:             2,
		Name:           "Cheerios",
		Description:    "A cereal",
		Calories:       100,
		ExpirationDate: "2022-12-31",
	},
	{
		ID:             3,
		Name:           "Doritos",
		Description:    "A snack",
		Calories:       150,
		ExpirationDate: "2022-12-31",
	},
}

var nextID = 4

func listFoodItems() []*FoodItem {
	return foodItems
}

func getFoodItem(id int) *FoodItem {
	for _, foodItem := range foodItems {
		if foodItem.ID == id {
			return foodItem
		}
	}
	return nil
}

func storeFoodItem(foodItem FoodItem) FoodItem {
	foodItem.ID = nextID
	nextID++
	foodItems = append(foodItems, &foodItem)
	return foodItem
}

func deleteFoodItem(id int) *FoodItem {
	for i, foodItem := range foodItems {
		if foodItem.ID == id {
			foodItems = append(foodItems[:i], foodItems[i+1:]...)
			return foodItem
		}
	}
	return nil
}

func updateFoodItem(id int, foodItemUpdate FoodItem) *FoodItem {
	for i, foodItem := range foodItems {
		if foodItem.ID == id {
			foodItemUpdate.ID = id
			foodItems[i] = &foodItemUpdate
			return &foodItemUpdate
		}
	}
	return nil
}
