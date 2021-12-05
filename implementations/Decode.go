package implementations

import (
	"fmt"

	"github.com/mrpiggy97/protocol-buffers/user"
	"google.golang.org/protobuf/proto"
)

// Decode will see how profiles are decoded.

func Decode(encodedMessage []byte) {
	var recoveredMessage *user.User = new(user.User)
	var decodingErr error = proto.Unmarshal(encodedMessage, recoveredMessage)
	if decodingErr != nil {
		panic(decodingErr)
	}
	fmt.Println(recoveredMessage)
}
