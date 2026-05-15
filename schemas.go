package main

type MemoryAddress struct {
	Register       string
	Offset         int
	OffsetRegister string
}

type SourceType uint8

const (
	Int SourceType = iota
	Memory
	Raw
	Register
	SmallRegister
)
