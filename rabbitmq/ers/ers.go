package ers

import (
	"fmt"
	"log"
)

func FailOnErrors(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
