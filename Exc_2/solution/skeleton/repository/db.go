package repository

import "ordersystem/model"

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo - DONE
func NewDatabaseHandler() *DatabaseHandler {
	drinks := []model.Drink{
		{ID: 1, Name: "Espresso",     Price: 2.2, Description: "Strong Italian coffee shot"},
		{ID: 2, Name: "Americano",    Price: 2.5, Description: "Espresso with hot water"},
		{ID: 3, Name: "Cappuccino",   Price: 2.9, Description: "Espresso with foamed milk"},
		{ID: 4, Name: "Latte",        Price: 3.1, Description: "Espresso with steamed milk"},
		{ID: 5, Name: "Water Sparkling", Price: 1.0, Description: "Bubbly, refreshing H2O"},
	}
	var orders []model.Order

	// Init orders slice with some test data.

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	totalledOrders := make(map[uint64]uint64)
	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += 1
	}
	// key = DrinkID, value = Amount of orders
	return totalledOrders
}

// todo - DONE
func (db *DatabaseHandler) AddOrder(order *model.Order) {
	db.orders = append(db.orders, *order)
	// add order to db.orders slice
}
