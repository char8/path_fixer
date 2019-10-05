package main

import "log"
import "os"
import "strings"
import "fmt"

const (
	SetPathCommand = `PATH="%v"; export PATH;`
)

func main() {
	path := os.Getenv("PATH")
	extractedDirs := strings.Split(path, ":")
	sortedDirs, err := GetOrderedPaths(extractedDirs)
	if err != nil {
		log.Fatalf("Error parsing paths: %v", err)
	}
	output := fmt.Sprintf(SetPathCommand, strings.Join(sortedDirs, ":"))
	fmt.Println(output)
}
