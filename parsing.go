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

	if immediate < -32768 || immediate > 65534 {
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

func ParseMemoryAddress(memory_lookup string) (MemoryAddress, error) {
	lookup := strings.Trim(memory_lookup, "[]")
	if !strings.Contains(lookup, "+") {
		address, err := strconv.Atoi(lookup)
		return MemoryAddress{Offset: address}, err
	}

	parts := strings.Split(lookup, "+")
	if len(parts) != 2 {
		return MemoryAddress{}, fmt.Errorf("invalid memory lookup: %s\n", memory_lookup)
	}
	offset, err := strconv.Atoi(parts[1])

	return MemoryAddress{Register: parts[0], Offset: offset}, err
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

	if isRegister(input) {
		return Register
	}

	_, err = hex.DecodeString(input)
	if err != nil {
		return Raw
	}

	return Register
}

func isRegister(input string) bool {
	return input == "ax" ||
		input == "bx" ||
		input == "cx" ||
		input == "dx" ||
		input == "sp" ||
		input == "bp" ||
		input == "si" ||
		input == "di" ||
		input == "cs" ||
		input == "ds" ||
		input == "ss" ||
		input == "es" ||
		input == "ah" ||
		input == "al" ||
		input == "bh" ||
		input == "bl" ||
		input == "ch" ||
		input == "cl" ||
		input == "dh" ||
		input == "dl"
}
