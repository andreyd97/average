package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var digits []float64             // переменна для хранения чисел полученных из файла
	open, err := os.Open("data.txt") // открываем файл
	if err != nil {                  // обрабатываем ошибку которая получилась при открытии файла
		log.Fatal(err)
	}

	reader := bufio.NewScanner(open) // создаем переменную для сканирования данных из файла

	for reader.Scan() { // создаем цикл для прохода по всем строкам файла
		test, err := strconv.ParseFloat(reader.Text(), 64) // получаем значение и преобразуем его String -> Float с битностью 64bit
		if err != nil {                                    // обрабатываем ошибку которая получилась при преобразовании строки в вещественное число
			log.Fatal(err)
		}
		digits = append(digits, test) //сохраняем полученные значения в массив

	}

	fmt.Println(digits) //выводим содержимое массива

	err = open.Close() // закрываем файл
	if err != nil {    // обрабатываем ошибку которая получилась при закрытии файла
		log.Fatal(err)
	}

	if reader.Err() != nil { // если сканирование произошло с ошибкойR
		fmt.Println(reader.Err()) // выводим эту ошибку
	}

}
