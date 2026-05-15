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
	CS [2]byte
	DS [2]byte
	SS [2]byte
	ES [2]byte
}

func (c *Chip) GetRawValue(register string) [2]byte {
	switch register {
	default:
		fmt.Println("invalid register, can't get raw")
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
	case "cs":
		return c.CS
	case "ds":
		return c.DS
	case "ss":
		return c.SS
	case "es":
		return c.ES

	case "ah":
		return [2]byte{0, c.AX[0]}
	case "al":
		return [2]byte{0, c.AX[1]}
	case "bh":
		return [2]byte{0, c.BX[0]}
	case "bl":
		return [2]byte{0, c.BX[1]}
	case "ch":
		return [2]byte{0, c.CX[0]}
	case "cl":
		return [2]byte{0, c.CX[1]}
	case "dh":
		return [2]byte{0, c.DX[0]}
	case "dl":
		return [2]byte{0, c.DX[1]}
	}
}

func (c *Chip) GetValue(register string) uint16 {
	switch register {
	default:
		fmt.Println("invalid register, can't get")
		os.Exit(1)
		return 0

	case "ax", "ah", "al":
		return castToU16(c.AX)
	case "bx", "bh", "bl":
		return castToU16(c.BX)
	case "cx", "ch", "cl":
		return castToU16(c.CX)
	case "dx", "dh", "dl":
		return castToU16(c.DX)
	case "sp":
		return castToU16(c.SP)
	case "bp":
		return castToU16(c.BP)
	case "si":
		return castToU16(c.SI)
	case "di":
		return castToU16(c.DI)

	case "cs":
		return castToU16(c.CS)
	case "ds":
		return castToU16(c.DS)
	case "ss":
		return castToU16(c.SS)
	case "es":
		return castToU16(c.ES)
	}
}

func (c *Chip) SetValue(register string, value [2]byte) {
	switch register {
	default:
		fmt.Println("invalid register, can't set")
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
	case "cs":
		c.CS = value
	case "ds":
		c.DS = value
	case "ss":
		c.SS = value
	case "es":
		c.ES = value

	case "ah":
		c.AX[0] = value[1]
	case "al":
		c.AX[1] = value[1]
	case "bh":
		c.BX[0] = value[1]
	case "bl":
		c.BX[1] = value[1]
	case "ch":
		c.CX[0] = value[1]
	case "cl":
		c.CX[1] = value[1]
	case "dh":
		c.DX[0] = value[1]
	case "dl":
		c.DX[1] = value[1]
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

func castToU16(data [2]byte) uint16 {
	return uint16(uint16(data[0])<<8 | uint16(data[1]))
}
