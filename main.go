package main

import (
	"fmt"
	"os"
	"strconv"
)

func timeconversion(str string) string {

	checkAMPM := str[len(str)-2:]
	checkHour := str[0:2]

	intHour, _ := strconv.Atoi(checkHour)

	newHour := intHour + 12

	strNewHour := strconv.Itoa(int(newHour))

	restTime := str[2:8]

	if checkAMPM == "AM" && newHour != 24 {
		return str[0:8]
	}

	if checkAMPM == "AM" && newHour == 24 {
		return "00" + str[2:8]
	}

	if checkAMPM == "PM" && newHour == 24 {
		return checkHour + restTime
	} else {
		return strNewHour + restTime
	}

}

func main() {

	stringTime1 := "00:45:30AM"  // 00:45:30
	stringTime2 := "00:45:30PM"  // 12:45:30
	stringTime3 := "12:45:30AM"  // 00:45:30
	stringTime4 := "12:45:30PM"  // 12:45:30
	stringTime5 := "11:45:30AM"  // 11:45:30
	stringTime6 := "11:45:30PM"  // 23:45:30
	stringTime7 := "00:00:00AM"  // 00:00:00
	stringTime8 := "00:00:00PM"  // 12:00:00
	stringTime9 := "12:00:00AM"  // 00:00:00
	stringTime10 := "12:00:00PM" // 12:00:00

	fmt.Println(timeconversion(stringTime1))
	fmt.Println(timeconversion(stringTime2))
	fmt.Println(timeconversion(stringTime3))
	fmt.Println(timeconversion(stringTime4))
	fmt.Println(timeconversion(stringTime5))
	fmt.Println(timeconversion(stringTime6))
	fmt.Println(timeconversion(stringTime7))
	fmt.Println(timeconversion(stringTime8))
	fmt.Println(timeconversion(stringTime9))
	fmt.Println(timeconversion(stringTime10))

	/// .exe
	fmt.Println("")
	fmt.Println("input 12-hour AM/PM format time to convert below: (example: 11:45:30PM)")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("")
	fmt.Println("==================================")
	fmt.Println("")
	fmt.Println("24-hour format time conversion result are below:")
	fmt.Println(timeconversion(input))

	fmt.Scanln()

}
