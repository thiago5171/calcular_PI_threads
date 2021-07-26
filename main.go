package main

import (
	"fmt"
	"math"
	"time"
)


const numExec = 5
var valoresPi, tempos, pi2 [numExec]float64
var parcela float64
var nThreads, termos int

func calcularPi(ninicio, termos int, c chan float64) {
	var pi float64

	for n := ninicio; n < ninicio+termos; n++ {
		pi += math.Pow(-1, float64(n)) / float64(2 * n + 1)
	}

	pi *= 4
	c <- pi
}

func main() {
	fmt.Printf("Numero de termos ")
	_, _ = fmt.Scanln(&termos)

	fmt.Printf("Numero de threads: ")
	_, _ = fmt.Scanln(&nThreads)

	var ninicio int

	for i := 0; i < numExec; i++ {
		inicio := time.Now()
		parcela := termos/nThreads
		ninicio = 0
		var canal []chan float64
		for j := 0; j < nThreads; j++ {
			canal = append(canal, make(chan float64))
			go calcularPi(ninicio, parcela ,canal[j])

			ninicio += parcela
		}

		for _, c := range canal {
			valoresPi[i] += <-c
		}

		tempos[i] = time.Since(inicio).Seconds()
	}

	var soma float64
	for _, tempoGasto := range tempos {
		soma += tempoGasto
	}
	media := soma / float64(numExec)

	var execTime float64
	for _, tempoGastoPorThread := range tempos {
		t := tempoGastoPorThread - media
		execTime += math.Pow(t, 2) / float64(numExec)
	}
	desvioPadrao := math.Sqrt(execTime)

	
	cv := (desvioPadrao / media) * 100

	fmt.Println("Valores obtidos para Pi -> ", valoresPi)
	fmt.Println("Tempo Médio ->", media)
	fmt.Println("Desvio padrão -> ", desvioPadrao)
	fmt.Println("Coeficiente de variação -> ", cv)
}

