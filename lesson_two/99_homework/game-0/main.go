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
	name           string
	isEnterable    bool
	isLoocked      bool
	items          []item
	characterIn    bool
	defDescription string
	defDescState   bool
	connection     map[string]bool
}

type character struct {
	bag        []string
	status     int
	bagAvaible bool
	bagPutOn   bool
}

/*ITEMS*/
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

/*ROOMS*/
var kitchen = room{
	name:           "кухня",
	isEnterable:    true,
	isLoocked:      false,
	items:          []item{cup},
	characterIn:    true,
	defDescription: "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
	defDescState:   true,
	connection: map[string]bool{
		"коридор": true,
		"комната": false,
		"улица":   false,
	},
}

var corridor = room{
	name:           "коридор",
	isEnterable:    true,
	isLoocked:      false,
	characterIn:    false,
	defDescription: "ничего интересного. можно пройти - кухня, комната, улица",
	defDescState:   true,
	connection: map[string]bool{
		"кухня":   true,
		"комната": true,
		"улица":   true,
	},
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
	characterIn:    false,
	defDescription: "ты в своей комнате. можно пройти - коридор",
	defDescState:   true,
	connection: map[string]bool{
		"кухня":   false,
		"коридор": true,
		"улица":   false,
	},
}

var outside = room{
	name:           "улица",
	isEnterable:    true,
	isLoocked:      false,
	characterIn:    false,
	defDescription: "на улице весна. можно пройти - домой",
	defDescState:   true,
	connection: map[string]bool{
		"кухня":   false,
		"коридор": true,
		"комната": false,
	},
}

var gameMap = []room{
	kitchen,
	corridor,
	myRoom,
	outside,
}

var ch = character{
	bag:        []string{},
	status:     0,
	bagAvaible: false,
	bagPutOn:   false,
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

/* CHARACTER */

func (ch *character) updateCharPosition(i int) {
	ch.status = i
}

func (ch *character) getCharPosition() int {
	// position := ch.status
	return ch.status
}

func (ch character) getBagAvaiability() bool {
	return ch.bagAvaible
}

func (ch character) getBagState() bool {
	return ch.bagPutOn
}

func (ch *character) changeBagAvaiability() {
	ch.bagAvaible = true
}

func (ch *character) changeBagState() {
	ch.bagPutOn = true
}

func (rm *room) characterEnteredRoom() {
	rm.characterIn = true
}

func (rm *room) characterLeftRoom() {
	rm.characterIn = false
}

/* COMMANDS */

func lookAround(i int) string {
	var result string
	// var table = "на столе "
	// switch gameMap[i].name {
	// case "кухня":
	// 	result = "ты находишься на" + " " + gameMap[i].name + ", "

	// 	for _, value := range gameMap[i].items {
	// 		// fmt.Printf("Value of %v is %v\n\n", value.name, value.status)
	// 		if value.status {
	// 			table += value.name
	// 		}
	// 	}
	// 	result += ". надо собрать рюкзак и идти в универ." + " "
	// 	result += "можно пройти -" + " " + gameMap[i+1].name
	// case "коридор":
	// 	result = "ничего интересного. можно пройти -" + gameMap[i-1].name + gameMap[i+1].name + gameMap[i+2].name
	// case "комната":
	// 	result = "на столе:" + gameMap[i].items[0].name + " " + gameMap[i].items[1].name + " " + "на стуле - " + gameMap[i].items[2].name + " " + "можно пройти - " + gameMap[i-1].name
	// }
	for _, room := range gameMap {
		if room.characterIn {
			if room.defDescState {
				result = room.defDescription
			}
		}
	}
	return result
}

func goTo() string {
	var result string
	var num int
	for _, room := range gameMap {
		fmt.Printf("Название комнты %v\n", room.name)
		fmt.Printf("Состояние игрока в этой комнате %v\n\n", room.characterIn)
	}
	for i, room := range gameMap {
		if room.characterIn {
			num = i
		}
	}

	if gameMap[num].name == params[0] {
		result = "вы уже находитесь в" + " " + gameMap[num].name
	}
	fmt.Printf("Комната из которой выходим %v\n", gameMap[num].name)
	for room, state := range gameMap[num].connection {
		fmt.Printf("room %v\n", room)
		fmt.Printf("state %v\n", state)
		if room == params[0] && state {
			fmt.Println("статус положительный")
			fmt.Printf("Статус игрока %v в комнате %v до его изменения\n", gameMap[num].characterIn, gameMap[num].name)
			gameMap[num].characterLeftRoom()
			fmt.Printf("Статус игрока %v в комнате %v после его изменения\n", gameMap[num].characterIn, gameMap[num].name)
			for i, room := range gameMap {
				if room.name == params[0] {
					fmt.Printf("Статус игрока %v в комнате %v до его изменения\n", room.characterIn, room.name)
					room.characterEnteredRoom()
					fmt.Printf("Статус игрока %v в комнате %v после его изменения\n", room.characterIn, room.name)
					fmt.Printf("Позиция игрока до ее изменения %v\n", ch.getCharPosition())
					ch.updateCharPosition(i)
					fmt.Printf("Позиция игрока после ее изменения %v\n\n", ch.getCharPosition())
					result = room.defDescription
				}
			}
			break
		} else {
			result = "нет пути в" + " " + params[0]
		}
	}

	cleanParams()

	return result
}

func takeItem() string {
	var result string
	if !ch.getBagAvaiability() {
		ch.changeBagAvaiability()
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

/* HELP COMMANDS */

func getTheCommand(userText []string) {
	command = userText[0]
}

func checkUserTextLength(userText []string) bool {
	if len(userText) > 1 {
		return true
	}
	return false
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

func changeItemStatus() {
	for i, val := range gameMap[ch.getCharPosition()].items {
		if val.name == params[0] {
			gameMap[ch.getCharPosition()].items[i].status = false
		}
	}
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
