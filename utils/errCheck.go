package utils

import "log"

func ErrorCheck(err error, arg string) {
	if err != nil {
		switch arg {
		case "Print":
			log.Println(err)
		case "Println":
			log.Println(err)
		case "Fatal":
			log.Fatal(err)
		case "Panic":
			log.Panic(err)
		default:
			log.Println("Method not implemented")
		}
	}
}
