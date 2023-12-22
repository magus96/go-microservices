package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var signInKey = []byte(os.Getenv("MY_SECRET_KEY"))

func index(w http.ResponseWriter, r *http.Request) {
	validToken, err := getJWT()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("Token Generation Failed: ")
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(validToken))
}

func getJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "anirudh"
	claims["aud"] = "aa.io"
	claims["iss"] = "bb.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(signInKey)
	if err != nil {
		fmt.Errorf("Something went wrong %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func handleRequests() {

	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println(signInKey)
	handleRequests()
}
