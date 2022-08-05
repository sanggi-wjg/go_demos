package clock

import "fmt"

type placeholder [5]string

var zero = placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var seperator = placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

var digits = [10][5]string{zero, one, two, three, four, five, six, seven, eight, nine}

func PrintDigit() {
	for digit := range digits {
		for line := range digits[0] {
			fmt.Println(digits[digit][line])
		}

		fmt.Println()
	}
}

func PrintTime(hour, minute, second int) {
	timeDigits := [...][5]string{
		digits[hour/10], digits[hour%10], seperator,
		digits[minute/10], digits[minute%10], seperator,
		digits[second/10], digits[second%10],
	}

	for line := range timeDigits[0] {
		for digit := range timeDigits {
			fmt.Print(timeDigits[digit][line], "  ")
		}
		fmt.Println()
	}
}
