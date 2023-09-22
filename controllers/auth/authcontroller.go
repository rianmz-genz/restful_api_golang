package authcontroler

import (
	"hello-world/helpers"
	"hello-world/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
var request struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

func Register(c *gin.Context) {

    if err := c.ShouldBindJSON(&request); err != nil {
        helpers.BadRequestResponse(c, "Data yang dimasukkan tidak valid")
        return
    }

    // Cek apakah username sudah digunakan
    var existingUser models.User
    if err := models.DB.Where("username = ?", request.Username).First(&existingUser).Error; err == nil {
        helpers.BadRequestResponse(c, "Username sudah digunakan")
        return
    }

    // Hash password sebelum menyimpan ke database
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
    if err != nil {
        helpers.InternalServerResponse(c, "Gagal mengenkripsi password")
        return
    }

    newUser := models.User{
        Username: request.Username,
        Password: string(hashedPassword),
    }

    if err := models.DB.Create(&newUser).Error; err != nil {
        helpers.InternalServerResponse(c, "Gagal membuat pengguna")
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "status":  true,
        "message": "Pengguna berhasil terdaftar",
    })
}

func Login(c *gin.Context) {
    

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  false,
            "message": "Data yang dimasukkan tidak valid",
        })
        return
    }

    var user models.User
    if err := models.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  false,
            "message": "Username atau password salah",
        })
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  false,
            "message": "Username atau password salah",
        })
        return
    }
	
    token := CreateToken(user.ID)
    c.JSON(http.StatusOK, gin.H{
        "status":  true,
        "message": "Berhasil login",
        "token":   token,
    })
}

func CreateToken(userID uint) string {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    tokenString, _ := token.SignedString([]byte("restfulapigolang"))

    return tokenString
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            helpers.UnAuthorizedResponse(c, "Token tidak ditemukan")
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("restfulapigolang"), nil
        })

        if err != nil || !token.Valid {
            helpers.UnAuthorizedResponse(c, "Token tidak valid")
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            helpers.UnAuthorizedResponse(c, "Token tidak valid")
            c.Abort()
            return
        }

        userID := uint(claims["user_id"].(float64))
        var user models.User
        if err := models.DB.First(&user, userID).Error; err != nil {
            switch err {
            case gorm.ErrRecordNotFound:
                helpers.NotFoundResponse(c)
            default:
                helpers.InternalServerResponse(c, err.Error())
            }
            c.Abort()
            return
        }

        c.Set("user", user)
        c.Next()
    }
}