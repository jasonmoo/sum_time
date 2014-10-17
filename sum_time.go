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

	for cr, sum := csv.NewReader(os.Stdin), time.Duration(0); ; {
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
		sum += time.Duration(t.Hour()) * time.Hour
		sum += time.Duration(t.Minute()) * time.Minute
		sum += time.Duration(t.Second()) * time.Second
	}

	fmt.Printf("Total time: %s\n", sum)

}
