package main

import (
	"fmt"
	"os"
	"strings"
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

	IP uint8

	Flags Flags

	Memory [2 << 20]byte
}

type Flags struct {
	Parity bool
	Sign   bool
	Zero   bool
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
	if strings.Contains(register, "[") {
		c.WriteToMemory(register, value)
		return
	}

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

func (c *Chip) GetFlags() string {
	var flags string

	if c.Flags.Parity {
		flags += "P"
	}

	if c.Flags.Sign {
		flags += "S"
	}

	if c.Flags.Zero {
		flags += "Z"
	}

	return flags
}

func (c *Chip) SetFlags(result uint16) {
	if result == 0 {
		c.Flags.Parity = true
		c.Flags.Sign = false
		c.Flags.Zero = true
		return
	}

	leading_bit := result >> 15
	if leading_bit == 1 {
		c.Flags.Parity = false
		c.Flags.Sign = true
		c.Flags.Zero = false
		return
	}

	c.Flags.Parity = false
	c.Flags.Sign = false
	c.Flags.Zero = false
}

func (c *Chip) WriteSrcToDest(instruction Instruction, source_type SourceType) {
	switch source_type {
	default:
	case Int:
		immediate, _ := ParseImmediate(instruction.Src)
		c.TransformSourceInt(&immediate, instruction)
		c.SetValue(instruction.Dest, WriteInt(immediate))
	case Memory:
		address, _ := ParseMemoryAddress(instruction.Src)
		value := c.GetFromMemory(address)
		c.SetValue(instruction.Dest, value)
	case Raw:
		value := ParseRaw(instruction.Src)
		c.TransformSourceRaw(&value, instruction)
		c.SetValue(instruction.Dest, value)
	case Register:
		src_value := c.GetRawValue(instruction.Src)
		c.TransformSourceRegister(&src_value, instruction)
		c.SetValue(instruction.Dest, src_value)
	}
}

func (c *Chip) WriteToMemory(memory_lookup string, value [2]byte) {
	memory_address, err := ParseMemoryAddress(memory_lookup)
	if err != nil {
		fmt.Println("invalid memory address")
		os.Exit(1)
	}

	if memory_address.Register == "" {
		c.Memory[memory_address.Offset+1] = value[0]
		c.Memory[memory_address.Offset] = value[1]
		return
	}

	index := int(c.GetValue(memory_address.Register)) + memory_address.Offset
	if memory_address.OffsetRegister != "" {
		index += int(c.GetValue(memory_address.OffsetRegister))
	}

	c.Memory[index+1] = value[0]
	c.Memory[index+0] = value[1]
}

func (c *Chip) GetFromMemory(address MemoryAddress) [2]byte {
	if address.Register == "" {
		return [2]byte{c.Memory[address.Offset+1], c.Memory[address.Offset]}
	}

	index := int(c.GetValue(address.Register)) + address.Offset
	if address.OffsetRegister != "" {
		index += int(c.GetValue(address.OffsetRegister))
	}

	return [2]byte{c.Memory[index+1], c.Memory[index]}
}

func NewChip() Chip {
	return Chip{}
}

func castToU16(data [2]byte) uint16 {
	return uint16(uint16(data[0])<<8 | uint16(data[1]))
}
