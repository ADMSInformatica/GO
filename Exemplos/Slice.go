package main
import "fmt"

func main() {
	numeros := [10]int{1,2,3,4,5,6,7,8,9,10}
	var x[]int = numeros[0:10]
	var y[]int = numeros[0:5]
	fmt.Println(x)
	fmt.Println(y)
}

/*
func main() {
	nomes := [10]string{"ADMS", "Informática"}
	n := nomes[0:1]
	fmt.Println(n)
}
*/
