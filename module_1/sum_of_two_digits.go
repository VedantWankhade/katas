package module1

import (
	"os"
	"text/scanner"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	in := scanner.Scanner{}
	in.Init(os.Stdout)

}
