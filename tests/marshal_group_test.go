package tests

import (
	"testing"

	"github.com/mrpiggy97/protocol-buffers/user"
	"google.golang.org/protobuf/proto"
)

func TestMarshalGroup(testCase *testing.T) {
	var user1 *user.User = &user.User{
		UserId: "12323",
		Email:  "email@example.com",
	}

	var user2 *user.User = &user.User{
		UserId: "asdas23e",
		Email:  "example@email.com",
	}

	var group *user.Group = &user.Group{
		Id:       21312313,
		Category: user.Category_DEVELOPER,
		Score:    43434.012,
		Users:    []*user.User{user1, user2},
		Winner:   user1,
	}

	_, encodingErr := proto.Marshal(group)
	if encodingErr != nil {
		testCase.Error("encodingErr is supposed to be nil got", encodingErr)
	}
}
