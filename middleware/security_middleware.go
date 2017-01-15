package middleware

import (
	"os"
	"strings"
	"time"
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/wilsontamarozzi/panda-api/services"
	"github.com/wilsontamarozzi/panda-api/services/models"
)

const ENV_JWT_SECRET_KEY = "JWT_SECRET_KEY"

var JWT_SECRET_KEY = "panda"

func init() {
	getEnvJWTSecretKey()
}

func getEnvJWTSecretKey() {
	jwtSecretKey := os.Getenv(ENV_JWT_SECRET_KEY)

	if len(jwtSecretKey) > 0 {
		JWT_SECRET_KEY = jwtSecretKey
	}
}

type Claims struct {
	UserId 		string  	`json:"user_id"`
	Username 	string 		`json:"user_name"`
	jwt.StandardClaims
}

func SetToken(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)

	hasher := md5.New()
    hasher.Write([]byte(user.Password))
    
	person := services.AuthenticationUser(user.Username, hex.EncodeToString(hasher.Sum(nil)))

	if person != (models.Person{}) {

		expireToken := time.Now().Add(time.Hour * 1).Unix()
		expireCookie := time.Now().Add(time.Hour * 1)

		claims := Claims{
			person.UUID,
			person.Name,
			jwt.StandardClaims{
				ExpiresAt: expireToken,
				Issuer:    "api.panda",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		signedToken, _ := token.SignedString([]byte(JWT_SECRET_KEY))

		cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
    	http.SetCookie(c.Writer, &cookie)
    	c.Writer.Header().Set("Authorization", signedToken)

		c.JSON(200, gin.H{"token": signedToken, "user_id" : person.UUID})
	} else {
		c.JSON(401, gin.H{"errors": "Usuário ou senha invalidos."})
	}
}

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
	
		/* Pega o header Authorization */        
        tokenString := c.Request.Header.Get("Authorization")

        /* Quebra a String do Header */
        result := strings.Split(tokenString, "Bearer ")

        /* Analisa se não veio vazio ou faltou o Bearer */
	    if tokenString == "" || len(result) <= 1 {
	        c.AbortWithStatus(401)
	    } else if CheckToken(result[1], c) {
	        c.Next()
	    } else {
	        c.AbortWithStatus(401)
	    }
    }
}

func CheckToken(tokenString string, c *gin.Context) bool {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

    if err == nil && token.Valid {
    	user := token.Claims.(*Claims)

    	c.Set("userRequest", user.UserId)

        return true
    } else {
        return false
    }
}