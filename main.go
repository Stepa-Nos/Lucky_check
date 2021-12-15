package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main () {
	fmt.Println("Привет, хочешь проверить свою удачу?")
	fmt.Print("Я загадал число от 1 до 10, но у тебя будет только 3 попытки\n", "\n")
	fmt.Println("Думаешь сможешь угадать? ")
	fmt.Println("Введи своё число и проверь это :)")

	// Генерация числа за счёт времени
	seed_number:= time.Now().Unix()
	rand.Seed(seed_number)
	my_random_number:= rand.Intn(10)+1

	lucky := false // очень нужна для правильного срабатывания цикла при удачи

	// Цикла и кол-во попыток
	for compare := 0;compare<3; compare++{
		fmt.Println("Осталось попыток: ", 3-compare)

		// получение номера пользователя
		var user_number int
		fmt.Scan(&user_number)

		if my_random_number > user_number{
			fmt.Println("Ты не угадал, твоё число меньше")
		} else if  my_random_number < user_number{
			fmt.Println("Ты не угадал, твоё число больше")
		} else {
			lucky = true
			fmt.Println("Вау, твоя удача - это нечто! Поздравляю ;)")
			break
		}
	}
	if !lucky{ // Переводит из false в true, для срабатывания
		fmt.Println("Вау, твоя удача - это нечто! Ты проиграл :( ")
		fmt.Println("Моё число было", my_random_number)
	}
}