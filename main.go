package main

import "fmt"

/*
Описание:
Тебе нужно написать код, который будет работать с двумя слайсами a и b, который содержит целые числа.
Задача:
Вернуть два слайса с уникальными элементами.
Найти пересечения значений в двух слайсах.
Вернуть слайс с уникальными элементами из слайсов a и b.
Примечания:
Порядок вывода чисел в слайсах не имеет значения.
Разбивать на файлы не нужно, просто пиши в одном файле main.go.
Решение задачи через AI недопустимо.
Допускается неполное решение, то есть могут быть выполнены не все пункты.


Пример работы программы:
Входные данные: a = [5, 1, 2, 5], b = [4, 2, 5, 1, 1, 2]
Выходные данные:
[5, 1, 2], [4, 2, 5, 1]
[5, 1, 2]
[5, 1, 2, 4]

*/

func main() {
	var a, b = []int{5, 1, 2}, []int{4, 2, 5, 1, 1, 2}

	fmt.Println(uniqueElements(a, b))
	fmt.Println(intersect(a, b))
	fmt.Println(uniqueAll(a, b))
}

func uniqueElements(a, b []int) ([]int, []int) {
	uniqueMapA, uniqueMapB := make(map[int]int), make(map[int]int)
	var uniqueA, uniqueB []int

	for _, elA := range a {
		uniqueMapA[elA] = 0
	}

	for _, elB := range b {
		uniqueMapB[elB] = 0
	}

	for i := range uniqueMapA {
		uniqueA = append(uniqueA, i)
	}

	for j := range uniqueMapB {
		uniqueB = append(uniqueB, j)
	}

	return uniqueA, uniqueB
}

func uniqueAll(a, b []int) []int {
	uniqueMap := make(map[int]int)
	var result []int

	for _, elA := range a {
		uniqueMap[elA] = 0
	}

	for _, elB := range b {
		uniqueMap[elB] = 0
	}

	for i := range uniqueMap {
		result = append(result, i)
	}

	return result
}

func intersect(a, b []int) []int {
	countA := make(map[int]int)

	for _, v := range a {
		countA[v]++
	}

	var result []int
	for _, v := range b {
		if countA[v] > 0 {
			result = append(result, v)
			countA[v]--
		}
	}

	return result
}
