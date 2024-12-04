package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Email string
	Name  string
}

func main() {
	http.HandleFunc("/", helloHandlerFunc)
	http.HandleFunc("/user", showUserHandlerFunc)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("server can't start")
	}
}

func helloHandlerFunc(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("Hello"))
	} else {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Only GET requests are accepted"))
	}
}

func showUserHandlerFunc(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		user := User{
			Email: "helloUser@mail.com",
			Name:  "User1",
		}

		marshalledUser, err := user.toJson()
		if err != nil {
			resp.Write([]byte("Error to convert to JSON"))
		}

		resp.Write(marshalledUser)

	} else {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Only GET requests are accepted"))
	}
}

func (u *User) toJson() ([]byte, error) {
	jsonUser, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return jsonUser, nil
}
