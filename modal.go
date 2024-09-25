package modal


type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // "user" or "admin"
	Email    string `json:"email"`
}

// InventoryItem represents an item in the inventory
type InventoryItem struct {
	ItemName       string `json:"item_name"`
	Quantity       int    `json:"quantity"`
	ExpirationDate string `json:"expiration_date"`
	UserID         uint   `json:"user_id"` // Foreign key to User
}

// FoodRecycleDestination represents a destination for food recycling
type FoodRecycleDestination struct {
	DestinationName string `json:"destination_name"`
	Address         string `json:"address"`
	ContactNumber   string `json:"contact_number"`
	OperationalHours string `json:"operational_hours"`
}

// ExpiringItem represents an item that is about to expire
type ExpiringItem struct {
	ItemName       string `json:"item_name"`
	Quantity       int    `json:"quantity"`
	ExpirationDate string `json:"expiration_date"`
	UserID         uint   `json:"user_id"` // Foreign key to User
}

// Recipe represents a recipe that uses ingredients from the inventory
type Recipe struct {
	RecipeName string `json:"recipe_name"`
	Ingredients string `json:"ingredients"` // Can be a comma-separated list or JSON
	Instructions string `json:"instructions"`
}
