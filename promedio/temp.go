package promedio

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"strings"
)


func Temperaturas(){
	scaner:= bufio.NewReader(os.Stdin)

	var mayor int
	var menor int
	var promedio int
	total := 0
	primero := true
	invalido := 0
	validos := 0

	frio := 0
	normal := 0
	calor := 0

	fmt.Println("ingrese la tenperatura(escribir [salir] para salir)")

	for {

	temperatura,_ := scaner.ReadString('\n')
	temperatura = strings.TrimSpace(temperatura)

	if temperatura == "salir"{
		break
	}

	estadistica,err:= strconv.Atoi(temperatura)

	if err != nil{
		invalido++
		continue
	}

	validos++
	
	if estadistica < -50 && estadistica > 50 {
		invalido++
		continue
	}

//temperatura mas alta y mas baja papa

if primero {
	mayor = estadistica
	menor = estadistica
	primero = false
}

	if estadistica < menor {
	menor = estadistica
}

	if estadistica > mayor {
		mayor = estadistica
	}

	total += estadistica

	
	
	if validos > 0 {
		promedio = total/ validos	
	}


// clasificaciones

if estadistica < 0 {
  frio++
}

if estadistica > 0 && estadistica < 25 {
	normal++
}

if estadistica > 25 {
	calor++
}

	}

	fmt.Println("promedio:",promedio)
	fmt.Println("temperatura mas alta:",mayor)
	fmt.Println("temperatura mas baja:",menor)
	fmt.Println("cantidad de datos validos:",validos)
	fmt.Println("cantidad de datos invalidos:",invalido)

	fmt.Println("||||CLASIFICACION||||")
	fmt.Println("frio:",frio)
	fmt.Println("normal:",normal)
	fmt.Println("calor:",calor)
}