package main

import "fmt"

func main() {
	var hello2 = "ShalomVsem"       //Один из вариантов обозначения переменной
	var hello3 = "Osobenniy Shalom" //Значение "string" определилось само.
	hello := "Shalom"               //Двоеточие перед равно выводит тип данных из присваеваемого значения
	var (
		name string = "somebody"
		age  int
	) //Вариант в виде списка
	fmt.Println(hello2, hello, name, age, hello3)
}

// Если возраст больше 35 - особенный шалом, спростиь возраст.
