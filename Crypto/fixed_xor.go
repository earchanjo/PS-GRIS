package main

import (
	"fmt"
	"encoding/hex"
)

// ---- Esta funcao converte a string em hexadecimal para array de bytes ascii correspondentes
func string_to_byte(input string) []byte {
	input_to_bytes := []byte(input)
	message := make([]byte, hex.DecodedLen(len(input))) 	// Cria um array de bytes prontos para receber os bytes em ascii 
	hex.Decode(message, input_to_bytes)

	return message
}

func x_o_r(input1, input2 []byte) []byte {

	if len(input1) != len(input2) {
		fmt.Println("Entradas com tamanhos diferentes")
		
	}
	var msg []byte

	for i,_ := range input1 {
			msg = append(msg, input1[i]^input2[i])		// "^" É o operador binario xor em go
		}
	return msg

}

func main() {
	fmt.Println("CRYPTOPALS - SET1 - EXERCICIO2 - FIXED XOR")
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"

	final_msg := x_o_r(string_to_byte(input1), string_to_byte(input2))
	fmt.Println("FIXED XOR MESSAGE IS: ", final_msg)
	fmt.Printf("%x",final_msg) // %x formatacao para saída em hexadecimal
}