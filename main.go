package main

import (
	"fmt"
	"log"

	"github.com/andreyd97/average/datafile"
)

func main() {
	var output float64
	test, err := datafile.GetFloat("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, sum := range test { // цикл для прохода по всем элементам сегмента
		output += sum //суммируем все элементы сегмента
	}

	output = output / float64(len(test)) //вычисляем среднее (сумма чисел массива делим на кол-во элементов в сегименте)
	fmt.Println(output)                  //выводим среднее

}
