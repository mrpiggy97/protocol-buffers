package implementations

import (
	"fmt"

	"github.com/mrpiggy97/protocol-buffers/user"
	"google.golang.org/protobuf/proto"
)

// Encode will see how protofiles are encoded.
func Encode(group *user.Group) []byte {
	encodedGroup, encodingErr := proto.Marshal(group)
	if encodingErr != nil {
		panic(encodingErr)
	}
	fmt.Println(encodedGroup)
	return encodedGroup
}
