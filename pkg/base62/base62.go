package base62

import "math"

var base62Str = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func MustInit(s string) {
	if len(s) != 62 {
		panic("base62Str长度不对")
	}
	base62Str = s
}

func Base62Encode(num int64) string {
	if num == 0 {
		return string(base62Str[0])
	}

	result := ""
	for num > 0 {
		remainder := num % 62
		result = string(base62Str[remainder]) + result
		num = num / 62
	}

	return result
}

func Base62Decode(str string) int64 {
	result := int64(0)
	for i, char := range str {
		index := -1
		for j, c := range base62Str {
			if c == char {
				index = j
				break
			}
		}
		if index == -1 {
			panic("Invalid character in input string")
		}
		result += int64(index) * int64(math.Pow(62, float64(len(str)-i-1)))
	}

	return result
}
