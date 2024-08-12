package dirstructure

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintFilePaths() {
	// fmt.Println(filepath.Dir("/gomistakes/conditions.go"))

	p := filepath.Join("/", "/", "initial.go")
	fmt.Println(p)

	path, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Working directory: ", path)

	fmt.Println("Dir(p)", filepath.Dir(p))
	fmt.Println("Base(p)", filepath.Base(p))

	abs_path, err := filepath.Abs(p)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Absolute(p)", abs_path)
}
