package entity

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusCooking   OrderStatus = "Cooking"
	OrderStatusCooked    OrderStatus = "Cooked"
	OrderStatusPickedUp  OrderStatus = "PickedUp"
	OrderStatusDelivered OrderStatus = "Delivered"
)

type OrderItemOption struct {
	Name   string  `json:"name"`
	Choice *string `json:"choice,omitempty"`
}

type OrderItem struct {
	CoreEntity
	DishID int
	Dish   Dish
	Option []OrderItemOption `gorm:"type:json"`
}

type Order struct {
	CoreEntity
	CustomerID   *int
	Customer     *User
	DriverID     *int
	Driver       *User
	RestaurantID *int
	Restaurant   *Restaurant
	Items        []OrderItem `gorm:"many2many:order_order_items;"`
	Total        *float64
	Status       OrderStatus `gorm:"default:'Pending'"`
}
