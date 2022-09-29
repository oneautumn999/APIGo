package controllers

import (
	"btpn/initializers"
	"btpn/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// rquest Untuk email dan password
	var body struct {
		Email    string
		Password string
		Username string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request tidak ditemukan",
		})

		return
	}

	// convert string menjadi hash untuk password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal Convert Password",
		})

		return
	}
	// membuat user
	user := models.User{Email: body.Email, Password: string(hash), Username: body.Username}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal Menambah User",
		})

		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	// mendapat email user
	var body struct {
		Email    string
		Password string
		Username string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request tidak ditemukan",
		})

		return
	}
	//req user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email Atau Password Salah",
		})

		return
	}

	//mengecek input user dan juga pass
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email Atau Password Salah",
		})

		return
	}
	//membuat jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal Membuat Token",
		})

		return
	}

	// respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func UserUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Email    string
		Password string
		Username string
	}

	c.Bind(&body)

	//menemukan user yang mau di update
	var user models.User
	initializers.DB.First(&user, id)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal Convert Password",
		})

		return
	}

	// memilih apa yang mau di update
	initializers.DB.Model(&user).Updates(models.User{
		Email:    body.Email,
		Password: string(hash),
		Username: body.Username,
	})

	c.JSON(200, gin.H{
		"message": user,
	})
}

func UserDelete(c *gin.Context) {
	//id user di database
	id := c.Param("id")

	//delete user
	initializers.DB.Delete(&models.User{}, id)

	//respond
	c.Status(200)
}
