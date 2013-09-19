package main
import (
	"fmt"
	"os"
	"flag"
	"io"
	"bufio"
)

// read file, and eval Period Language
func read_fileline(path string) []string {
	var reader *bufio.Reader
	var lines []string

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader = bufio.NewReader(file)

	readline, err := reader.ReadString('\n')
	for err != io.EOF {
		lines = append(lines, readline)

		// check line error
		if err != nil && err != io.EOF {
			fmt.Println(err)
			os.Exit(1)
		}

		readline, err = reader.ReadString('\n')
	}
	return lines 
}


func main() {
	flag.Parse()
	argv := flag.Args()
	
	if len(argv) > 0 {
		lines := read_fileline(argv[0])
		fmt.Println(lines)
	}
}
