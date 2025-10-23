package main

import "fmt"

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	if num >= 12307 {
		fmt.Printf("Число %d слишком большое — цикл не запущен.\n", num)
		return
	}

	for num < 12307 {
		if num < 0 {
			num *= -1
		} else if num%7 == 0 {
			num *= 39
		} else if num%9 == 0 {
			num = num*13 + 1
			continue
		} else {
			num = (num + 2) * 3
		}

		if num%13 == 0 && num%9 == 0 {
			fmt.Println("service error")
			break
		} else {
			num += 1
		}
	}

	fmt.Printf("Результат работы программы: %d\n", num)
}
