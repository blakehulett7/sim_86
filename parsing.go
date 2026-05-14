package main

import (
	"encoding/hex"
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

func ParseRaw(input string) [2]byte {
	data, _ := hex.DecodeString(strings.TrimPrefix(input, "0x"))

	if len(data) > 2 || len(data) == 0 {
		fmt.Printf("invalid raw source, to big or empty: %s\n", input)
		os.Exit(1)
	}

	if len(data) == 1 {
		return [2]byte{0, data[0]}
	}

	return [2]byte{data[0], data[1]}
}

func ParseSrcType(input string) SourceType {
	_, err := strconv.Atoi(input)
	if err == nil {
		return Int
	}

	_, err = hex.DecodeString(input)
	if err != nil {
		return Raw
	}

	return Register
}
