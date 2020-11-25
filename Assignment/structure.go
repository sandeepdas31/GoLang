package main
import "fmt"

type person struct{
	 firstname string
	 lastname string
	 age int
	 ph_no int
}

func main() {
	var p person
	fmt.Println("enter name,age,ph_no")
	fmt.Scanln(&p.name)
	fmt.Scanln(&p.age)
	fmt.Scanln(&p.ph_no)
	fmt.Println(p)
}
