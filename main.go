package ContainerFromScatch

import (
	"fmt"
	"os"
)

// docker run image <cmd> <params>
// go run main.go

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func run() {
	fmt.Println("Running %v\n", os.Args[2:])
}

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("bad command ")
	}

}
