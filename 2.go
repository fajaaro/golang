package main

import "fmt"

func main() {
	var str1 string
	var str2 string
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	str1usage := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}
	str2usage := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}

	fmt.Print("Input String 1 : ")
	_, err := fmt.Scanln(&str1)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	fmt.Print("Input String 2 : ")
	_, err = fmt.Scanln(&str2)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	//selisih := len(str1) - len(str2)
	//if selisih < -1 || selisih > 1 {
	//	fmt.Println("Result : False")
	//} else {
	//	//for _
	//
	//	fmt.Println("Result : True")
	//}

	for pos, char := range str1 {
		fmt.Println(pos, string(char))
		str1usage[string(char)] += 1
	}
	for pos, char := range str2 {
		fmt.Println(pos, string(char))
		str2usage[string(char)] += 1
	}
	fmt.Println(str1usage)
	fmt.Println(str2usage)

	totalDifferent := 0
	for _, letter := range letters {
		if str1usage[letter] != str2usage[letter] {
			totalDifferent++
		}
	}

	fmt.Println("Result:", totalDifferent <= 2)
}
