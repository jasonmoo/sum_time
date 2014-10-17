package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	f = flag.Int("f", 0, "field to sum")
)

func main() {

	flag.Parse()

	if *f == 0 {
		fmt.Println("Usage: ")
		fmt.Println("./sum_time -f <num> < file.csv")
		flag.PrintDefaults()
		os.Exit(1)
	}

	cr := csv.NewReader(os.Stdin)
	sum := time.Duration(0)
	zero, _ := time.Parse("15:04:05", "00:00:00")

	for {
		line, err := cr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if len(line) < *f {
			fmt.Printf("Insufficient columns: %v\n", line)
			continue
		}
		t, err := time.Parse("15:04:05", line[*f-1])
		if err != nil {
			fmt.Printf("Unable to parse time: %s\n", line[*f-1])
			continue
		}
		sum += t.Sub(zero)
	}

	fmt.Printf("Total time: %s\n", sum)

}
