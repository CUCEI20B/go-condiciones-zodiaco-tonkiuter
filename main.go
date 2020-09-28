package main

import "fmt"

func main(){
	var dia int
	var mes int

	fmt.Scan(&dia)
	
	fmt.Scan(&mes)

	switch{
		case dia >= 21 && dia <= 31 && mes == 1 || dia >= 1 && dia <= 18 && mes == 2:
			fmt.Println("acuario")
		case dia >= 19 && dia <= 29 && mes == 2 || dia >= 1 && dia <= 20 && mes == 3:
			fmt.Println("piscis")
		case dia >= 21 && dia <= 31 && mes == 3 || dia >= 1 && dia <= 20 && mes == 4:
			fmt.Println("aries")
		case dia >= 21 && dia <= 30 && mes == 4 || dia >= 1 && dia <= 20 && mes == 5:
			fmt.Println("tauro")
		case dia >= 21 && dia <= 30 && mes == 5 || dia >= 1 && dia <= 21 && mes == 6:
			fmt.Println("geminis")
		case dia >= 22 && dia <= 30 && mes == 6 || dia >= 1 && dia <= 22 && mes == 7:
			fmt.Println("cancer")
		case dia >= 23 && dia <= 31 && mes == 7 || dia >= 1 && dia <= 22 && mes == 8:
			fmt.Println("leo")
		case dia >= 23 && dia <= 31 && mes == 8 || dia >= 1 && dia <= 22 && mes == 9:
			fmt.Println("virgo")
		case dia >= 23 && dia <= 30 && mes == 9 || dia >= 1 && dia <= 22 && mes == 10:
			fmt.Println("libra")
		case dia >= 23 && dia <= 31 && mes == 10 || dia >= 1 && dia <= 22 && mes == 11:
			fmt.Println("escorpion")
		case dia >= 23 && dia <= 30 && mes == 11 || dia >= 1 && dia <= 21 && mes == 12:
			fmt.Println("sagitario")
		case dia >= 22 && dia <= 31 && mes == 12 || dia >= 1 && dia <= 20 && mes == 1:
			fmt.Println("capricornio")
		default:
			fmt.Println("Opcion no Valida!")
	}

}