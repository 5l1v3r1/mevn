package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`^\[(.*)\]=(.*)$`)

func main() {
	cs := os.Getenv("CIRCLE_SHA1")
	if cs == "" {
		log.Print("Error: Export CIRCLE_SHA1")
		return
	}

	out := "export "
	messages := strings.Split(cs, " ")
	for _, message := range messages {
		matches := re.FindStringSubmatch(message)
		if matches != nil {
			out += fmt.Sprintf("%s=%s ", matches[1], matches[2])
		}
	}
	fmt.Print(out)
}
