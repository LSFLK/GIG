package authentication

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"log"
)

func decodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secretKey, _ := revel.Config.String("app.secret")
		return []byte(secretKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println("email and nbf:", claims["email"], claims["nbf"])
	} else {
		log.Println(err)
		return nil, err
	}
	return claims, nil
}
