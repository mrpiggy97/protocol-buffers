package main

import (
	"fmt"

	"github.com/mrpiggy97/protocol-buffers/implementations"
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

	fmt.Println("to encode", usersGroup)
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

	var golangDeveloper *user.Developer = new(user.Developer)
	golangDeveloper = &user.Developer{Language: "golang"}
	var userType *user.User_Developer = new(user.User_Developer)
	userType = &user.User_Developer{Developer: golangDeveloper}
	var userD *user.User = new(user.User)
	userD = &user.User{
		UserId: "122323",
		Email:  "exampleemail@com",
		Type:   userType,
	}

	var aksOperator *user.Operator = new(user.Operator)
	aksOperator = &user.Operator{Platform: "aks"}

	var operatorType *user.User_Operator = new(user.User_Operator)
	operatorType = &user.User_Operator{Operator: aksOperator}

	var userE *user.User = &user.User{
		UserId: "this is the user e",
		Email:  "this@email.com",
		Type:   operatorType,
	}

	var secondUserGroup *user.Group = &user.Group{
		Id:       123213123,
		Category: user.Category_OPERATOR,
		Score:    123123123.13123,
		Users:    []*user.User{userD, userE},
		Winner:   userE,
	}

	var recoveredSecondGroup *user.Group = new(user.Group)

	encodedSecondGroup, marshalingErr := proto.Marshal(secondUserGroup)

	if marshalingErr != nil {
		panic(marshalingErr)
	}

	var recoveringErr error = proto.Unmarshal(encodedSecondGroup, recoveredSecondGroup)
	if recoveringErr != nil {
		panic(recoveringErr)
	}

	fmt.Println("\n recovered group:", recoveredSecondGroup)

	implementations.ProtoMaps()
	implementations.ProtoToJson()
}
