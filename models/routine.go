package models

import (
	"code.google.com/p/go.net/websocket"

	"encoding/json"

	"fmt"

	//	"github.com/fhbzyc/poker/libs/array"
	//	"math/rand"
	"strconv"

	"time"
)

var WsList []*websocket.Conn
var chanList []chan string

var WsListNum int = 0

var table Table

func Run(reply string, ws *websocket.Conn) {

	var GetData GetData

	err := json.Unmarshal([]byte(reply), &GetData)

	if err == nil {
		switch GetData.Action {
		case "bet":
			//这里还缺个类型验证
			num, err := strconv.Atoi(GetData.Data.(string))
			if err != nil {
				return
			}
			Id := GetId(ws)

			for i := range table.Players {
				if table.Players[i].id == Id {
					if !table.Players[i].cool {
						//没到自己下注的时候
						return
					}
					if table.Players[i].chip < num {
						//下的注超出了自己的上限
						return
					}
					if table.Players[i].GetBet()+num < table.GetMaxChip() && table.Players[i].GetChip()-num > 0 {
						//下的注不够
						return
					}
					table.Players[i].Bet(num)
					nextIndex := table.NextPlayer(i)
					table.addMaxChip(table.Players[i].GetBet())
					table.addSumChip(num)

					if table.Players[nextIndex].GetBet() >= table.MaxChip {
						//没有任何人加注了
						table.Next()

						timer1 := time.NewTimer(time.Second * 2)
						<-timer1.C

						Play()
						return
					}
					//这里还要判断是否其他人都不能加注了
					//					table.Next()
					//					Play()

					//					fmt.Println(table)
					//					table.Players[i].cool(false)
					//					table.Players[table.NextPlayer(i)].cool(true)

					fmt.Println("桌面上筹码是", table.MaxChip)
					return
				}
			}

		}
		return
	} else {

	}
}

func GetId(ws *websocket.Conn) int {
	for i, val := range WsList {
		if val == ws {
			return i + 1
		}
	}
	return 0
}

func Play() {
	step := table.GetStep()
	switch step {
	case 0:
		startingHand()
	case 1:
		flopCards()
	case 2:
		turnCards()
	case 3:
		riverCards()
	case 5:
		step5()
	}
}

type SendData struct {
	Action string
	Data   interface{}
}

type GetData struct {
	Action string
	Data   interface{}
}

func startingHand() {

	table.Init()
	table.shuffle()

	for i, ws := range WsList {
		var player Player
		player.id = i + 1
		player.chip = 10000
		player.cool = false
		player.poker = table.startingHand()

		if i == 0 {
			player.cool = true
		}

		table.addPlayer(player)

		msg, _ := json.Marshal(SendData{Action: "startingHand", Data: player.poker})

		if ws != nil {
			err := websocket.Message.Send(ws, string(msg))
			if err != nil {
				WsList[i] = nil
			}
		} else {

		}

		//		return
	}

	//	table.Next()
}

func flopCards() {

	//	table.Next()

	table.flopCards()

	communityCards := table.GetCommunityCards()

	result, _ := json.Marshal(SendData{Action: "flopCards", Data: communityCards})

	send(string(result))
}

func turnCards() {

	//	table.Next()

	table.turnCards()

	communityCards := table.GetCommunityCards()

	result, _ := json.Marshal(SendData{Action: "turnCards", Data: communityCards[3:len(communityCards)]})

	send(string(result))
}

func riverCards() {

	//	table.Next()

	table.riverCards()

	communityCards := table.GetCommunityCards()

	result, _ := json.Marshal(SendData{Action: "riverCards", Data: communityCards[4:len(communityCards)]})

	send(string(result))
}

func step5() {

}

//群发数据
func send(msg string) {
	var err error
	for i := 0; i < len(WsList); i++ {
		if WsList[i] == nil {
			continue
		}

		if err = websocket.Message.Send(WsList[i], msg); err != nil {
			WsList[i] = nil
		} else {

		}
	}
}

