package whatis

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func WhatIsNumber() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100) + 1
	defer fmt.Println("Загаданное число было:", number)
	fmt.Println("Число загадано от 1 до 100")
	fmt.Printf("У вас %v попыток.\n", 5)
	r := bufio.NewReader(os.Stdin)
	for x := 0; x < 5; x++ {

		fmt.Println("Какое это число? Попытка:", x+1)
		fmt.Print("Введите число: ")
		i, _ := r.ReadString('\n')
		i = strings.TrimSpace(i)
		k, _ := strconv.Atoi(i)

		switch {
		case k < number:
			log.Println("Загаданное число больше")
			continue
		case k > number:
			log.Println("Загаданное число меньше")
			continue
		default:
			log.Println("Вы угадали!!")
			return
		}

	}

	fmt.Println("Вы проиграли :(")

}
