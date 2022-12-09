// declarando o pacote principal
package main

// importando a função fmt
import "fmt"

// declarando a variável CONST do ponto de ebulição da água em kelvin
const ebulicaoK float64 = 373.0

// função principal
func main() {

	var tempK float64 = ebulicaoK   // variável do valor da temperatura em kelvin
	var tempC float64 = tempK - 273 // conversão de kelvin para graus celsius

	//mostra na tela o resultado
	fmt.Println("A temperatura de ebulição da água em kelvin é", tempK)
	fmt.Println("A temperatura de ebulição da água em graus celsius é", tempC)

}
