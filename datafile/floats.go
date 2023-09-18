package datafile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetFloat(filename string) ([]float64, error) {
	var digits []float64           // переменна для хранения чисел полученных из файла
	open, err := os.Open(filename) // открываем файл
	if err != nil {                // обрабатываем ошибку которая получилась при открытии файла
		return digits, err
	}

	reader := bufio.NewScanner(open) // создаем переменную для сканирования данных из файла

	for reader.Scan() { // создаем цикл для прохода по всем строкам файла
		test, err := strconv.ParseFloat(reader.Text(), 64) // получаем значение и преобразуем его String -> Float с битностью 64bit
		if err != nil {                                    // обрабатываем ошибку которая получилась при преобразовании строки в вещественное число
			return digits, err
		}
		digits = append(digits, test) //сохраняем полученные значения в сегмент

	}

	fmt.Println(digits) //выводим содержимое сегмента

	err = open.Close() // закрываем файл
	if err != nil {    // обрабатываем ошибку которая получилась при закрытии файла
		return digits, err
	}

	if reader.Err() != nil { // если сканирование произошло с ошибкойR
		return digits, reader.Err() // выводим эту ошибку
	}
	return digits, err
}
