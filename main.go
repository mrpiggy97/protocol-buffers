package main

import (
	"fmt"

	"github.com/mrpiggy97/protocol-buffers/user"
	"google.golang.org/protobuf/proto"
)

func main() {
	var userA *user.User = &user.User{UserId: "fabian123123", Email: "email@example.com"}
	var userB *user.User = &user.User{UserId: "ddfd111", Email: "another@example.com"}
	var userC *user.User = &user.User{UserId: "caasd222", Email: "example@email.com"}
	var userSlice []*user.User = []*user.User{userA, userB, userC}
	var usersGroup *user.Group = &user.Group{
		Id:       2323,
		Score:    12323.0,
		Category: user.Category_DEVELOPER,
		Users:    userSlice,
	}

	fmt.Println("to encode", usersGroup.String())
	encoded, err := proto.Marshal(usersGroup)
	if err != nil {
		panic(err)
	}
	var recoveredUsersGroup *user.Group = new(user.Group)
	var decodingErr error = proto.Unmarshal(encoded, recoveredUsersGroup)
	if decodingErr != nil {
		panic(decodingErr)
	}

	fmt.Println("\nrecovered:", recoveredUsersGroup.Users[0])
}
