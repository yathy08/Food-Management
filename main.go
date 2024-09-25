package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // User Routes
    router.POST("/register", RegisterUser)  // User registration
    router.POST("/login", LoginUser)        // User login

    // Inventory Routes
    router.GET("/inventory", GetInventory)               // Get all inventory items
    router.POST("/inventory", AddInventoryItem)          // Add a new inventory item
    router.PUT("/inventory/:id", UpdateInventoryItem)    // Update an existing inventory item
    router.DELETE("/inventory/:id", DeleteInventoryItem)  // Delete an inventory item

    // Recipe Suggestions Route
    router.GET("/recipes", GetRecipeSuggestions) // Get recipe suggestions based on inventory

    // Start the server
    router.Run(":8080") // You can change the port if needed
}
