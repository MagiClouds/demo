package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main()  {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJsb3JhLWFwcC1zZXJ2ZXIiLCJleHAiOjE2MzU3MjQ4MDAsImlzcyI6ImxvcmEtYXBwLXNlcnZlciIsIm5iZiI6MTYyNTA0NzMwNiwic3ViIjoidXNlciIsInVzZXJuYW1lIjoiYWRtaW4ifQ.6J56Y6B1mV0jiaT0fhmbgKhIeyH94rpg9WWPXVcF2LQ"
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("verysecret"), nil
	})

	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["iss"], claims["sub"])
	}

}
