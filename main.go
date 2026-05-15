package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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
	lines := strings.Split(string(file), "\n")

	fmt.Printf("--- test\\%s execution ---\n", strings.TrimSuffix(path, ".asm"))

	for chip.IP < uint8(len(lines)) {
		line := lines[chip.IP]
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
	fmt.Printf("ip: %d\n", chip.IP)
	fmt.Printf("flags: %s", chip.GetFlags())
	fmt.Println()
}

func Execute(chip *Chip, instruction Instruction) {
	if instruction.Op == "jnz" {
		if chip.Flags.Zero {
			chip.IP++
			return
		}

		dest, err := strconv.ParseUint(instruction.Dest, 10, 8)
		if err != nil {
			fmt.Println("invalid jnz destination")
			os.Exit(1)
		}
		fmt.Println()
		chip.IP = uint8(dest)
		return
	}

	dest := instruction.Dest
	if strings.Contains(instruction.Dest, "h") {
		dest = strings.Replace(instruction.Dest, "h", "x", 1)
	}
	if strings.Contains(instruction.Dest, "l") {
		dest = strings.Replace(instruction.Dest, "l", "x", 1)
	}

	if !strings.Contains(instruction.Dest, "[") {
		fmt.Printf("%s:%#x->", dest, chip.GetValue(instruction.Dest))
	}
	initial_flags := chip.GetFlags()
	flag_string := fmt.Sprintf("flags:%s", initial_flags)

	source_type := ParseSrcType(instruction.Src)
	chip.WriteSrcToDest(instruction, source_type)

	final_flags := chip.GetFlags()
	flag_string += fmt.Sprintf("->%s", final_flags)

	if !strings.Contains(instruction.Dest, "[") {
		fmt.Printf("%#x ", chip.GetValue(instruction.Dest))
	}
	if initial_flags != final_flags {
		fmt.Print(flag_string)
	}
	fmt.Printf(" ip: %d", chip.IP)
	fmt.Println()

	chip.IP++
}

func WriteInt(value int16) [2]byte {
	var b [2]byte
	b[0] = byte(value >> 8)
	b[1] = byte(value)

	return b
}

func WriteUint(value uint16) [2]byte {
	var b [2]byte
	b[0] = byte(value >> 8)
	b[1] = byte(value)

	return b
}
