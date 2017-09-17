package middleware

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/Panda-CRM/panda-api/helpers"
	"github.com/Panda-CRM/panda-api/repositories"
	"time"
	"strings"
)

var (
	JWT_REALM       = "api.panda"
	JWT_SECRET_KEY  = "panda"
	MY_JWT          *jwt.GinJWTMiddleware
	REPOSITORY_USER repositories.UserRepository
)

func init() {
	MY_JWT = initJTW()
	REPOSITORY_USER = repositories.NewUserRepository()
}

func AuthRequired() gin.HandlerFunc {
	return MY_JWT.MiddlewareFunc()
}

func initJTW() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:         JWT_REALM,
		Key:           []byte(JWT_SECRET_KEY),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: Authenticator,
		Authorizator:  Authorization,
		Unauthorized:  Unauthorized,
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

type AccountController struct {
	Repository repositories.UserRepository
}

func (a AccountController) ForgotPassword(gc *gin.Context) {
	/*type params struct {
		id string `json:"id"`
	}
	var meu params
	if err := gc.BindJSON(&meu); err == nil {
		log.Printf("ID PESSOA: %s", meu.id)
		person := a.Repository.Get(meu.id)
		if !person.IsEmpty() {
			token := MY_JWT.TokenGenerator(person.UUID)
			var email helpers.ResetEmail
			email.UserName = person.Name
			email.ToEmails = []string{person.Email}
			email.ResetLink = fmt.Sprintf("https://panda.com/reset-password?token=%s", token)
			helpers.SMTP_SERVER.SendEmail(&email)
		}
		gc.JSON(200, nil)
	} else {
		gc.JSON(500, gin.H{"error": err.Error()})
	}*/
}

func (a AccountController) CheckTokenValid(gc *gin.Context) {
	claims := jwt.ExtractClaims(gc)
	userRequest := claims["id"].(string)
	exp := claims["exp"].(string)
	gc.JSON(200, gin.H{
		"user_id": userRequest,
		"exp":     exp,
	})
}

func (a AccountController) ResetPassword(gc *gin.Context) {

}

func Authenticator(userId string, password string, gc *gin.Context) (string, bool) {
	person := REPOSITORY_USER.Authenticator(userId, helpers.GetMD5Hash(password))
	if !person.IsEmpty() && person.Role.UUID != "" {
		return person.UUID, true
	}
	return userId, false
}

func Authorization(userId string, gc *gin.Context) bool {
	person := REPOSITORY_USER.Get(userId)
	rPath := gc.Request.URL.Path
	rMethod := gc.Request.Method
	for _, perm := range person.Role.Permissions {
		if rMethod == perm.Method {
			pathSplit := strings.Split(rPath, "/")
			permSplit := strings.Split(perm.Route, "/")
			if len(pathSplit) == len(permSplit) {
				var equals = 0
				for i := 0; i < len(pathSplit); i++ {
					if permSplit[i] == "*" || pathSplit[i] == permSplit[i] {
						equals++
					}
				}
				if equals == len(pathSplit) {
					return true
				}
			}
		}
	}
	return false
}

func Unauthorized(gc *gin.Context, code int, message string) {
	gc.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
