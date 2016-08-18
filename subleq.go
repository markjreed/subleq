package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// run a SUBLEQ program
func RunProg(memory []int) error {
	pc := 0
	reader := bufio.NewReader(os.Stdin)
	for {
		if pc < 0 {
			return nil
		}
		if pc > len(memory)-3 {
			return fmt.Errorf("Address out of range: %d", pc)
		}
		a, b, c := memory[pc], memory[pc+1], memory[pc+2]
		pc += 3
		if b < 0 {
			_, err := fmt.Printf("%c", memory[a])
			if err != nil {
				panic(err)
				return err
			}
		} else if a < 0 {
			char, _, err := reader.ReadRune()
			if err != nil {
				panic(err)
				return err
			}
			memory[b] = int(char)
		} else {
			memory[b] -= memory[a]
			if memory[b] <= 0 {
				if c == 27 || c == 30 {
				}
				pc = c
			}
		}
	}
}

func main() {
	var reader io.Reader
	var err error
	var handle *os.File
	closeFlag := false
	if len(os.Args) > 1 {
		handle, err = os.Open(os.Args[1])
		check(err)
		closeFlag = true
		reader = handle
	} else {
		reader = bufio.NewReader(os.Stdin)
	}
	memory, err := ReadInts(reader)
	if closeFlag {
		handle.Close()
	}
	check(err)
	RunProg(memory)
}
