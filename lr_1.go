package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("===Start===")

	var (
		inputFileName  = flag.String("i", "", "Use a file with the name file-name as an input.")
		outputFileName = flag.String("o", "", "Use a file with the name file-name as an output.")
		ignoreHeader   = flag.Bool("h", false, "The first line is a header that must be ignored during sorting but included in the output.")
		sortingField   = flag.Int("f", 0, "Sort input lines by value number N.")
		reverseSort    = flag.Bool("r", false, "Sort input lines in reverse order.")
	)

	flag.Parse()

	if outputFileName != nil && *outputFileName != "" {
		content := readFile(*outputFileName)
		sortContent(content, *ignoreHeader, *sortingField, *reverseSort)
		fmt.Println(content)
	}

	if inputFileName != nil && *inputFileName != "" {
		writeFile(*inputFileName)
	}

	fmt.Println("===Finish===")
}

func readFile(name string) [][]string {
	f, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	content := [][]string{}
	n := 0

	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		row := strings.Split(line, ",")
		if n == 0 {
			n = len(row)
		}
		if n != len(row) {
			log.Fatalf("ERROR: row has %d columns, but must have %d\n", len(row), n)
		}
		content = append(content, row)
	}

	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	return content
}

func writeFile(name string) {
	s := bufio.NewScanner(os.Stdin)

	n := 0

	f, err := os.Create(name)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		row := strings.Split(line, ",")

		if n == 0 {
			n = len(row)
		}

		if n != len(row) {
			fmt.Printf("Error: row has %d column, but must have %d\n", len(row), n)
			os.Exit(1)
		}

		_, err2 := f.WriteString(line + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
	}
	fmt.Println("done")
}

func sortContent(content [][]string, header bool, sortField int, rev bool) {
	if rev {
		if header {
			sort.Slice(content[1:], func(i, j int) bool {
				return content[1:][i][sortField] > content[1:][j][sortField]
			})
		} else {
			sort.Slice(content, func(i, j int) bool {
				return content[i][sortField] > content[j][sortField]
			})
		}
	} else {
		if header {
			sort.Slice(content[1:], func(i, j int) bool {
				return content[1:][i][sortField] < content[1:][j][sortField]
			})
		} else {
			sort.Slice(content, func(i, j int) bool {
				return content[i][sortField] < content[j][sortField]
			})
		}
	}

}
