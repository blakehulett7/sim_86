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
	fmt.Printf("es: %#x (%d)\n", chip.ES, chip.GetValue("es"))
	fmt.Printf("ss: %#x (%d)\n", chip.SS, chip.GetValue("ss"))
	fmt.Printf("ds: %#x (%d)\n", chip.DS, chip.GetValue("ds"))
	fmt.Println()
}

func Execute(chip *Chip, instruction Instruction) {
	dest := instruction.Dest
	if strings.Contains(instruction.Dest, "h") {
		dest = strings.Replace(instruction.Dest, "h", "x", 1)
	}
	if strings.Contains(instruction.Dest, "l") {
		dest = strings.Replace(instruction.Dest, "l", "x", 1)
	}

	fmt.Printf("%s:%#x->", dest, chip.GetValue(instruction.Dest))

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
