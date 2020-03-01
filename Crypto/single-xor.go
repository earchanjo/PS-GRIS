package main

import (
	"fmt"
	"encoding/hex"
	"sort"
)

func calculate_score_of_xor (inpt string) int {
	score := 0
	
	for _, char := range inpt {

        switch {
        case byte(char) < 0x65:
            score = score - 1
        }

        switch string(char) {
        case "e":
			score = score + 5
        case "t":
            score = score + 5
        case "a":
            score = score + 4
        case "o":
            score = score + 4
        case "i":
            score = score + 4
        }
    }


	return score
}

func hex_to_string(input string) string {
	input_hex, _ := hex.DecodeString(input)

	dec_string := make([]byte, 94)
	score_for_char := make(map[int]string)
	score_array := make([]int, 94)

	for n,_ := range dec_string { 			// Existem ao todo 94 characteres para se fazer xor
		array_of_xor := make([]byte, 0)
		
		for _, char := range input_hex { // vamos aplicar para cada character o xor e criar um array dos resultados
			array_of_xor = append(array_of_xor, byte(n + 32) ^ char)  // n+32 representa o inicio dos characteres na tabela ascii
		}

		string_possibilities := string(array_of_xor)
		score_array[n] = calculate_score_of_xor(string_possibilities)	
		score_for_char[n] = string_possibilities

	}

	sortedScoreArray := make([]int, 94)
    copy(sortedScoreArray, score_array)

    sort.Ints(sortedScoreArray)
    sort.Sort(sort.Reverse(sort.IntSlice(sortedScoreArray)))

    highestScore := sortedScoreArray[0]
    var decoderIndex int

    for i, v := range score_array {
        if v == highestScore {
            decoderIndex = i
        }
    }

    //fmt.Println(score_for_char[decoderIndex])

	
	//fmt.Println(score_for_char)

	return score_for_char[decoderIndex]

}


func main() {
	fmt.Println("CRYPTOPALS - SET1 - EXERCICIO3 - SINGLE-BYTE XOR")
	frase_hex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	fmt.Println(hex_to_string(frase_hex))
}