package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	csvFileName := flag.String("csv file", "problems.csv","a csv file in 'question,answer' format")
	timeLimit := flag.Int("time limit", 30, "time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil{
		fmt.Printf("File %v could not be opened\n",*csvFileName)
		os.Exit(1)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		fmt.Printf("Error in parsing csv file")
	}

	pr := parseLine(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	ansChn := make(chan string)
	go func() {
		var ans string

		fmt.Scanf("%s\n",&ans)
		ansChn <- ans
	}()


	score := 0
	for i, p := range pr{

		fmt.Printf("Problem #%v: %v = \n", i+1, p.q)

		select {

		case <-timer.C:
			fmt.Printf("\nTime's up. Your score is %d",score)
			return
		case ans := <-ansChn:
			if (ans == p.a){
				fmt.Printf("Correct\n")
				score ++
			} else {
				fmt.Printf("Incorrect\n")
			}
		}

	}

	fmt.Printf("Your score is %d",score)

}

type problem struct {
	q string
	a string
}

func parseLine(lines [][]string) (ret []problem) {

	for _,v := range lines{

		var pr problem

		pr.q = v[0]
		pr.a = strings.TrimSpace(v[1])

		ret = append(ret, pr)
	}



	return
}