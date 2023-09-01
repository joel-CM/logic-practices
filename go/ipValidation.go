package main

import (
	"log"
	"regexp"
)

func IPValidation(ip string) bool {
	regex, err := regexp.Compile(`^([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)

	if err != nil {
		log.Fatal(err.Error())
	}

	return regex.Match([]byte(ip))
}
