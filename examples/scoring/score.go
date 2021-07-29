package main

import (
	"bytes"
	"fmt"
)

var VOWELS = []byte("AEIOU")
var NUMBERS = []byte("0123456789")
var NOT_VOWELS = []byte("BCDFGHJKLMNPQRSTXYZ")

type TEST struct {
	nono int
}

type Code []byte

func (c Code) Is(index int, chars []byte) bool {
	return bytes.ContainsAny(chars, string(c[index]))
}

func main() {
	str := []byte("ABCDEFGHIJKLMNOPQRSTUXYZ123456789")
	// str := []byte("AE12R")
	allLexicographic(str, 6)

	// codes := []Code{
	// 	// []byte("I9ROPZ"),
	// 	// []byte("PIP898"),
	// 	// []byte("REP70d"),
	// 	// []byte("TIPKEK"),
	// 	// []byte("REPRER"),
	// 	// []byte("REPREP"),
	// 	// []byte("012345"),
	// 	// []byte("123POP"),
	// 	// []byte("123CWQ"),
	// 	// []byte("DIDIDI"),
	// 	// []byte("123ABC"),
	// 	// []byte("888888"),
	// 	// []byte("TIKTOK"),
	// 	// []byte("PTBTBT"),
	// 	// []byte("123457"),
	// 	// []byte("FAABAA"),
	// 	// []byte("FAABAB"),
	// 	[]byte("FAB039"),
	// }

	// for _, v := range codes {
	// 	fmt.Printf("%4d - %v\n", score6(v), string(v))
	// }

	// fmt.Println(codes)

}

func score6(code Code) (score int) {

	// vowel for second cahracter
	if code.Is(1, VOWELS) {
		score += 50

		// vowels be next to each other has negative impact
		if code.Is(0, VOWELS) {
			score -= 250
		}
		if code.Is(2, VOWELS) {
			score -= 250
		}

		// same char around vowel
		if code[0] == code[3] {
			score += 100
		}

		// number around vowel has bad impact
		if code.Is(0, NUMBERS) || code.Is(2, NUMBERS) {
			score -= 350
		}

	}

	// vowel for fourth cahracter
	if code.Is(4, VOWELS) {
		score += 50
		if code.Is(3, VOWELS) {
			score -= 25
		}
		if code.Is(5, VOWELS) {
			score -= 25
		}

		// same char around vowel
		if code[3] == code[5] {
			score += 100
		}

		// number around vowel has bad impact
		if code.Is(3, NUMBERS) || code.Is(5, NUMBERS) {
			score -= 350
		}
	}

	if code.Is(1, VOWELS) && code.Is(4, VOWELS) {
		if code[1] == code[4] {
			score += 200
		}

		if code[0] == code[3] {
			score += 110
		}

		if code[2] == code[5] {
			score += 110
		}
	}

	// compare two part
	if bytes.Compare(code[0:3], code[3:6]) == 0 {
		score += 450
	}

	// compare [1,2] and [4,5]
	if code[1] == code[4] && code[2] == code[5] {
		score += 110
	}

	// pair scoring [0,1] and [1,2] and [2,3] and [3,4] and [4,5]
	if code[0] == code[1] {
		score += 15
	}
	if code[1] == code[2] {
		score += 15
	}
	if code[2] == code[3] {
		score += 15
	}
	if code[3] == code[4] {
		score += 15
	}
	if code[4] == code[5] {
		score += 15
	}

	// compare three part
	if bytes.Compare(code[0:2], code[2:4]) == 0 && bytes.Compare(code[2:4], code[4:6]) == 0 {
		score += 300
	}

	// compare two part at the begining
	if bytes.Compare(code[0:2], code[2:4]) == 0 {
		score += 90
	}

	// compare two part at the end
	if bytes.Compare(code[2:4], code[4:6]) == 0 {
		score += 90
	}

	if code[0]+1 == code[1] && code[1]+1 == code[2] {
		score += 40
	}

	if code[3]+1 == code[4] && code[4]+1 == code[5] {
		score += 40
	}

	count := countChar(code)

	switch count {
	case 5:
		score += 5
	case 4:
		score += 10
	case 3:
		score += 20
	case 2:
		score += 40
	case 1:
		score += 100
	}

	if code.Is(0, NOT_VOWELS) &&
		code.Is(1, VOWELS) &&
		code.Is(2, NOT_VOWELS) &&
		code.Is(3, NUMBERS) &&
		code.Is(4, NUMBERS) &&
		code.Is(5, NUMBERS) {

		score += 600
	}

	return
}

func countChar(code []byte) (count int) {
	charMap := make(map[byte]int)
	for _, v := range code {
		if _, ok := charMap[v]; ok {
			charMap[v]++
		} else {
			charMap[v] = 1
			count++
		}
	}

	return
}

// permutation
func allLexicographicRecur(str []byte, data []byte, last, index int) {
	l := len(str)

	for i := 0; i < l; i++ {
		data[index] = str[i]

		if index == last {
			// fmt.Println(string(data))
			score := score6(data)
			if score > 310 {
				fmt.Printf("%04d,%v\n", score, string(data))
			}
		} else {
			allLexicographicRecur(str, data, last, index+1)
		}
	}
}

func allLexicographic(str []byte, l int) {

	data := make([]byte, l)
	// var temp = str

	// fmt.Println(data, temp)

	allLexicographicRecur(str, data, l-1, 0)

}
