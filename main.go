package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	file_name := flag.Arg(0)
	path := fmt.Sprintf("%s", file_name)

	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("could not open file")
		os.Exit(1)
	}

	var chip Chip

	fmt.Printf("--- test\\%s execution ---\n", strings.TrimSuffix(path, ".asm"))

	for line := range strings.SplitSeq(string(file), "\n") {
		fmt.Print(line + " ; ")
		instruction := ParseLine(line)
		Execute(&chip, instruction)
	}

	fmt.Println()
	fmt.Println("Final registers:")
	fmt.Printf("ax: %#x (%d)\n", chip.AX, chip.GetValue("ax"))
	fmt.Printf("bx: %#x (%d)\n", chip.BX, chip.GetValue("bx"))
	fmt.Printf("cx: %#x (%d)\n", chip.CX, chip.GetValue("cx"))
	fmt.Printf("dx: %#x (%d)\n", chip.DX, chip.GetValue("dx"))
	fmt.Printf("sp: %#x (%d)\n", chip.SP, chip.GetValue("sp"))
	fmt.Printf("bp: %#x (%d)\n", chip.BP, chip.GetValue("bp"))
	fmt.Printf("si: %#x (%d)\n", chip.SI, chip.GetValue("si"))
	fmt.Printf("di: %#x (%d)\n", chip.DI, chip.GetValue("di"))
	fmt.Println()
}

func Execute(chip *Chip, instruction Instruction) {
	fmt.Printf("%s:%#x->", instruction.Dest, chip.GetValue(instruction.Dest))

	if instruction.Op != "mov" {
		fmt.Println("\nOp not supported, use mov")
		os.Exit(1)
	}

	source_type := ParseSrcType(instruction.Src)
	chip.WriteSrcToDest(instruction, source_type)

	fmt.Printf("%#x\n", chip.GetValue(instruction.Dest))
}

func Write(value int16) [2]byte {
	var b [2]byte
	b[0] = byte(value >> 8)
	b[1] = byte(value)

	return b
}
