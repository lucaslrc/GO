package main

import "fmt"

func main() {

	var w int

	for w != 1 {
		var a int
		fmt.Println("\n1 - Soma")
		fmt.Println("2 - Subtração")
		fmt.Println("3 - Multiplicação")
		fmt.Println("4 - Divisão")
		fmt.Printf("5 - Sair\n")

		fmt.Scanf("%d", &a)

		switch a {
		case 1:
			var b, c int

			fmt.Println("\nDigite um número: ")
			fmt.Scanf("%d", &b)
			fmt.Println("Digite outro número: ")
			fmt.Scanf("%d", &c)
			fmt.Printf("Resultado: %d\n", b+c)

			break

		case 2:
			var b, c int

			fmt.Println("\nDigite um número: ")
			fmt.Scanf("%d", &b)
			fmt.Println("Digite outro número: ")
			fmt.Scanf("%d", &c)
			fmt.Printf("Resultado: %d\n", b-c)

			break

		case 3:
			var b, c int

			fmt.Println("\nDigite um número: ")
			fmt.Scanf("%d", &b)
			fmt.Println("Digite outro número: ")
			fmt.Scanf("%d", &c)
			fmt.Printf("Resultado: %d\n", b*c)

			break

		case 4:
			var b, c int

			fmt.Println("\nDigite um número: ")
			fmt.Scanf("%d", &b)
			fmt.Println("Digite outro número: ")
			fmt.Scanf("%d", &c)
			fmt.Printf("Resultado: %d\n", b/c)

			break

		case 5:
			fmt.Println("Até mais!")
			w++
			break
		}
	}
}
