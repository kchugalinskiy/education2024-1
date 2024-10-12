package hello

import "fmt"

func privateSay() {
	fmt.Println("Hail robots")
}

func Say() {
	defer privateSay()
	fmt.Println("Hello, World!")
}
