package sinonimo

import (
	"fmt"
	"log"
	"os"
)

func checkIfError(err error) {
	if err == nil {
		return
	}

	log.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
