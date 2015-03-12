package main

import "fmt"

type Status int

const (
	StatusOk Status = iota
	StatusError
)

func (s Status) String() string {
	return fmt.Sprintf("Status--%d", s)
}

func main() {
	fmt.Println(StatusError.String())
}
