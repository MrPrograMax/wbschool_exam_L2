package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println(UnpackString("a4bc2d5e"))
	fmt.Println(UnpackString("abcd"))
	fmt.Println(UnpackString("45"))
	fmt.Println(UnpackString(""))

}

func UnpackString(str string) string {
	var resultSb strings.Builder
	var curLetter rune

	for _, v := range []rune(str) {
		if unicode.IsLetter(v) {
			resultSb.WriteRune(v)
			curLetter = v
		} else if unicode.IsDigit(v) {
			count, err := strconv.Atoi(string(v))
			if err != nil {
				return ""
			}

			if curLetter == 0 {
				return ""
			}

			for i := 0; i < count-1; i++ {
				resultSb.WriteRune(curLetter)
			}
		}
	}

	return resultSb.String()
}
