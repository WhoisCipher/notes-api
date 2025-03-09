package utils

import(
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(id uint) (string, error){
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": id,
        "exp": time.Now().Add(time.Hour * 2).Unix(),
    })

    return token.SignedString(secretKey)
}
