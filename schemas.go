package main

type SourceType uint8

const (
	Int SourceType = iota
	Raw
	Register
	SmallRegister
)
