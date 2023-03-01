package main


import "fmt"

// Code by Rhett Applestone


func main() {

// Usr input of hex to be converted to binary blocks
	fmt.Print("Hex to be converted: ")
	hex := ""
	fmt.Scanln(&hex)

// Defining Slice that each char of input will be put in
	slice := make([]string, 64)


// Defining output variable which data will be put into
	var output string = ""

// For loop that puts individual characters into slice
	for i := 0; i < len(hex); i++ {
		slice[i] = string(hex[i])

	}


// for loop that puts each binary representation of the hex into the output varible with switch

	for j := 0; j < 64; j++ {
		switch slice[j] {
		case "0": output += "0000"
		case "1": output += "0001"
		case "2": output += "0010"
		case "3": output += "0011"
		case "4": output += "0100"
		case "5": output += "0101"
		case "6": output += "0110"
		case "7": output += "0111"
		case "8": output += "1000"
		case "9": output += "1001"
		case "a": output += "1010"
		case "b": output += "1011"
		case "c": output += "1100"
		case "d": output += "1101"
		case "e": output += "1110"
		case "f": output += "1111"
		}

	}


	fmt.Println(output)
}
