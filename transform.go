package main

import (
	"fmt"
	"os"
)

func (c *Chip) TransformSourceInt(source *int16, instruction Instruction) {
	switch instruction.Op {
	default:
		fmt.Println("Op not supported. Use mov, add, sub, or cmp")
		os.Exit(1)
	case "mov":
		return
	case "add":
	case "sub":
	case "cmp":
	}
}

func (c *Chip) TransformSourceRaw(source *[2]byte, instruction Instruction) {
	switch instruction.Op {
	default:
		fmt.Println("Op not supported. Use mov, add, sub, or cmp")
		os.Exit(1)
	case "mov":
		return
	case "add":
	case "sub":
	case "cmp":
	}
}

func (c *Chip) TransformSourceRegister(source *[2]byte, instruction Instruction) {
	switch instruction.Op {
	default:
		fmt.Println("Op not supported. Use mov, add, sub, or cmp")
		os.Exit(1)
	case "mov":
		return
	case "add":
	case "sub":
		dest := c.GetValue(instruction.Dest)
		src := castToU16(*source)
		result := dest - src
		*source = WriteUint(result)
	case "cmp":
	}
}
