package server

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type Users struct {
	Users []*User `json:"Users"`
}

func (t *User) Create(user *User) error {
	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return err
	}

	defer file.Close()
	byteValue, _ := io.ReadAll(file)

	user.Id = uuid.New().ClockSequence()

	if err != nil {
		return err
	}

	err = file.Truncate(0)
	file.Seek(0, 0)

	if err != nil {
		fmt.Println("Error during file trunc:", err)
	}
	var UsersData Users
	if len(byteValue) > 0 {
		err = json.Unmarshal(byteValue, &UsersData)
		if err != nil {
			fmt.Println("Error unmarshaling JSON: ", err)
		}
	}

	var appUsers []*User
	for _, value := range UsersData.Users {
		appUsers = append(appUsers, &User{Id: value.Id, Name: value.Name, Email: value.Email, Phone: value.Phone})
	}
	appUsers = append(appUsers, &User{Id: user.Id, Name: user.Name, Email: user.Email, Phone: user.Phone})
	jsonData, err := json.MarshalIndent(Users{appUsers}, "", "\t")

	if err != nil {
		fmt.Println("Error serializing object: ", err)
		return err
	}

	_, err = file.Write(jsonData)

	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}
	return nil
}
func (t *User) Read(id int) (interface{}, error) {
	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, err
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	var users Users
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data: ", err)
		return nil, err
	}
	user := findUser(users.Users, id)
	if user != nil {
		fmt.Println("User found: ")
		fmt.Println("User ID: ", user.Id)
		fmt.Println("User name: ", user.Name)
		fmt.Println("User email: ", user.Email)
		fmt.Println("User phone: ", user.Phone)
		return user, nil
	} else {
		fmt.Println("User not found")
		return nil, err
	}
}
func (t *User) ReadAll(display bool) (interface{}, error) {
	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, err
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	var users Users
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println("Error during unmarshalling JSON: ", err)
	}
	if display {
		for i := 0; i < len(users.Users); i++ {
			user := users.Users[i]
			fmt.Println("--------------------")
			fmt.Println("User ID: ", user.Id)
			fmt.Println("User name: ", user.Name)
			fmt.Println("User email: ", user.Email)
			fmt.Println("User phone: ", user.Phone)
		}
	}
	return users, err
}
func (t *User) Update(id int, data *User) error {
	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return err
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)

	var users Users

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data: ", err)
		return err
	}

	user := findUser(users.Users, id)
	if user == nil {
		fmt.Println("User not found")
		return nil
	}
	if data.Name != "" {
		user.Name = data.Name
	}

	if data.Email != "" {
		user.Email = data.Email
	}

	if data.Phone != "" {
		user.Phone = data.Phone
	}

	jsonData, err := json.MarshalIndent(Users{users.Users}, "", "\t")

	if err != nil {
		fmt.Println("Error serializing object: ", err)
		return err
	}
	err = file.Truncate(0)
	file.Seek(0, 0)
	_, err = file.Write(jsonData)

	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}
	return nil
}
func (t *User) Delete(id int) error {
	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	byteValue, _ := io.ReadAll(file)
	var users Users

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data: ", err)
		return err
	}

	user := findUser(users.Users, id)
	if user == nil {
		fmt.Println("User not found")
		return nil
	}
	var updatedUsers []*User
	for _, user := range users.Users {
		if user.Id == id {
			continue
		} else {
			updatedUsers = append(updatedUsers, &User{Id: user.Id, Name: user.Name, Email: user.Email, Phone: user.Phone})
		}
	}

	jsonData, err := json.MarshalIndent(Users{updatedUsers}, "", "\t")

	if err != nil {
		fmt.Println("Error serializing object: ", err)
		return err
	}

	_, err = file.Write(jsonData)

	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}
	return nil
}

func findUser(users []*User, id int) *User {
	for _, user := range users {
		if user.Id == id {
			return user
		}
	}
	return nil
}
func checkIfUserExists(id int) bool {
	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return false
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	var users Users
	if len(byteValue) > 0 {
		err = json.Unmarshal(byteValue, &users)
		if err != nil {
			fmt.Println("Error unmarshaling JSON: ", err)
		}
	}
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Id == id {
			return true
		}
	}
	return false
}
