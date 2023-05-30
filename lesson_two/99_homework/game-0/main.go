package main

import (
	"strings"
)

var command string
var params = []string{}

type item struct {
	name   string
	status bool
	useble bool
}

type furniture struct {
	name  string
	items []string
}

type room struct {
	name        string
	isEnterable bool
	isLocked    bool
	lockedItems []string
	items       []item
	characterIn bool
	description string
	descState   bool
	connection  map[string]bool
	doorsStatus map[string]bool
	invetory    []furniture
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
	useble: false,
}

var keys = item{
	name:   "ключи",
	status: true,
	useble: false,
}

var workbooks = item{
	name:   "конспекты",
	status: true,
	useble: false,
}

var bag = item{
	name:   "рюкзак",
	status: true,
	useble: true,
}

/*ROOMS*/
//var kitchen = room{
//	name:        "кухня",
//	isEnterable: true,
//	isLocked:    false,
//	items:       []item{cup},
//	characterIn: true,
//	description: "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
//	descState:   true,
//	connection: map[string]bool{
//		"коридор": true,
//		"комната": false,
//		"улица":   false,
//	},
//}
//
//var corridor = room{
//	name:        "коридор",
//	isEnterable: true,
//	isLocked:    false,
//	characterIn: false,
//	description: "ничего интересного. можно пройти - кухня, комната, улица",
//	descState:   true,
//	connection: map[string]bool{
//		"кухня":   true,
//		"комната": true,
//		"улица":   true,
//	},
//	doorsStatus: map[string]bool{
//		"кухня":   false,
//		"комната": false,
//		"улица":   true,
//	},
//}
//
//var myRoom = room{
//	name:        "комната",
//	isEnterable: true,
//	isLocked:    false,
//	items: []item{
//		keys,
//		workbooks,
//		bag,
//	},
//	characterIn: false,
//	description: "ты в своей комнате. можно пройти - коридор",
//	descState:   true,
//	connection: map[string]bool{
//		"кухня":   false,
//		"коридор": true,
//		"улица":   false,
//	},
//	invetory: []furniture{
//		table,
//		chair,
//	},
//}
//
//var outside room{
//	name:        "улица",
//	isEnterable: true,
//	isLocked:    true,
//	lockedItems: []string{"дверь"},
//	characterIn: false,
//	description: "на улице весна. можно пройти - домой",
//	descState:   true,
//}

var gameMap []*room

var ch character

func main() {
	initGame()
}

func initGame() {

	var table = furniture{
		name: "стол",
		items: []string{
			"ключи",
			"конспекты",
		},
	}

	var chair = furniture{
		name: "стул",
		items: []string{
			"рюкзак",
		},
	}
	var kitchen = room{
		name:        "кухня",
		isEnterable: true,
		isLocked:    false,
		items:       []item{cup},
		characterIn: true,
		description: "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
		descState:   true,
		connection: map[string]bool{
			"коридор": true,
			"комната": false,
			"улица":   false,
		},
	}

	var corridor = room{
		name:        "коридор",
		isEnterable: true,
		isLocked:    false,
		characterIn: false,
		description: "ничего интересного. можно пройти - кухня, комната, улица",
		descState:   true,
		connection: map[string]bool{
			"кухня":   true,
			"комната": true,
			"улица":   true,
		},
		doorsStatus: map[string]bool{
			"кухня":   false,
			"комната": false,
			"улица":   true,
		},
	}

	var myRoom = room{
		name:        "комната",
		isEnterable: true,
		isLocked:    false,
		items: []item{
			keys,
			workbooks,
			bag,
		},
		characterIn: false,
		description: "ты в своей комнате. можно пройти - коридор",
		descState:   true,
		connection: map[string]bool{
			"кухня":   false,
			"коридор": true,
			"улица":   false,
		},
		invetory: []furniture{
			table,
			chair,
		},
	}

	var outside = room{
		name:        "улица",
		isEnterable: true,
		isLocked:    true,
		lockedItems: []string{"дверь"},
		characterIn: false,
		description: "на улице весна. можно пройти - домой",
		descState:   true,
	}
	gameMap = []*room{
		&kitchen,
		&corridor,
		&myRoom,
		&outside,
	}

	ch = character{
		bag:        []string{},
		status:     0,
		bagAvaible: false,
		bagPutOn:   false,
	}
	//testCommands := []string{
	//	"осмотреться",
	//	"идти коридор",
	//	"идти комната",
	//	"осмотреться",
	//	"одеть рюкзак",
	//	"взять ключи",
	//	"взять конспекты",
	//	"идти коридор",
	//	"применить ключи дверь",
	//	"идти улица",
	//}

	//testCommands2 := []string{
	//	"осмотреться",
	//	"завтракать",
	//	"идти комната",
	//	"идти коридор",
	//	"применить ключи дверь",
	//	"идти комната",
	//	"осмотреться",
	//	"взять ключи",
	//	"одеть рюкзак",
	//	"осмотреться",
	//	"взять ключи",
	//	"взять телефон",
	//	"взять ключи",
	//	"осмотреться",
	//	"взять конспекты",
	//	"осмотреться",
	//	"идти коридор",
	//	"идти кухня",
	//	"осмотреться",
	//	"идти коридор",
	//	"идти улица",
	//	"применить ключи дверь",
	//	"применить телефон шкаф",
	//	"применить ключи шкаф",
	//	"идти улица",
	//}
	//
	//for _, textCommand := range testCommands2 {
	//	fmt.Printf("Команда: %v\n", textCommand)
	//	result := handleCommand(textCommand)
	//	fmt.Printf("Ответ: %v\n\n", result)
	//}
	// for {
	// 	fmt.Println("\nВведите команду с вашим действием: ")
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	scanner.Scan()
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// 	result := handleCommand(line)
	// 	fmt.Println(result)
	// 	if result == "на улице весна. можно пройти - домой" {
	// 		break
	// 	} else {
	// 		continue
	// 	}

	// 	// if handleCommand(line) == "ничего интересного. можно пройти - кухня, комната, улица" {
	// 	// 	fmt.Println(handleCommand(line))
	// 	// 	break
	// 	// }
	// 	//
	// }

}

