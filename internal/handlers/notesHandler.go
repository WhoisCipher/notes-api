package handlers

import (
	"strconv"

	"github.com/WhoisCipher/notes-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func CreateNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userToken := c.Locals("user").(*jwt.Token)
		claims := userToken.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(float64)

		// Parse request
		var note models.Note
		if err := c.BodyParser(&note); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		// Assign user ID from JWT
		note.UserID = uint(userID)

		// Save note in DB
		if err := db.Create(&note).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create note"})
		}

		return c.Status(fiber.StatusCreated).JSON(note)
	}
}

// Get all notes of the logged in User
func GetNotes(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userToken := c.Locals("user").(*jwt.Token)
		claims := userToken.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(float64)

		var notes []models.Note
		if err := db.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch notes"})
		}

		return c.Status(fiber.StatusOK).JSON(notes)
	}
}

// Update the note of given ID of logged in User
func UpdateNotes(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userToken, ok := c.Locals("user").(*jwt.Token)
        if !ok {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
        }

        claims, ok := userToken.Claims.(jwt.MapClaims)
        if !ok {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
        }

        userID, ok := claims["user_id"].(float64)
        if !ok {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID"})
        }

        noteID, err := strconv.Atoi(c.Params("id"))
        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid note ID"})
        }

        var note models.Note

        if err := db.First(&note, "id = ? AND user_id = ?", noteID, uint(userID)).Error; err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
        }

        var updateData models.Note
        if err := c.BodyParser(&updateData); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
        }

        if updateData.Title != "" {
            note.Title = updateData.Title
        }
        if updateData.Content != "" {
            note.Content = updateData.Content
        }

        if err := db.Save(&note).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update note"})
        }

        return c.Status(fiber.StatusOK).JSON(note)
    }
}

// Delete a note of given userID
func DeleteNote(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userToken := c.Locals("user").(*jwt.Token)
        claims := userToken.Claims.(jwt.MapClaims)
        userID := uint(claims["user_id"].(float64))

        noteID, err := strconv.Atoi(c.Params("id"))
        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid note ID"})
        }

        var note models.Note

        if err := db.First(&note, "id = ? AND user_id = ?", noteID, userID).Error; err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
        }
        if err := db.Delete(&note).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete note"})
        }

        return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Note deleted successfully"})
    }
}

