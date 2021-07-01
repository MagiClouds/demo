package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)


/*
iss(Issuser)：代表这个JWT的签发主体；
sub(Subject)：代表这个JWT的主体，即它的所有人；
aud(Audience)：代表这个JWT的接收对象；
exp(Expiration time)：是一个时间戳，代表这个JWT的过期时间；
nbf(Not Before)：是一个时间戳，代表这个JWT生效的开始时间，意味着在这个时间之前验证JWT是会失败的；
iat(Issued at)：是一个时间戳，代表这个JWT的签发时间；
jti(JWT ID)：是JWT的唯一标识。
*/

func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "lora-app-server",
		"aud": "lora-app-server",
		"nbf": time.Now().Unix(),
		"exp": time.Date(2021, 11, 1, 0, 0, 0, 0, time.UTC).Unix(),
		"sub": "user",
		"username": "admin",
	})

	tokenString, err := token.SignedString([]byte("verysecret"))
	fmt.Println(tokenString, err)
}
