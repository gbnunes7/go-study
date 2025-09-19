package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Question struct {
	Text string
	Options []string
	Answer int
}

type GameState struct {
	Name string
	Points string
	Questions []Question
}

func (game *GameState)Init() {
	fmt.Println("Iniciando o jogo...")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite seu nome: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		panic("Erro ao ler o nome: " + err.Error())
	}
	game.Name = name

	fmt.Printf("Bem-vindo, %s Vamos começar o quiz.\n", game.Name)
}

func (game *GameState) ProcessCSV() {
	 file, err := os.Open("quizgo.csv")

	 if err != nil {
		 panic("Erro ao abrir o arquivo: " + err.Error())
	 }

	 defer file.Close() // feito por ultimo independente

	 reader := csv.NewReader(file)
	 records, err := reader.ReadAll()

	 if err != nil {
		panic("Erro ao ler CSV")
	 }

	 for index, record := range records {
		if index > 0 {
			correctAnswer, _ := toInt(record[5])
			question := Question {
				Text: record[0],
				Options: record[1:5],
				Answer: correctAnswer,
			}
			game.Questions = append(game.Questions, question)
		}
	 }
}

func (game *GameState) RunGame() {
	for index, question := range game.Questions {
		fmt.Printf("\033[33m %d. %s \033[0m\n", index+1, question.Text)
		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Println("Digite uma alternativa")

		var answer int
		var err error 

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')
		
			answer, err = toInt(read[:len(read)-1])

			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			break
		}

		if answer == question.Answer {
			fmt.Println("Resposta correta!")
			game.Points = "10"
		} else {
			fmt.Printf("Resposta incorreta! A resposta correta é %d\n", question.Answer)
			game.Points = "0"
		}


	}
}

func main() {
	game := &GameState{}
	go game.ProcessCSV()
	game.Init()
	game.RunGame()

	fmt.Printf("Obrigado por jogar, %s! Você fez %s pontos.\n", game.Name, game.Points)
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("erro ao converter string para int: %w", err)
	}
	return i, nil
}