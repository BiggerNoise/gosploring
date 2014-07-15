package csv_test

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("CSV Operations", func() {
	It("Just load a dang file", func() {
		file, err := os.Open("input.tsv")
		// automatically call Close() at the end of current method
		defer file.Close()
		if err != nil {
			Fail(fmt.Sprintf("Unable to load file, bummer: %v\n", err.Error()))
		}
		reader := csv.NewReader(file)
		reader.Comma = '\t'
		count := 0
		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			if count++; count < 5 {
				fmt.Printf("Line: %v\n", line)
			}

			if count == 2 {
				fmt.Printf("Second line: %v, type %v", line[0], type.(line[0]))
			}
		}
	})
})
