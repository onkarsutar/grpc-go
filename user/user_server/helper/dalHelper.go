package helper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/onkarsutar/grpc-go/user/userpb"
)

func FindUser(id int64) (userpb.User, error) {
	users := []userpb.User{}
	data, err := ioutil.ReadFile("./data/users.json")
	if err != nil {
		log.Printf("Error while reading file %v\n", err)
		return userpb.User{}, err
	}
	err = json.Unmarshal(data, &users)

	// for _, user := range users {
	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			return users[i], nil
		}
	}

	return userpb.User{}, errors.New("no user found")

}
