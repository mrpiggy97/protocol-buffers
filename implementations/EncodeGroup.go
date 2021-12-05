package implementations

import (
	"fmt"

	"github.com/mrpiggy97/protocol-buffers/user"
	"google.golang.org/protobuf/proto"
)

// EncodeGroup will serialize a group.
func EncodeGroup() {
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
	fmt.Println(recoveredSecondGroup.String())
}