/* CHARACTER */

func (ch *character) updateCharPosition(i int) {
	ch.status = i
}

func (ch *character) getCharPosition() int {
	// position := ch.status
	return ch.status
}

func (ch *character) getBagAvaiability() bool {
	return ch.bagAvaible
}

func (ch *character) getBagState() bool {
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

func (rm *room) changeLockStatus() {
	rm.isLocked = false
}

func (fn *furniture) deleteItem() {
	for i, v := range fn.items {
		if v == params[0] {
			fn.items[i] = fn.items[len(fn.items)-1]
			fn.items[len(fn.items)-1] = ""
			fn.items = fn.items[:len(fn.items)-1]
		}
	}
}

/* COMMANDS */

func lookAround() string {

	var result string

	roomIn := gameMap[ch.getCharPosition()]

	// Проверка если статус у описания по умолчанию комнаты true
	// то вернуть текст из поля description
	if roomIn.descState {
		result = roomIn.description
	}

	if roomIn.name == "кухня" {
		if ch.bagAvaible {
			result = "ты находишься на кухне, на столе чай, надо идти в универ. можно пройти - коридор"
		}
	}
	if roomIn.name == "комната" {
		if len(roomIn.invetory[0].items) > 0 && len(roomIn.invetory[1].items) > 0 {
			sum := 0
			var itemsStr string
			for _, v := range roomIn.invetory[0].items {

				if sum == len(roomIn.invetory[0].items)-1 {
					itemsStr += " " + v
				} else {
					itemsStr += " " + v + ","
				}
				sum += 1
			}

			result = "на столе:" + itemsStr
			result += "," + " " + "на стуле -"
			itemsStr = ""
			sum = 0

			for _, v := range roomIn.invetory[1].items {

				if sum == len(roomIn.invetory[1].items)-1 {
					itemsStr += v
				} else {
					itemsStr += " " + v + ","
				}
				sum += 1
			}

			result += " " + itemsStr + "." + " " + "можно пройти - коридор"
		}

		if len(roomIn.invetory[0].items) == 0 && len(roomIn.invetory[1].items) > 0 {
			sum := 0
			var itemsStr string

			for _, v := range roomIn.invetory[1].items {

				if sum == len(roomIn.invetory[1].items)-1 {
					itemsStr += v
				} else {
					itemsStr += " " + v + ","
				}
				sum += 1
			}

			result = "на стуле -" + itemsStr + "." + " " + "можно пройти - коридор"
		}

		if len(roomIn.invetory[0].items) > 0 && len(roomIn.invetory[1].items) == 0 {
			sum := 0
			var itemsStr string
			for _, v := range roomIn.invetory[0].items {

				if sum == len(roomIn.invetory[0].items)-1 {
					itemsStr += " " + v
				} else {
					itemsStr += " " + v + ","
				}
				sum += 1
			}

			result = "на столе:" + itemsStr + "." + " " + "можно пройти - коридор"
		}

		if len(roomIn.invetory[0].items) == 0 && len(roomIn.invetory[1].items) == 0 && roomIn.name != "комната" {

			result = "ничего интересного." + " " + "можно пройти - коридор"
		}
		if len(roomIn.invetory[0].items) == 0 && len(roomIn.invetory[1].items) == 0 && roomIn.name == "комната" {

			result = "пустая комната." + " " + "можно пройти - коридор"
		}
	}

	return result
}

func goTo() string {
	var resRoom *room
	var result string
	var needToGo bool = false
	var roomIn *room

	charPosition := ch.getCharPosition()
	roomIn = gameMap[charPosition]

	if roomIn.name == params[0] {
		result = "вы уже находитесь в" + " " + gameMap[charPosition].name
	} else {
		for conectRoom, state := range roomIn.connection {
			if conectRoom == params[0] && state {
				roomIn.characterLeftRoom()
				needToGo = true
				break
			} else {
				result = "нет пути в" + " " + params[0]
			}

		}
		if needToGo {
			for i, room := range gameMap {
				if room.name == params[0] && !room.isLocked && room.name == "кухня" && ch.bagAvaible {
					resRoom = room
					resRoom.characterEnteredRoom()
					ch.updateCharPosition(i)
					result = "кухня, ничего интересного. можно пройти - коридор"
					break
				} else if room.name == params[0] && !room.isLocked {
					resRoom = room
					resRoom.characterEnteredRoom()
					ch.updateCharPosition(i)
					result = room.description
					break
				} else {
					result = "дверь закрыта"
				}
			}
		}
	}

	cleanParams()
	return result
}

func takeItem() string {
	var result string
	charPosition := ch.getCharPosition()
	roomIn := gameMap[charPosition]

	if !ch.bagAvaible && params[0] != "рюкзак" {
		cleanParams()
		return "некуда класть"
	}

	if !ch.bagAvaible && params[0] == "рюкзак" {
		ch.changeBagAvaiability()
		roomIn.invetory[1].deleteItem()
		cleanParams()
		return "вы одели: рюкзак"
	}

	for _, val := range roomIn.invetory[0].items {
		if val == params[0] {
			ch.bag = append(ch.bag, val)
			result = "предмет добавлен в инвентарь:"
			result += " " + val
			roomIn.invetory[0].deleteItem()
			break
		} else {
			result = "нет такого"
		}
	}

	cleanParams()
	return result
}

func checkItemInBag() bool {
	var bagItemStatus bool
	for _, bagItem := range ch.bag {
		if params[0] == bagItem {
			bagItemStatus = true
			break
		}
	}
	return bagItemStatus
}
func checkItemInRoom(roomName string) bool {
	var roomItemStatus bool
	for _, rm := range gameMap {
		if rm.name == roomName {
			for _, v := range rm.lockedItems {
				if v == params[1] {
					roomItemStatus = true
					break
				}
			}
		}
	}
	return roomItemStatus
}

func useItem() string {
	var result string

	roomIn := gameMap[ch.getCharPosition()]

	var openedRoom *room
	var closedRoom string

	if !checkItemInBag() {
		searchItem := params[0]
		cleanParams()
		return "нет предмета в инвентаре -" + " " + searchItem

	} else {
		for key, value := range roomIn.doorsStatus {
			if value {
				closedRoom = key
				break
			}
		}
	}

	if !checkItemInRoom(closedRoom) {
		cleanParams()
		return "не к чему применить"
	} else {
		for _, rm := range gameMap {
			if rm.name == closedRoom {
				openedRoom = rm
				break
			}
		}
	}

	if checkItemInBag() && checkItemInRoom(closedRoom) {
		openedRoom.changeLockStatus()
		result = "дверь открыта"
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
			params = append(params, word)
		}
	}
}

func cleanParams() {
	params = nil
}

func changeItemStatus() {
	for i, val := range gameMap[ch.getCharPosition()].items {
		if val.name == params[0] {
			gameMap[ch.getCharPosition()].items[i].status = false
		}
	}
}

func handleCommand(userText string) string {
	var result string

	getTheCommand(strings.Split(userText, " "))
	commands := map[string]interface{}{
		"осмотреться": lookAround,
		"идти":        goTo,
		"взять":       takeItem,
		"одеть":       takeItem,
		"применить":   useItem,
	}
	for key, _ := range commands {
		if key == command {
			getCommandsParams(strings.Split(userText, " "))
			result = commands[command].(func() string)()
			break
		} else {
			result = "неизвестная команда"
		}
	}

	cleanParams()
	return result
}
