package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gilbraz28/aula01GO/calculos"
	"github.com/gilbraz28/aula01GO/message"
	"github.com/gilbraz28/aula01GO/message/say"
)

func main() {

	var soma int

	listaEntradas := os.Args[1:len(os.Args)]

	valor1, err1 := strconv.Atoi(listaEntradas[0])
	valor2, err2 := strconv.Atoi(listaEntradas[1])

	if err1 != nil || err2 != nil {
		fmt.Println("")
	}

	msg := "VAMOS FAZER A SOMA e MULTIPLICAÇÃO DOS NÚMEROS RECEBIDOS: " + strconv.Itoa(valor1) + " e " + strconv.Itoa(valor2)
	message.Show(msg)

	soma = execSoma(int(valor1), int(valor2))
	multiplicao := execMultiplicacao(int(valor1), int(valor2))

	msg = "O resultado da soma é: " + strconv.Itoa(soma) + ". "

	say.Show(msg)

	msg = "O resultado da multiplicação é: " + strconv.Itoa(multiplicao) + ". "

	say.Show(msg)

}

func execSoma(vl1, vl2 int) int {

	result := calculos.Soma(int(vl1), int(vl2))

	//soma++

	return result

}

func execMultiplicacao(vl1, vl2 int) int {

	result := calculos.Multiplica(int(vl1), int(vl2))

	//soma++

	return result

}
