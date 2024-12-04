package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Age        int    `json:"age"`
}

type Users struct {
	mu   sync.Mutex
	list map[int]User
}

var users = Users{
	list: make(map[int]User),
}

var lastID int

func (us *Users) Create() {
	us.list = make(map[int]User)
}

func (us *Users) FindById(id int) (_ User, found bool) {
	if _, ok := us.list[id]; !ok {
		return User{}, ok
	}
	return us.list[id], true
}

func (us *Users) AddUser(user User, id int) {
	us.mu.Lock()
	defer us.mu.Unlock()
	us.list[id] = user
}

func (us *Users) UpdateUser(user User, id int) {
	us.mu.Lock()
	defer us.mu.Unlock()
	us.list[id] = user
}

// domain module ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

func PostUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	lastID++
	users.AddUser(user, lastID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": lastID})

}

func PutUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "user/")
	idStr := pathParts[1]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var user User
	if _, ok := users.FindById(id); !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	users.UpdateUser(user, id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users.list[id])

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "user/")
	idStr := pathParts[1]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, ok := users.FindById(id); !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users.list[id])
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "user/")
	idStr := pathParts[1]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, ok := users.FindById(id); !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(users.list, id)

	w.WriteHeader(http.StatusOK)
}

func UserRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUser(w, r)
	case http.MethodPut:
		PutUser(w, r)
	case http.MethodDelete:
		DeleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UserCreater(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		PostUser(w, r)
	}
}

// handlers module ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

func main() {

	http.HandleFunc("/user", UserCreater)
	http.HandleFunc("/user/", UserRouter)

	log.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

}
