package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseImmediate(input string) (int16, error) {
	immediate, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("not an immediate")
	}

	if immediate < -32768 || immediate > 32767 {
		fmt.Println("immediate overflow")
		os.Exit(1)
	}

	return int16(immediate), nil
}

func ParseLine(line string) Instruction {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		fmt.Printf("invalid instruction: %s\n", line)
		os.Exit(1)
	}

	return Instruction{
		Op:   parts[0],
		Dest: strings.TrimSuffix(parts[1], ","),
		Src:  parts[2],
	}
}
