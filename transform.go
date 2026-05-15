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
		dest := c.GetValue(instruction.Dest)
		result := int16(dest) + *source
		c.SetFlags(uint16(result))
		*source = result
	case "sub":
		dest := c.GetValue(instruction.Dest)
		result := int16(dest) - *source
		c.SetFlags(uint16(result))
		*source = result
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
		dest := c.GetValue(instruction.Dest)
		src := castToU16(*source)
		result := dest + src
		c.SetFlags(result)
		*source = WriteUint(result)
	case "sub":
		dest := c.GetValue(instruction.Dest)
		src := castToU16(*source)
		result := dest - src
		c.SetFlags(result)
		*source = WriteUint(result)
	case "cmp":
		dest := c.GetValue(instruction.Dest)
		src := castToU16(*source)
		result := dest - src
		c.SetFlags(result)
		*source = WriteUint(dest)
	}
}
