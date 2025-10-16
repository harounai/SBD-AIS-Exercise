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
		{ID: 1, Name: "Beer", Price: 2.0, Description: "Hagenberger Gold"},
		{ID: 2, Name: "Spritzer", Price: 1.4, Description: "Wine with Soda"},
		{ID: 3, Name: "Coffee", Price: 0.0, Description: "Mifare isn't that secure ;)"},
	}
	var orders []model.Order

	// Init orders slice with some test data

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

// todo - DONE
// func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
// 	for _, order := range db.orders {
// 		totalledOrders := make(map[uint64]uint64)
// 		totalledOrders[order.DrinkID] += 1
// 	}
// 	// key = DrinkID, value = Amount of orders
// 	// totalledOrders map[uint64]uint64

// 	return totalledOrders

func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	totalledOrders := make(map[uint64]uint64)
	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += 1
	}
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64

	return totalledOrders
}

// todo - DONE
func (db *DatabaseHandler) AddOrder(order *model.Order) {
	db.orders = append(db.orders, *order)
	// add order to db.orders slice
}
