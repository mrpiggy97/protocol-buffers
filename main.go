package main

import (
	"github.com/mrpiggy97/protocol-buffers/implementations"
	"github.com/mrpiggy97/protocol-buffers/user"
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

	var encodedUsersGroup []byte = implementations.Encode(usersGroup)
	implementations.Decode(encodedUsersGroup)
	implementations.ProtoMaps()
	implementations.ProtoToJson()
	implementations.EncodeGroup()
}