/*
//牌桌
type Table struct {
	Step           int       //步骤  0没开始 1最先三张牌 2第四张牌 3第五张牌
	Players        [9]player //玩家最多9人
	Pokers         []int     //扑克
	CommunityCards []int     //公共牌
}

//洗牌
func initPokers() []int {

	poker := []int{
		2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, //红桃
		102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, //方快
		1002, 1003, 1004, 1005, 1006, 1007, 1008, 1009, 1010, 1011, 1012, 1013, 1014, //梅花
		10002, 10003, 10004, 10005, 10006, 10007, 10008, 10009, 10010, 10011, 10012, 10013, 1014, //黑桃
	}

	array.IntShuffle(&poker)

	return poker
}

type player struct {
}

const WIDTH int = 40
const HEIGHT int = 40

var WsList [1000]*websocket.Conn
var chanList [1000]chan string
var Postion [1000]people
var WsListNum int = 0

type Message struct {
	Action string
	Data   interface{}
}

type people struct {
	Id        int
	X         int
	Y         int
	Color     string
	Direction string
}

type cannonball struct {
	Id int
	X  int
	Y  int
}

type squire struct {
	X int
	Y int
}

func Squire() squire {
	return squire{X: 1000, Y: 800}
}

type data struct {
	Action string
	Data   string
}



func send(msg string) {
	var err error
	for i := 0; i < len(WsList); i++ {
		if WsList[i] == nil {
			continue
		}

		if err = websocket.Message.Send(WsList[i], msg); err != nil {

		} else {

		}
	}
}

func Ready(ID int) {
	if Postion[ID].Color == "" {
		colors := [6]string{"red", "write", "black", "blue", "green", "gold"}
		ra := rand.New(rand.NewSource(time.Now().UnixNano()))
		color := colors[ra.Intn(6)]
		Postion[ID] = people{ID, 0, 0, color, "right"}
		str := move(ID, "right")
		send(str)
	}
}

func collide(id int, x int, y int, action string) bool {
	if action == "down" {

		for i := 0; i < len(Postion); i++ {
			if id != i && Postion[i].Color != "" {
				if x >= Postion[i].X && x <= Postion[i].X+WIDTH && y+HEIGHT+10 >= Postion[i].Y && y+HEIGHT+10 <= Postion[i].Y+HEIGHT {
					return true
				}
				if x+WIDTH >= Postion[i].X && x+WIDTH <= Postion[i].X+WIDTH && y+HEIGHT+10 >= Postion[i].Y && y+HEIGHT+10 <= Postion[i].Y+HEIGHT {
					return true
				}
			}
		}

	} else if action == "right" {
		for i := 0; i < len(Postion); i++ {
			if id != i && Postion[i].Color != "" {
				if y >= Postion[i].Y && y <= Postion[i].Y+HEIGHT && x+WIDTH+10 >= Postion[i].X && x+10 <= Postion[i].X+WIDTH {
					return true
				}
				if y+HEIGHT >= Postion[i].Y && y+HEIGHT <= Postion[i].Y+HEIGHT && x+WIDTH+10 >= Postion[i].X && x+10 <= Postion[i].X+WIDTH {
					return true
				}
			}
		}
	} else if action == "left" {
		for i := 0; i < len(Postion); i++ {
			if id != i && Postion[i].Color != "" {
				if y >= Postion[i].Y && y <= Postion[i].Y+HEIGHT && x-10 <= Postion[i].X+WIDTH && x-WIDTH-10 >= Postion[i].X {
					return true
				}
				if y+HEIGHT >= Postion[i].Y && y+HEIGHT <= Postion[i].Y+HEIGHT && x-10 <= Postion[i].X+WIDTH && x-WIDTH-10 >= Postion[i].X {
					return true
				}
			}
		}
	} else if action == "up" {
		for i := 0; i < len(Postion); i++ {
			if id != i && Postion[i].Color != "" {
				if x >= Postion[i].X && x <= Postion[i].X+WIDTH && y-10 <= Postion[i].Y+HEIGHT && y+HEIGHT-10 >= Postion[i].Y {
					return true
				}
				if x+WIDTH >= Postion[i].X && x+WIDTH <= Postion[i].X+WIDTH && y-10 <= Postion[i].Y+HEIGHT && y+HEIGHT-10 >= Postion[i].Y {
					return true
				}
			}
		}
	}
	return false
}

func chonghe(id int, x int, y int, action string) bool {
	if action == "down" {

		for i := 0; i < len(Postion); i++ {
			if id != i && Postion[i].Color != "" {
				if x >= Postion[i].X && x <= Postion[i].X+WIDTH && y+HEIGHT >= Postion[i].Y && y+HEIGHT <= Postion[i].Y+HEIGHT {
					return true
				}
				if x+WIDTH >= Postion[i].X && x+WIDTH <= Postion[i].X+WIDTH && y+HEIGHT >= Postion[i].Y && y+HEIGHT <= Postion[i].Y+HEIGHT {
					return true
				}
			}
		}

	} else if action == "right" {
		for i := 0; i < len(Postion); i++ {
			if id != i && Postion[i].Color != "" {
				if y >= Postion[i].Y && y <= Postion[i].Y+HEIGHT && x+WIDTH >= Postion[i].X && x <= Postion[i].X+WIDTH {
					return true
				}
				if y+HEIGHT >= Postion[i].Y && y+HEIGHT <= Postion[i].Y+HEIGHT && x+WIDTH >= Postion[i].X && x <= Postion[i].X+WIDTH {
					return true
				}
			}
		}
	} else if action == "left" {
		for i := 0; i < len(Postion); i++ {
			if id != i && Postion[i].Color != "" {
				if y >= Postion[i].Y && y <= Postion[i].Y+HEIGHT && x <= Postion[i].X+WIDTH && x-WIDTH >= Postion[i].X {
					return true
				}
				if y+HEIGHT >= Postion[i].Y && y+HEIGHT <= Postion[i].Y+HEIGHT && x <= Postion[i].X+WIDTH && x-WIDTH >= Postion[i].X {
					return true
				}
			}
		}
	} else if action == "up" {
		for i := 0; i < len(Postion); i++ {
			if id != i && Postion[i].Color != "" {
				if x >= Postion[i].X && x <= Postion[i].X+WIDTH && y <= Postion[i].Y+HEIGHT && y+HEIGHT >= Postion[i].Y {
					return true
				}
				if x+WIDTH >= Postion[i].X && x+WIDTH <= Postion[i].X+WIDTH && y <= Postion[i].Y+HEIGHT && y+HEIGHT >= Postion[i].Y {
					return true
				}
			}
		}
	}
	return false
}

func move(ID int, action string) string {

	x := Postion[ID].X
	y := Postion[ID].Y

	if collide(ID, x, y, action) {
		if !chonghe(ID, x, y, action) {
			return ""
		}

	}
	find := false
	if action == "down" {
		if Postion[ID].Y < Squire().Y-40 {
			Postion[ID].Y += 10
			find = true
		}
	} else if action == "right" {
		if Postion[ID].X < Squire().X-40 {
			Postion[ID].X += 10
			find = true
		}
	} else if action == "left" {
		if Postion[ID].X >= 10 {
			Postion[ID].X -= 10
			find = true
		}
	} else if action == "up" {
		if Postion[ID].Y >= 10 {
			Postion[ID].Y -= 10
			find = true
		}
	}

	if find {
		Postion[ID].Direction = action
		result, _ := json.Marshal(Message{Action: "move", Data: Postion[ID]})
		return string(result)
	}

	return ""
}

func shoot(p people) {
	direction := p.Direction

	x := p.X
	y := p.Y

	if direction == "down" {
		x += WIDTH / 2
		y += HEIGHT + 1
	} else if direction == "right" {
		x += WIDTH + 1
		y += HEIGHT / 2

	} else if direction == "left" {
		x -= 1
		y += HEIGHT / 2
	} else if direction == "up" {
		x += WIDTH / 2
		y -= 1
	}

	if x >= Squire().X || y >= Squire().Y || x <= 0 || y <= 0 {
		return
	}

	//	str := ""
OVER:
	for {

		if direction == "down" {
			if y < Squire().Y-1 {
				y += 1
			} else {
				break
			}

		} else if direction == "right" {
			if x < Squire().X-1 {
				x += 1
			} else {
				break
			}
		} else if direction == "left" {
			if x >= 1 {
				x -= 1
			} else {
				break
			}
		} else if direction == "up" {
			if y >= 1 {
				y -= 1
			} else {
				break
			}
		}

		result, _ := json.Marshal(Message{Action: "shoot", Data: cannonball{Id: p.Id, X: x, Y: y}})
		send(string(result))
		//		continue

		for i := 0; i < len(Postion); i++ {
			if Postion[i].Color != "" && x >= Postion[i].X && x <= Postion[i].X+WIDTH && y >= Postion[i].Y && y <= Postion[i].Y+HEIGHT {
				result, _ := json.Marshal(Message{Action: "message", Data: "NUM `" + strconv.Itoa(i) + "` 被击中了！！ "})
				send(string(result))
				break OVER

				//					send(str)
				//					Postion[i] = people{x: -100, y: -100, color: "null", direction: "null"}
				//					WsList[i] = nil
			}
		}

		time.Sleep(time.Millisecond * 2)
	}
}
*/
