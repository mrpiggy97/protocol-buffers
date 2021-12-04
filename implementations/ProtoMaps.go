package implementations

import (
	"github.com/mrpiggy97/protocol-buffers/handleRecovery"
	"github.com/mrpiggy97/protocol-buffers/user"
	"google.golang.org/protobuf/proto"
)

// ProtoMaps will show how proto maps work.
func ProtoMaps() {
	defer handleRecovery.RecoverBadEncoding()
	var userA *user.User = &user.User{UserId: "fabian123123", Email: "email@example.com"}
	var userB *user.User = &user.User{UserId: "ddfd111", Email: "another@example.com"}
	var userC *user.User = &user.User{UserId: "caasd222", Email: "example@email.com"}
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

	var userSlice []*user.User = []*user.User{userA, userB, userC, userD, userE}

	var list *user.UserList = &user.UserList{Users: userSlice}
	var team map[string]*user.UserList = make(map[string]*user.UserList)
	team["team1"] = list

	var teamMap *user.Teams = &user.Teams{Team: team}

	var recoveredTeamMap *user.Teams = new(user.Teams)

	encodedTeamMap, encodingErr := proto.Marshal(teamMap)
	if encodingErr != nil {
		panic(encodingErr)
	}

	var decodingErr error = proto.Unmarshal(encodedTeamMap, recoveredTeamMap)
	if decodingErr != nil {
		panic(decodingErr)
	}
}
