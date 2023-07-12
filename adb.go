package adb

import "fmt"

func connect(name string) string{
	message := fmt.Sprintf("Hi, %v. adb connect running ", name)
	return message
}