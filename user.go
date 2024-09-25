package user

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// Global DB instance
var DB *gorm.DB

// User model (define it appropriately if not already done)
type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"` // Ideally should be hashed
    Role     string `gorm:"not null"` // e.g., "user" or "admin"
}

// InventoryItem model (define it appropriately if not already done)
type InventoryItem struct {
    ID             uint   `gorm:"primaryKey"`
    ItemName       string `gorm:"not null"`
    Quantity       int    `gorm:"not null"`
    ExpirationDate string `gorm:"type:date;not null"`
}

// Register a new user
func RegisterUser(c *gin.Context) {
    var newUser User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    result := DB.Create(&newUser)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login a user
func LoginUser(c *gin.Context) {
    var loginUser User
    if err := c.ShouldBindJSON(&loginUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var user User
    result := DB.Where("username = ?", loginUser.Username).First(&user)
    if result.Error != nil || user.Password != loginUser.Password { // Normally verify hashed password
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "role": user.Role})
}
func LogoutUser(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// Get all inventory items (for the user)
func GetInventory(c *gin.Context) {
    var inventory []InventoryItem
    result := DB.Find(&inventory)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventory"})
        return
    }

    c.JSON(http.StatusOK, inventory)
}

// Get a specific inventory item by ID
func GetInventoryItem(c *gin.Context) {
    id := c.Param("id")
    var item InventoryItem

    result := DB.First(&item, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        return
    }

    c.JSON(http.StatusOK, item)
}

// Get recipe suggestions based on inventory items
func GetRecipeSuggestions(c *gin.Context) {
    // Sample recipe suggestions (you can integrate this with your AI/ML logic later)
    recipes := []string{
        "Scrambled Eggs", "Omelette", "Pancakes", "Milkshake",
    }
    c.JSON(http.StatusOK, gin.H{"recipes": recipes})
}
