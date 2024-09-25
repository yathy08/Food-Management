package admin

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "FoodWaste-Management/modal"
)

func AddInventoryItem(c *gin.Context) {
    var newItem modal.InventoryItem
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Assuming admin role is verified
    result := DB.Create(&newItem)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Item added successfully", "item": newItem})
}

func GetInventory(c *gin.Context) {
    var inventory []modal.InventoryItem
    result := DB.Preload("User").Find(&inventory)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventory"})
        return
    }

    c.JSON(http.StatusOK, inventory)
}

func UpdateInventoryItem(c *gin.Context) {
    id := c.Param("id")
    var item modal.InventoryItem

    // Check if item exists
    result := DB.First(&item, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        return
    }

    // Update item
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    DB.Save(&item)
    c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully", "item": item})
}

func DeleteInventoryItem(c *gin.Context) {
    id := c.Param("id")
    var item modal.InventoryItem

    // Check if item exists
    result := DB.First(&item, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        return
    }

    DB.Delete(&item)
    c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

func GetAllUsers(c *gin.Context) {
    var users []modal.User
    result := DB.Find(&users)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
        return
    }

    c.JSON(http.StatusOK, users)
}
