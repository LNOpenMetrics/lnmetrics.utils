package goutils

import "fmt"

func Todo(message string, args ...interface{}) error {
	return fmt.Errorf(message, args...)
}

func Errf(message string, args ...interface{}) error {
	return fmt.Errorf(message, args...)
}

func NotImplementYet() error {
	return Todo("not implemented yet")
}
