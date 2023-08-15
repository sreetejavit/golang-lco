package main

import "fmt"

type Logger struct {
	Filename string
}

var instance *Logger

func (l *Logger) GetFileName() string {
	return l.Filename
}

func GetLogInstance() *Logger {

	if instance == nil {
		instance = &Logger{Filename: "logfile.log"}

	} else {
		fmt.Println("There is already logger instance")
	}

	return instance

}

func Recoverfunc() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}

func main() {
	fmt.Println("Create singletone Logging instance")
	Logger := GetLogInstance()
	fmt.Println(Logger.GetFileName())

	Logger2 := GetLogInstance()
	fmt.Println(Logger2.GetFileName())
	Logger23 := GetLogInstance()
	fmt.Println(Logger23.GetFileName())
}
