package main

import (
	"fmt"
	"github.com/kchugalinskiy/education2024-1/internal/server"
	"github.com/kchugalinskiy/education2024-1/internal/wg"
	"github.com/kchugalinskiy/education2024-1/pkg/hello"
)

func main() {
	grp := wg.New()
	go func() {
		defer grp.Done()
		hello.Say()
	}()
	grp.Wait()

	srv := server.New(":8080")
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
