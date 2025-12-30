package main

import (
	"fmt"
	"os"
)

const mov_op_code = 0b10001000

const reg_field_mask = 0b00111000

const mod_memory_mode = 0b00
const mod_memory_8bit = 0b01
const mod_memory_16bit = 0b10
const mod_reg_mode = 0b11

const reg_AL_AX = 0b00000000
const reg_CL_CX = 0b00001000
const reg_DL_DX = 0b00010000
const reg_BL_BX = 0b00011000
const reg_AH_SP = 0b00100000
const reg_CH_BP = 0b00101000
const reg_DH_SI = 0b00110000
const reg_BH_DI = 0b00111000

const rm_AL_AX = 0b00000000
const rm_CL_CX = 0b00000001
const rm_DL_DX = 0b00000010
const rm_BL_BX = 0b00000011
const rm_AH_SP = 0b00000100
const rm_CH_BP = 0b00000101
const rm_DH_SI = 0b00000110
const rm_BH_DI = 0b00000111

func main() {

	args := os.Args[1:]
	file, error := os.Open(args[0])
	check(error)
	b := make([]byte, 2)
	_, error = file.Read(b)
	check(error)

	// b7 := (b[0] & 0b10000000) >> 7
	// b6 := (b[0] & 0b01000000) >> 6
	// b5 := (b[0] & 0b00100000) >> 5
	// b4 := (b[0] & 0b00010000) >> 4
	// b3 := (b[0] & 0b00001000) >> 3
	// b2 := (b[0] & 0b00000100) >> 2
	// b1 := (b[0] & 0b00000010) >> 1
	// b0 := (b[0] & 0b00000001) >> 0

	// b15 := (b[1] & 0b10000000) >> 7
	// b14 := (b[1] & 0b01000000) >> 6
	// b13 := (b[1] & 0b00100000) >> 5
	// b12 := (b[1] & 0b00010000) >> 4
	// b11 := (b[1] & 0b00001000) >> 3
	// b10 := (b[1] & 0b00000100) >> 2
	// b09 := (b[1] & 0b00000010) >> 1
	// b08 := (b[1] & 0b00000001) >> 0

	opcode := (b[0] & 0b11111100)
	// direction := (b[0] & 0b00000010)
	wordop := (b[0] & 0b00000001)
	// mod := (b[1] & 0b11000000)
	reg := (b[1] & reg_field_mask)
	rm := (b[1] & 0b00000111)
	reg_addr := eval_register_name(reg, wordop)
	rm_addr := eval_rm_name(rm, wordop)

	if opcode == mov_op_code {
		fmt.Printf("mov %s, %s", reg_addr, rm_addr)
	}

	//////////////////////////////////
	// fmt.Printf("%b", b7)	        //
	// fmt.Printf("%b", b6)	        //
	// fmt.Printf("%b", b5)	        //
	// fmt.Printf("%b", b4)	        //
	// fmt.Printf("%b", b3)	        //
	// fmt.Printf("%b", b2)	        //
	// fmt.Printf("%b", b1)	        //
	// fmt.Printf("%b", b0)	        //
	// 			        //
	// fmt.Print(" ")	        //
	// 			        //
	// fmt.Printf("%b", b15)        //
	// fmt.Printf("%b", b14)        //
	// fmt.Printf("%b", b13)        //
	// fmt.Printf("%b", b12)        //
	// fmt.Printf("%b", b11)        //
	// fmt.Printf("%b", b10)        //
	// fmt.Printf("%b", b09)        //
	// fmt.Printf("%b", b08)        //
	//////////////////////////////////

	///////////////////////////////////////////
	// for _, v := range b {		 //
	// 	fmt.Printf("%08b ", v)		 //
	// }					 //
	// fmt.Println()			 //
	///////////////////////////////////////////
}

func eval_register_name(reg byte, wordop byte) string {
	register_prefix := ""
	register_full := ""

	switch {
	case reg == reg_AL_AX:
		register_prefix = "A"
	case reg == reg_CL_CX:
		register_prefix = "C"
	case reg == reg_DL_DX:
		register_prefix = "D"
	case reg == reg_BL_BX:
		register_prefix = "B"
	}

	if wordop == 0b0 {
		register_full = register_prefix + "L"
	} else {
		register_full = register_prefix + "X"
	}

	return register_full
}

func eval_rm_name(rm byte, wordop byte) string {
	register_prefix := ""
	register_full := ""

	switch {
	case rm == rm_AL_AX:
		register_prefix = "A"
	case rm == rm_CL_CX:
		register_prefix = "C"
	case rm == rm_DL_DX:
		register_prefix = "D"
	case rm == rm_BL_BX:
		register_prefix = "B"
	}

	if wordop == 0b0 {
		register_full = register_prefix + "L"
	} else {
		register_full = register_prefix + "X"
	}

	return register_full
}

func check(error error) {
	if error != nil {
		panic(error)
	}
}
