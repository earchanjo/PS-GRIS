package main

import(
	"fmt"
	"encoding/hex"
	"strconv"
	"strings"
)
var dec64 = []string{"A","B","C","D","E","F","G","H","I","J","K","L",
					"M","N","O","P","Q","R","S","T","U","V","W","X",
					"Y","Z","a","b","c","d","e","f","g","h","i","j",
					"k","l","m","n","o","p","q","r","s","t","u","v",
					"w","x","y","z","0","1","2","3","4","5","6","7",
					"8","9","+","/"}

var hex_frase string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

/*---- Esta funcao converte os caracteres (em dupla) hexadecimais
  ----da frase em seus codigos ascii correspondentes */
func ts( f string ) []byte {
	h, _ := hex.DecodeString(f)
	return h
}
/*---- Esta funcao pega cada elemento do array de string fornecido pela
 ----- funcao ts e os converte em binarios*/

 func ASCII_to_bin(input []byte) string {
	var ret []string
	var ret2 []string

	for _, char := range input {
		var char2 int64  =  int64(char)
		ret = append(ret,strconv.FormatInt(char2,2))
	}

	for _, ch := range ret {
		if len(ch) < 8 {
			ch2 := ch
			for i := 0; i < (8-len(ch)); i++ {
				ch2 = "0" + ch2
			}
			ret2 =  append(ret2,ch2)
		}
	}
	
	return strings.Join(ret2, "")
}

func bin_to_b64 (input string) string {
	//var ret string
	var array_of_bins []string

	if (len(strings.Split(input, "")) % 3) == 0 {
		input_splited := strings.Split(input, "")
		pos := 0
		p := 0
		//fmt.Println("len(input_splited)",len(input_splited))
		//fmt.Println(input_splited)

		for p < len(input_splited) {
			if pos == len(input_splited) {break}
			str := strings.Join(input_splited[pos:pos+6], "")
			array_of_bins = append(array_of_bins, str)
			pos += 6
			p++
		}
		
	}

	// Convertendo binario para decimal
	var bin_int []int64
        //var str_bin string
        for _, bin := range array_of_bins{
                int_bin, _ := strconv.ParseInt(bin,2,64)
                //fmt.Println(int_bin)
                //str_bin = string(int_bin)
				bin_int = append(bin_int, int_bin)
		}


		var list_char []string
        for _, num := range bin_int{
                list_char = append(list_char, dec64[num])
        }
		decoded_msg := strings.Join(list_char, "")
		
	return decoded_msg
}


func main() {
	fmt.Println("\tCRYPTOPALS SET 1 - HEX TO B64:")
    ret := ASCII_to_bin(ts(hex_frase))
	ret2 := bin_to_b64(ret)
	fmt.Println("Frase em HEX: ", hex_frase)
	fmt.Println("Frase codificada em B64: ", ret2)
}