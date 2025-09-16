package main

import (
	"fmt"

	"github.com/gbnunes7/meet"
)

const nome string = "Gustavo" // Global Variable Imutável



type Cliente struct {
	nome string
	idade int
	endereco string
}

type Pessoa struct {
	nome string
	idade int
}

func (p *Pessoa) Apresentar() {
	p.nome = "Maria"
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.nome, p.idade)
}

func main() {
	meet.SayHello()
	meet.Say("Go is awesome!")

	var texto string = "I love programming!"
	textoNovo := "oba"

	var idade int = 30 
	var contador int32 = 2
	var indice int64 = 1234567890
	var i2 int8 = 127

	var pi float32 = 3.14
	var raio float64 = 1.23456789012345

	var booleano bool = true

	var maior bool = idade > 18

	var array [2] string

	array[0] = "Gustavo"
	array[1] = "Nunes"

	var gavetas []string
	gavetas = append(gavetas, "caneta", "azul caneta")
	gavetas = append(gavetas, "lapis")
	gavetas = append(gavetas, "borracha")
	gavetas = append(gavetas, "caderno")
	fmt.Println(gavetas[:2])
	fmt.Println(gavetas)
	gavetas = gavetas[:2]
	fmt.Println(gavetas)

	fmt.Println(array)

	fmt.Println(texto, textoNovo, nome, idade, contador, indice, i2, pi, raio, booleano, maior)

	var pessoas = map[string]int {
		"João": 25,
		"Maria": 30,
		"Pedro": 20,
	}

	fmt.Println(pessoas)
	fmt.Println(pessoas["Maria"])
	delete(pessoas, "Pedro")

	if idade, ok := pessoas["João"]; ok {
		fmt.Println("Idade do João:", idade)
	} else {
		fmt.Println("João não encontrado")
	}

	nota := 7
	if nota >= 7 {
		fmt.Println("Aprovado")
	} else if nota >= 5 {
		fmt.Println("Recuperação")
	} else {
		fmt.Println("Reprovado")
	}

	switch nota {
	case 10:
		fmt.Println("Nota máxima")
	case 7, 8, 9:
		fmt.Println("Aprovado")
	case 5, 6:
		fmt.Println("Recuperação")
	default:
		fmt.Println("Reprovado")
	}

	sum := 0
	for i:= 0; i <= 10; i++ {
		sum += 1
	}

	fmt.Println("Soma:", sum)

	nums := []int {1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Println("Index:", i, "Value:", nums[i])
	}

	for i, num := range nums {
		fmt.Println("Index:", i, "Value:", num)
	}

	clienteJoao := Cliente {
		nome: "João",
		idade: 25,
		endereco: "Rua A, 123",
	}

	fmt.Println("Cliente:", clienteJoao)
	fmt.Println("Nome:", clienteJoao.nome)
	fmt.Println("Idade:", clienteJoao.idade)
	fmt.Println("Endereço:", clienteJoao.endereco)

	clienteJoao.idade = 32

	fmt.Println("Idade atualizada:", clienteJoao.idade)
	
	resultado := soma(10, 20)
	fmt.Println("Resultado da soma:", resultado)

	tax := func(valor float64) float64 {
		return valor * 0.1
	}
	fmt.Println("Taxa sobre 100.0:", tax(100.0))

	p := Pessoa {
		nome: "Ana",
		idade: 28,
	}

	fmt.Println("Pessoa antes:", p)

	p.Apresentar()
	
	fmt.Println("Pessoa depois:", p)
	
	var p1 Pessoa = Pessoa{ nome: "Carlos", idade: 35}

	var p3 *Pessoa = &p1
	fmt.Println("Ponteiro:", &p3)

	p3.nome = "Roberto"

	fmt.Println("Pessoa 1:", p1)
	fmt.Println("Pessoa 3:", *p3)

}

func soma(a, b int) int {
	return a + b
}