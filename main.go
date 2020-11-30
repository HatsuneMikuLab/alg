package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/HatsuneMikuLab/alg/zalg"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Answer: %v", zalg.CalcZValues(input))
}
