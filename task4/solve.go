package main

//import "fmt"
import "unicode"

func RemoveEven(a []int) []int {
	ret := make([]int, 0)
	for i := range a {
		if a[i] % 2 == 1 {
			ret = append(ret, a[i])
		}
	}
	return ret
}

func PowerGenerator(a int) func() int {
	b := a
	return func()(c int){
		c = b
		b = b * a
		return
	}
}


func DifferentWordsCount(s string) (ret int) {
	words := make(map[string]int)
	word := ""
	for i:= range s{
		if unicode.IsLetter(rune(s[i])){
			word = word + string(unicode.ToLower(rune(s[i])))
		} else {
			if word != "" {
				words[word]++;
			}
			word = ""
		}
	}
	if word != "" {
		words[word]++;
	}
	ret = len(words)
	return
}

/*
func main() {
	fmt.Println(DifferentWordsCount("hi hhhhi"))
	// Должно напечатать 2
}
*/

