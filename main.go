package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Options struct {
	Data []string
}

type GameData struct {
	selectedTheme, wordToGuess, displayValue, alreadyGuessed string
	userGuess string
	guessCount, incorrectGuessCount int 
	incorrectGuessLimit int
}

type GameHistory struct {
	data [] GameData
}



func main() {



	 gameData := GameData{}
	 themes:= []string {"data/presidents.json" ,"data/cars.json"  ,"data/anime.json"}


	getStartInto(&gameData)


	gameData.wordToGuess =	getWordFromTheme("cars", themes)
	gameData.displayValue = createDefaultDisplay(len(gameData.wordToGuess))


	handleUserGuess(&gameData)

	handleDisplay(&gameData)



	for {
		gameComplete := !strings.Contains(gameData.displayValue, "_")
		gameOver := (gameData.incorrectGuessLimit - gameData.incorrectGuessCount) <= 0
		if gameComplete {
			fmt.Printf("Good Job! \n %v\n", gameData.wordToGuess)
			fmt.Printf("Here are your stats: %+v", gameData)
			break
			}
		if gameOver	{
			fmt.Println("Sorry, Too many incorrect attempts")
			break
		}

		handleUserGuess(&gameData)
		handleDisplay(&gameData)
		}

	}

	func handleUserGuess (game *GameData) {
		fmt.Println(game.displayValue)
		fmt.Println("Guess a letter ")
		fmt.Scan(&game.userGuess)

		//if len(game.userGuess)

		current := strings.ToUpper(game.userGuess)
		alreadyUsed := strings.Contains(game.alreadyGuessed, current)
		if alreadyUsed {
			fmt.Printf("You have already Guessed: %v \n", game.userGuess)
			 handleUserGuess(game)
		}
			game.alreadyGuessed += strings.ToUpper(game.userGuess)

	}

	func createDefaultDisplay (lengthOfWord int) string{
		defaultDisplay :=""
		for i := 0; i < lengthOfWord; i++ {
			defaultDisplay += "_"
		}
		return defaultDisplay
	}



	func handleDisplay (game *GameData)  {
		var listOfIndexes []int
		x := strings.ToUpper(game.userGuess)
		newGuess := []rune(x)[0]

		for i, char := range game.wordToGuess {
			if char == newGuess {
				listOfIndexes=	append(listOfIndexes, i)
			}
		}

		decodedDisplayVal := []rune(game.displayValue)
		correctGuess := len(listOfIndexes) > 0

		if correctGuess {
			for i, _ := range listOfIndexes {
				//range through the list of indexs and then assign each item and
				// and replace it with the guessed value
				decodedDisplayVal[listOfIndexes[i]] = newGuess
			}
			game.displayValue = string(decodedDisplayVal)
			game.guessCount ++
		}  else  {
			game.incorrectGuessCount ++
			fmt.Printf("Number of Incorrect guesses left: %v \n", game.incorrectGuessLimit - game.incorrectGuessCount)
		}

}

	func getWordFromTheme (selectedThemeStr string, options []string ) string {
	selectedWord := ""
	for _, val := range options {
		if strings.Contains(val, selectedThemeStr) {
		words, _ :=	ParseFromJson(val)
		selectedWord = strings.ToUpper(words.Data[0])
			}
		}
		return selectedWord
}

	func getStartInto (game *GameData)    {
		difficultyLevels := map[string] int {"EASY" : 10, "MEDIUM": 6, "HARD": 4}
		fmt.Println("Welcome to Hangman! \nChoose your difficulty level: \n ")

		for k, val := range difficultyLevels {
		fmt.Printf("Enter %v for %v Attempts \n", k, val)
	}
	var difficultyChoice string
	fmt.Scan(&difficultyChoice)
		game.incorrectGuessLimit = difficultyLevels[difficultyChoice]


	fmt.Printf("\n \nChoose your theme: \n Cars \n Presidents \n Anime\n")
	fmt.Scan(&game.selectedTheme)
}

	func ParseFromJson (fileName string) (Options, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	data := Options{}

	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}
	
	return data, err
	
}

	func ParseToJson(data Options, fileName string) bool {
	jsonFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fileName, jsonFile, 0644)
	if err != nil {
		panic(err)
	}

	return true
}