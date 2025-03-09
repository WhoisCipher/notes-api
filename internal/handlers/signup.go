package handlers

import (
    "net/mail"
    "github.com/WhoisCipher/notes-api/internal/models"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type SignUpReq struct{
    Username    string  `json:"username"`
    Email       string  `json:"email"`
    Password    string  `json:"password"`
}

func Signup(db *gorm.DB) fiber.Handler{
    return func (c *fiber.Ctx) error {
        var req SignUpReq

        if err := c.BodyParser(&req); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
        }

        if _, err := mail.ParseAddress(req.Email); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Email"})
        }

        hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Password Hashing Failed"})
        }

        user := models.User{
            Username: req.Username,
            Email: req.Email,
            Password: string(hashedPass),
        }
        if err := db.Create(&user).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User Creation Failed"})
        }

        return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User Created Successfully"})
    }
}
