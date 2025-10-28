package main

import "fmt"

// cтруктура Bashnya — стек строк
type Bashnya struct {
	bashnya []string
}

// добавить элемент
func (b *Bashnya) Push(x string) {
	b.bashnya = append(b.bashnya, x)
}

// снять верхний элемент
func (b *Bashnya) Pop() (string, bool) {
	var result string
	var ok bool

	if len(b.bashnya) > 0 {
		result = b.bashnya[len(b.bashnya)-1]
		b.bashnya = b.bashnya[:len(b.bashnya)-1]
		ok = true
	}

	return result, ok
}

// проверить пустая ли башня
func (b *Bashnya) IsEmpty() bool {
	var empty bool
	if len(b.bashnya) == 0 {
		empty = true
	}
	return empty
}

// нзнать высоту башни
func (b *Bashnya) Size() int {
	size := len(b.bashnya)
	return size
}

// очистить башню
func (b *Bashnya) Clear() {
	b.bashnya = nil
}

// показать содержимое башни
func (b *Bashnya) Show() {
	if b.IsEmpty() {
		fmt.Println("Башня пуста")
	} else {
		fmt.Println("Содержимое башни:")
		for i := len(b.bashnya) - 1; i >= 0; i-- {
			fmt.Printf("%s\n", b.bashnya[i])
		}
	}
}

func main() {
	var ostankino Bashnya

	ostankino.Push("первый кирпич")
	ostankino.Push("второй кирпич")
	ostankino.Push("третий кирпич")

	ostankino.Show()
	fmt.Println("Высота:", ostankino.Size())

	brick, ok := ostankino.Pop()
	if ok {
		fmt.Println("Сняли:", brick)
	} else {
		fmt.Println("Башня уже пуста.")
	}

	ostankino.Show()
	ostankino.Clear()
	ostankino.Show()
}
