package handleRecovery

import "fmt"

func RecoverBadEncoding() {
	r := recover()
	if r != nil {
		fmt.Println("recovered from encoding error", r)
	}
}
