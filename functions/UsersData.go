package groupietrackers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteUsersData(users []User) {
	file, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	_ = ioutil.WriteFile("server/data/users.json", file, 0644)
}

func GetUserData(username string) User{
	users := _GetUsers()

	for i, user := range users {
		if user.Username == username {	
			return users[i]
		}
	}
	return User{Username:""}
}

func SetUserData(user User) []User{
	users := _GetUsers()
	for i, userInL := range users {
		if userInL.Username == user.Username {	
			users[i] = user
			return users
		}
	}
	fmt.Println("Error User not found")
	return users

}

func AddUser(user User){
	users := _GetUsers()
	if GetUserData(user.Username).Username != ""{
		fmt.Println("Error, User Already exist !")
	} else {
		users = append(users, user)
	}
	WriteUsersData(users)
}

func _GetUsers()[]User{
	var users []User

	jsonFile, err := os.Open("server/data/users.json")

	if err != nil {
		fmt.Println(err)
	}
	jsonFileValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(jsonFileValue, &users)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	return users
}
