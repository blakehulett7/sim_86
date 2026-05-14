package main

import (
	"fmt"
	"os"
)

type Chip struct {
	AX [2]byte
	BX [2]byte
	CX [2]byte
	DX [2]byte
	SP [2]byte
	BP [2]byte
	SI [2]byte
	DI [2]byte
}

func (c *Chip) GetRawValue(register string) [2]byte {
	switch register {
	default:
		fmt.Println("invalid register")
		os.Exit(1)
		return [2]byte{}

	case "ax":
		return c.AX
	case "bx":
		return c.BX
	case "cx":
		return c.CX
	case "dx":
		return c.DX
	case "sp":
		return c.SP
	case "bp":
		return c.BP
	case "si":
		return c.SI
	case "di":
		return c.DI
	}
}

func (c *Chip) GetValue(register string) int16 {
	switch register {
	default:
		fmt.Println("invalid register")
		os.Exit(1)
		return 0

	case "ax":
		return castToI16(c.AX)
	case "bx":
		return castToI16(c.BX)
	case "cx":
		return castToI16(c.CX)
	case "dx":
		return castToI16(c.DX)
	case "sp":
		return castToI16(c.SP)
	case "bp":
		return castToI16(c.BP)
	case "si":
		return castToI16(c.SI)
	case "di":
		return castToI16(c.DI)
	}
}

func (c *Chip) SetValue(register string, value [2]byte) {
	switch register {
	default:
		fmt.Println("invalid register")
		os.Exit(1)
		return

	case "ax":
		c.AX = value
	case "bx":
		c.BX = value
	case "cx":
		c.CX = value
	case "dx":
		c.DX = value
	case "sp":
		c.SP = value
	case "bp":
		c.BP = value
	case "si":
		c.SI = value
	case "di":
		c.DI = value
	}
}

func (c *Chip) WriteSrcToDest(instruction Instruction, source_type SourceType) {
	switch source_type {
	default:
	case Int:
		immediate, _ := ParseImmediate(instruction.Src)
		c.SetValue(instruction.Dest, Write(immediate))
	case Raw:
		value := ParseRaw(instruction.Src)
		c.SetValue(instruction.Dest, value)
	case Register:
		src_value := c.GetRawValue(instruction.Src)
		c.SetValue(instruction.Dest, src_value)
	}
}

func NewChip() Chip {
	return Chip{}
}

func castToI16(data [2]byte) int16 {
	return int16(uint16(data[0])<<8 | uint16(data[1]))
}
