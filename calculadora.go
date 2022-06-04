package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

//Esta funcion se vuelve un metodo de calc
func (calc) operate(entrada string, operador string) int {

	valores := strings.Split(entrada, operador)

	operador1 := parsear(valores[0])

	operador2 := parsear(valores[1])

	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println(operador, " No esta soportado")
		return 0
	}

}

func parsear(entrada string) int {
	operador, err := strconv.Atoi(entrada)
	if err != nil {
		fmt.Println(err)
	}
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	return operacion

}
