package helloworld

import (
	"fmt"

	hello "github.com/drymifolia/private-hello"
)

func HelloWorld() string {
	return fmt.Sprintf("%s world!", hello.Hello())
}
