package main

import (
	"fmt"
	"os"
	"time"
)

func division(a, b int) int {
	return a / b
}

type ErrorWithTime struct {
	text string
	time string
}

func New(text string) error {
	return &ErrorWithTime{
		text: text,
		time: fmt.Sprintf(time.Now().Format("01-02-2006 15:04:05"))}
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("%s  %s", e.time, e.text)
}

func CreateFile(name string) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = fmt.Fprintln(f, "")
}

func main() {
	defer func() {
		if v := recover(); v != nil {
			var err error
			err = New(fmt.Sprintf("%s", v))
			fmt.Println(err)
		}
	}()
	var a = 1
	var b int
	fmt.Println(division(a, b))

	for i := 1; i <= 1000; i++ {
		name := fmt.Sprintf("New_file%v.txt", i)
		CreateFile(name)
	}

}
