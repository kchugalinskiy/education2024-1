package main

import (
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
}
