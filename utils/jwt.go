package utils

import "github.com/dgrijalva/jwt-go"

// specify encryption key
var jwtSecret = []byte("2d9a0da267bee9c14d8e7aaedeca907c")

// Claim is the state and additional metadata of some entity (usually a user)
type Claims struct {
	UserId      int    `json:"id"`
	ClientId    string `json:"client_id"`
	AccessToken string `json:"access_token"`
	jwt.StandardClaims
}

// Get the Claims object information according to the incoming token value, (and then get the username and password in it)
func ParseToken(token string) (*Claims, error) {

	//It is used to parse the authentication statement. The internal method is mainly the specific decoding and verification process, and finally returns *Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// Obtain the Claims object from tokenClaims, and use the assertion to convert the object to our own defined Claims
		// To pass in pointers, the structures in the project are passed by pointers to save space.
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
