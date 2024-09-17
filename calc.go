package main

import (
	"fmt"
	"log"

	"strconv"
)

func processResult(firstVal int, operator string, secVal int) (int, error) {
	var result int

	switch operator {
	case "+":
		result = add(firstVal, secVal)
		return result, nil
	case "-":
		result = subtract(firstVal, secVal)
		return result, nil
	case "*":
		result = multiply(firstVal, secVal)
		return result, nil
	case "/":
		result = divide(firstVal, secVal)
		return result, nil
	default:
		err := fmt.Errorf("ошибка: неверный оператор")
		return 0, err
	}

}
func main() {

	firstInput, secInput, operator, err := getInput()
	if err != nil {
		log.Fatal(err)
	}
	var result int
	var isFirstRomeNumber, isSecondRomeNumber bool
	romeToArabicNumbers := make(map[string]int)
	addValues := func() {
		romeToArabicNumbers["I"] = 1
		romeToArabicNumbers["II"] = 2
		romeToArabicNumbers["III"] = 3
		romeToArabicNumbers["IV"] = 4
		romeToArabicNumbers["V"] = 5
		romeToArabicNumbers["VI"] = 6
		romeToArabicNumbers["VII"] = 7
		romeToArabicNumbers["VIII"] = 8
		romeToArabicNumbers["IX"] = 9
		romeToArabicNumbers["X"] = 10
	}
	addValues()
	fmt.Println("firstInput is " + firstInput)
	fmt.Println("secInput is " + secInput)
	firstVal, ok := romeToArabicNumbers[firstInput]
	if ok {
		fmt.Println("First Val:", firstVal)
		isFirstRomeNumber = true
	} else {
		isFirstRomeNumber = false
		fmt.Println("First Val not found")
		firstVal, err = strconv.Atoi(firstInput)
		if err != nil {
			log.Fatal(err)
		}

	}

	secVal, ok := romeToArabicNumbers[secInput]
	if ok {
		fmt.Println("Sec Val:", secVal)
		isSecondRomeNumber = true
	} else {
		isSecondRomeNumber = false
		fmt.Println("Sec Val not found")
		secVal, err = strconv.Atoi(secInput)
		if err != nil {
			log.Fatal(err)
		}

	}
	if isFirstRomeNumber == true && isSecondRomeNumber == true {
		arabicToRomeNumbers := make(map[int]string) //init map
		func() {
			// Заполнение мапы
			arabicToRomeNumbers[1] = "I"
			arabicToRomeNumbers[2] = "II"
			arabicToRomeNumbers[3] = "III"
			arabicToRomeNumbers[4] = "IV"
			arabicToRomeNumbers[5] = "V"
			arabicToRomeNumbers[6] = "VI"
			arabicToRomeNumbers[7] = "VII"
			arabicToRomeNumbers[8] = "VIII"
			arabicToRomeNumbers[9] = "IX"
			arabicToRomeNumbers[10] = "X"
			arabicToRomeNumbers[11] = "XI"
			arabicToRomeNumbers[12] = "XII"
			arabicToRomeNumbers[13] = "XIII"
			arabicToRomeNumbers[14] = "XIV"
			arabicToRomeNumbers[15] = "XV"
			arabicToRomeNumbers[16] = "XVI"
			arabicToRomeNumbers[17] = "XVII"
			arabicToRomeNumbers[18] = "XVIII"
			arabicToRomeNumbers[19] = "XIX"
			arabicToRomeNumbers[20] = "XX"

			for i := 21; i <= 100; i++ {
				switch {
				case i < 30:
					arabicToRomeNumbers[i] = "XX" + arabicToRomeNumbers[i-20]
				case i < 40:
					arabicToRomeNumbers[i] = "XXX" + arabicToRomeNumbers[i-30]
				case i < 50:
					arabicToRomeNumbers[i] = "XL" + arabicToRomeNumbers[i-40]
				case i < 60:
					arabicToRomeNumbers[i] = "L" + arabicToRomeNumbers[i-50]
				case i < 70:
					arabicToRomeNumbers[i] = "LX" + arabicToRomeNumbers[i-60]
				case i < 80:
					arabicToRomeNumbers[i] = "LXX" + arabicToRomeNumbers[i-70]
				case i < 90:
					arabicToRomeNumbers[i] = "LXXX" + arabicToRomeNumbers[i-80]
				case i < 100:
					arabicToRomeNumbers[i] = "XC" + arabicToRomeNumbers[i-90]
				default:
					arabicToRomeNumbers[i] = "C"
				}
			}
		}()
		result, err = processResult(firstVal, operator, secVal)
		if err != nil {
			log.Fatal(err)
		}
		if result < 1 {
			log.Fatal("Римские цифры  не <1")
		}
		resultInRomeNumber, _ := arabicToRomeNumbers[result]
		fmt.Printf("Результат: %s", resultInRomeNumber)
	} else if isFirstRomeNumber == false && isSecondRomeNumber == false {
		result, err = processResult(firstVal, operator, secVal)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Результат: %d", result)

	} else {
		log.Fatal("Ошибка ввода")
	}
}

func add(a, b int) int {
	if a < 1 || a > 10 || b < 1 || b > 10 {
		log.Fatal("Входные значения >=1 && <=10")
		return 0
	}
	return a + b
}
func subtract(a, b int) int {
	if a < 1 || a > 10 || b < 1 || b > 10 {
		log.Fatal("Входные значения >=1 && <=10")
		return 0
	}
	return a - b
}
func multiply(a, b int) int {
	if a < 1 || a > 10 || b < 1 || b > 10 {
		log.Fatal("Входные значения >=1 && <=10")
		return 0
	}
	return a * b
}
func divide(a, b int) int {
	if b == 0 {
		log.Fatal("деление на ноль")
		return 0
	} else if a < 1 || a > 10 || b < 1 || b > 10 {
		log.Fatal("Входные значения >=1 && <=10")
		return 0
	}
	return a / b

}

func getInput() (string, string, string, error) {

	var a, b string
	var operator string

	fmt.Println("Введите три значения через пробел:")
	_, err := fmt.Scanln(&a, &operator, &b)
	if err != nil {
		return "", "", "", err
	}

	return a, b, operator, nil
}
