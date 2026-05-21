package raindrops

import "strconv"

func Convert(number int) string {
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		return strconv.Itoa(number)
	}

	raindrops := ""
	if number%3 == 0 {
		raindrops += "Pling"
	}

	if number%5 == 0 {
		raindrops += "Plang"
	}

	if number%7 == 0 {
		raindrops += "Plong"
	}

	return raindrops
}
