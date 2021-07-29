/****************************
idea from
https://www.geeksforgeeks.org/print-all-permutations-with-repetition-of-characters/
*****************************/
package pkg

import "fmt"

func allLexicographicRecur(str []byte, data []byte, last, index int) {
	l := len(str)

	for i := 0; i < l; i++ {
		data[index] = str[i]
		if index == last {
			fmt.Println(string(data))
		} else {
			allLexicographicRecur(str, data, last, index+1)
		}
	}
}

func allLexicographic(str []byte, l int) {

	data := make([]byte, l+1)
	var temp = str

	fmt.Println(data, temp)

	allLexicographicRecur(str, data, l-1, 0)

}

func main() {
	// str := []byte("ABCDEFGHJKLMNPQRSTUWXYZ123456789")
	str := []byte("ABCDE")
	allLexicographic(str, 6)
}
