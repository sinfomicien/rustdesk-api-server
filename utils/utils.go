package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJwtToken(userId int, username, token, clientId, uuid string) (string, error) {
	//Set token valid time
	nowTime := time.Now()
	expireTime := nowTime.Add(399999 * time.Hour)

	claims := Claims{
		UserId: userId,
		//Username: username,
		//Password: password,
		AccessToken: token,
		ClientId:    clientId,
		//Uuid:     uuid,
		StandardClaims: jwt.StandardClaims{
			// Expiration
			ExpiresAt: expireTime.Unix(),
			// Specify the token issuer
			Issuer: "baozier",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//This method internally generates a signature string, which is then used to obtain a complete, signed token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
