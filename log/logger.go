package logs

import (
	. "fmt"
)

var D bool

func Debug(debug ...interface{}) {
	if D {
		Println("[DEBUG]", debug)
	}
}

func Error(err ...interface{}) {
	Println("[ERROR]", err)
}

func Info(info ...interface{}) {
	Println("[INFO]", info)
}
