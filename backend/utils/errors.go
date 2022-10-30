package utils

import "log"

func HandleWarning(err error) {
	if err != nil {
		log.Println(err)
	}
}

func HandleFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
