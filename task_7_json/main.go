package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	var userList []string
	firstUser := user{
		Name:  "Alex",
		Age:   12,
		Email: "alex123@mail.com",
	}

	secondtUser := user{
		Name:  "Petya",
		Age:   21,
		Email: "Petr1@mail.com",
	}

	firstUser.writeToFile()
	secondtUser.writeToFile()

	userList = append(userList, firstUser.Email+`.txt`)
	userList = append(userList, secondtUser.Email+`.txt`)

	for _, fileName := range userList {
		newUser := user{}
		newUser.readFromFile(fileName)
		fmt.Println(newUser)
	}

}

func (u *user) toJson() (res []byte, err error) {
	res, err = json.Marshal(u)
	if err != nil {
		fmt.Printf("Ошибка в превращении в юзера в JSON - %v", err)
	}
	return res, nil
}

func (u *user) writeToFile() {
	fileName := u.Email + `.txt`
	fileData, _ := u.toJson()
	finalData := []byte(string(fileData))
	os.WriteFile(fileName, finalData, 0666)
}

func (u *user) readFromFile(fileName string) {
	bytedData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Ошибка: нельзя прочесть из файла %v", err)
	}
	json.Unmarshal(bytedData, u)
}
