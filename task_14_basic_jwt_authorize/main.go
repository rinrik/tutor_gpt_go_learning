package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

const signature string = "Right signature"

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (cr *Credentials) Create() {
	cr.Login = "admin"
	cr.Password = "admin@"

}

type LoginResponse struct {
	Token string `json:"token"`
}

func (lr *LoginResponse) Marshal() []byte {
	a, err := json.Marshal(lr)
	if err != nil {
		log.Println("ERROR: Can't marshall jwt token")
		return nil
	}
	return a
}

func main() {

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/protected", protectedHandler)
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user Credentials
	var trueCreds Credentials
	trueCreds.Create()

	if r.Method != "POST" {
		log.Printf("Used method %v is not allowed for login", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Can't unmarshal request body")
		w.WriteHeader(http.StatusBadRequest)
	}

	if user.Login == trueCreds.Login && user.Password == trueCreds.Password {

		token := LoginResponse{
			Token: createJwt(user.Login, time.Now().Unix()),
		}

		w.WriteHeader(http.StatusOK)
		w.Write(token.Marshal())
	} else {
		log.Printf("Wrong credentials attempt!!!\n")
		w.WriteHeader(http.StatusUnauthorized)
	}

}

func createJwt(login string, reqTime int64) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: reqTime + 3600,
		NotBefore: reqTime - 3600,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "Ya",
		Id:        login,
	})

	signedToken, err := t.SignedString([]byte(signature))

	if err != nil {
		log.Println("Can't sign JWT token")
	}

	return signedToken
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Used method %v is not allowed for login", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	token, err := jwt.Parse(r.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		return []byte(signature), nil
	})

	if token.Method != jwt.SigningMethodHS256 {
		err = fmt.Errorf("signing method is wrong - %v", token.Method.Alg())
	}

	if !token.Valid {
		err = fmt.Errorf("invalid token - %v", token.Claims.Valid().Error())
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ERROR: " + err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
