package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lsflk/gig-sdk/models"
	"github.com/revel/revel"
	"time"
)

func CreateSignedUserToken(user models.User) (models.UserToken, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Subject:   user.Role,
	})

	secretKey, _ := revel.Config.String("app.secret")

	token, err := claims.SignedString([]byte(secretKey))

	return models.UserToken{Name: user.Name, Email: user.Email, Role: user.Role, Token: token}, err
}
