package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var command string
var params = []string{}

type item struct {
	name   string
	status bool
}

type room struct {
	name        string
	isEnterable bool
	isLoocked   bool
	items       []item
}

type character struct {
	bag       []string
	status    int
	bagStatus bool
}

var cup = item{
	name:   "кружка",
	status: true,
}

var keys = item{
	name:   "ключи",
	status: true,
}

var workbooks = item{
	name:   "конспекты",
	status: true,
}

var bag = item{
	name:   "рюкзак",
	status: true,
}

var kitchen = room{
	name:        "кухня",
	isEnterable: true,
	isLoocked:   false,
	items:       []item{cup},
}

var corridor = room{
	name:        "коридор",
	isEnterable: true,
	isLoocked:   false,
}

var myRoom = room{
	name:        "комната",
	isEnterable: true,
	isLoocked:   false,
	items: []item{
		keys,
		workbooks,
		bag,
	},
}

var outside = room{
	name:        "улица",
	isEnterable: true,
	isLoocked:   false,
}

var gameMap = []room{
	kitchen,
	corridor,
	myRoom,
	outside,
}

var ch = character{
	bag:       []string{},
	status:    0,
	bagStatus: false,
}

func main() {
	initGame()
}

func initGame() {

	for {
		fmt.Println("\nВведите команду с вашим действием: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		if line == "конец" {
			break
		} else {
			fmt.Println(handleCommand(line))
			continue
		}

	}

}

func (ch *character) updateCharPosition(i int) {
	ch.status = i
}

func (ch *character) getCharPosition() int {
	position := ch.status
	return position
}

func (ch character) getBagStatus() bool {
	return ch.bagStatus
}
func (ch *character) changeBagStatus() {
	ch.bagStatus = true
}

func lookAround(i int) string {
	var result string
	var table = "на столе "
	switch gameMap[i].name {
	case "кухня":
		result = "ты находишься на" + " " + gameMap[i].name + ", "

		for _, value := range gameMap[i].items {
			// fmt.Printf("Value of %v is %v\n\n", value.name, value.status)
			if value.status {
				table += value.name
			}
		}
		result += ". надо собрать рюкзак и идти в универ." + " "
		result += "можно пройти -" + " " + gameMap[i+1].name
	case "коридор":
		result = "ничего интересного. можно пройти -" + gameMap[i-1].name + gameMap[i+1].name + gameMap[i+2].name
	case "комната":
		result = "на столе:" + gameMap[i].items[0].name + " " + gameMap[i].items[1].name + " " + "на стуле - " + gameMap[i].items[2].name + " " + "можно пройти - " + gameMap[i-1].name
	}

	return result
}

func goTo() string {
	var result string

	for i, name := range gameMap {
		// fmt.Printf("Character position before update is %v\n", ch.getCharPosition())
		// fmt.Printf("User entered is %v\n", params[0])
		// fmt.Printf("Loop value is %v\n", name.name)
		// fmt.Printf("i is %v\n", i)
		// fmt.Printf("Difference between i and char postion is %v\n", i-ch.getCharPosition())

		if name.name == params[0] && i-ch.getCharPosition() == 1 || name.name == params[0] && i-ch.getCharPosition() == -1 {
			//fmt.Printf("Name.name is %v\n\n", name.name)
			//	fmt.Printf("Params[0] is %v\n\n", params[0])

			ch.updateCharPosition(i)
			//	fmt.Printf("Character position after update is %v\n\n", ch.getCharPosition())
			result = lookAround(i)
			break
		} else if name.name == params[0] && i-ch.getCharPosition() == 0 {
			result = "вы уже находитесь в " + " " + gameMap[i].name

		} else if name.name == params[0] && i-ch.getCharPosition() > 1 || name.name == params[0] && i-ch.getCharPosition() < -1 {
			result = "нет пути в" + " " + gameMap[i].name

		}
	}
	cleanParams()
	return result
}

func takeItem() string {
	var result string
	if !ch.getBagStatus() {
		ch.changeBagStatus()
		changeItemStatus()
		result = "вы одели: рюкзак"
	} else {
		for i, val := range gameMap[ch.getCharPosition()].items {
			if val.name == params[0] {
				gameMap[ch.getCharPosition()].items[i].status = false
				ch.bag = append(ch.bag, gameMap[ch.getCharPosition()].items[i].name)
				//	fmt.Printf("bag contains %v\n\n", ch.bag)
				result += gameMap[ch.getCharPosition()].items[i].name
			}
		}

	}
	cleanParams()
	return result
}

func changeItemStatus() {
	for i, val := range gameMap[ch.getCharPosition()].items {
		if val.name == params[0] {
			gameMap[ch.getCharPosition()].items[i].status = false
		}
	}
}

func checkUserTextLength(userText []string) bool {
	if len(userText) > 1 {
		return true
	}
	return false
}

func getTheCommand(userText []string) {
	command = userText[0]
}

func getCommandsParams(userText []string) {
	for i, word := range userText {
		if i > 0 {
			//fmt.Printf("Функция 'getCommandsParams' \nИндекс %v Переданный параметр %v\n\n", i, word)
			params = append(params, word)
			//	fmt.Printf("Функция 'getCommandsParams' значение параметра %v\n\n", params)
		}
	}
}

func cleanParams() {
	//fmt.Printf("Params before cleaning %v\n\n", params)
	params = nil
	// fmt.Printf("Params after cleaning %v\n\n", params)
}

func handleCommand(userText string) string {
	getTheCommand(strings.Split(userText, " "))
	commands := map[string]interface{}{
		"осмотреться": lookAround,
		"идти":        goTo,
		"взять":       takeItem,
	}
	for key, _ := range commands {
		if key == command {
			if checkUserTextLength(strings.Split(userText, " ")) {
				getCommandsParams(strings.Split(userText, " "))
				return commands[command].(func() string)()
			} else if !checkUserTextLength(strings.Split(userText, " ")) {
				return commands[command].(func(int) string)(ch.getCharPosition())
			}
		}
	}

	return "неизвестная команда"
}
