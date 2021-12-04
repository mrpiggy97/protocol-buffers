package implementations

import (
	"encoding/json"
	"fmt"

	"github.com/mrpiggy97/protocol-buffers/handleRecovery"
	"github.com/mrpiggy97/protocol-buffers/user"
	"google.golang.org/protobuf/proto"
)

// ProtoToJson will trasform a proto serialized object
// to json.
func ProtoToJson() {
	defer handleRecovery.RecoverBadEncoding()
	//first we will create a bunch of data
	var pythonDeveloper *user.Developer = &user.Developer{Language: "python"}
	var golangDeveloper *user.Developer = &user.Developer{Language: "golang"}
	var typescriptDeveloper *user.Developer = &user.Developer{Language: "typescript"}

	var pythonUser *user.User = &user.User{
		UserId: "33423432",
		Email:  "email@example.com",
		Type:   &user.User_Developer{Developer: pythonDeveloper},
	}

	var golangUser *user.User = &user.User{
		UserId: "thisis the id",
		Email:  "email@com",
		Type:   &user.User_Developer{Developer: golangDeveloper},
	}

	var typescriptUser *user.User = &user.User{
		UserId: "adsadsad",
		Email:  "emailexample.com",
		Type:   &user.User_Developer{Developer: typescriptDeveloper},
	}

	var userSlice []*user.User = []*user.User{pythonUser, golangUser, typescriptUser}
	var list *user.UserList = &user.UserList{
		Users: userSlice,
	}

	var teamMap map[string]*user.UserList = make(map[string]*user.UserList)
	teamMap["developers"] = list

	var firstTeam *user.Teams = &user.Teams{
		Team: teamMap,
	}

	//now we play with encodings and decodings
	firstTeamProtoEncoded, protoEncodingErr := proto.Marshal(firstTeam)
	if protoEncodingErr != nil {
		panic(protoEncodingErr)
	}

	var recoveredTeam *user.Teams = new(user.Teams)
	var decodingErr error = proto.Unmarshal(firstTeamProtoEncoded, recoveredTeam)
	if decodingErr != nil {
		panic(decodingErr)
	}

	teamAsJson, encodingErr := json.Marshal(recoveredTeam)
	if encodingErr != nil {
		panic(encodingErr)
	}

	fmt.Println("\n", string(teamAsJson))
}
