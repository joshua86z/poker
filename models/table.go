package models

import (
	//	"fmt"
	"github.com/fhbzyc/poker/libs/array"
)

type Table struct {
	Step           int      //步骤  0没开始 1最先三张牌 2第四张牌 3第五张牌
	Players        []Player //玩家最多9人
	Poker          []int    //扑克
	CommunityCards []int    //公共牌
	PlayerId       int      //现在玩家
	SumChip        int      //本局筹码
	MaxChip        int      //当然玩家最大的下注数
}

//初始化
func (this *Table) Init() {
	this.Step = 0
	this.Players = []Player{}
	this.Poker = []int{}
	this.CommunityCards = []int{}
}

//下一步
func (this *Table) Next() {
	switch this.Step {
	case 0:
		//		this.shuffle()
		//		this.flopCards()
		this.Step += 1
	case 1:
		//		this.turnCards()
		this.Step += 1
	case 2:
		//		this.riverCards()
		this.Step += 1
	case 3:
		//		亮牌
		this.Step += 1
	case 4:
		//		下一局
		this.Step = 0
	}
}

//
func (this *Table) GetStep() int {
	return this.Step
}

func (this *Table) GetCommunityCards() []int {
	return this.CommunityCards
}

//
func (this *Table) GetPlayers() []Player {
	return this.Players
}

//添加玩家
func (this *Table) GetPlayerById(id int) Player {
	for _, player := range this.Players {
		if player.id == id {
			return player
		}
	}
	var player Player
	return player
}

//当前玩家ID
func (this *Table) GetPlayerId() int {
	return this.PlayerId
}

//下一个下注的玩家
func (this *Table) NextBetPlayer(index int) int {

	var num int

	for i := 0; i < len(this.Players); i++ {
		if this.Players[i].id != 0 &&
			this.Players[i].GetFold() == false &&
			this.Players[i].GetChip() > 0 {
			num++
		}
	}

	if num <= 1 {
		return -1
	}

	for i := index + 1; i < len(this.Players); i++ {
		if this.Players[i].id != 0 &&
			this.Players[i].GetFold() == false &&
			this.Players[i].GetChip() > 0 {
			this.Players[i].SetCool(true)
			this.Players[index].SetCool(false)
			return i
		}
	}

	for i := 0; i < index; i++ {
		if this.Players[i].id != 0 &&
			this.Players[i].GetFold() == false &&
			this.Players[i].GetChip() > 0 {
			this.Players[i].SetCool(true)
			this.Players[index].SetCool(false)
			return i
		}
	}

	return -1
}

func (this *Table) GetSumChip() int {
	return this.SumChip
}

func (this *Table) GetMaxChip() int {
	return this.MaxChip
}

//洗牌
func (this *Table) shuffle() {
	poker := GetPoker()
	array.IntShuffle(&poker)
	this.Poker = poker
}

//发3张公共牌
func (this *Table) flopCards() {
	this.CommunityCards = this.Poker[0:3]
	this.Poker = this.Poker[3:len(this.Poker)]
}

//第4张公共牌
func (this *Table) turnCards() {
	this.CommunityCards = append(this.CommunityCards, this.Poker[0])
	this.Poker = this.Poker[1:len(this.Poker)]
}

//第5张公共牌
func (this *Table) riverCards() {
	this.CommunityCards = append(this.CommunityCards, this.Poker[0])
	this.Poker = this.Poker[1:len(this.Poker)]
}

//起手牌
func (this *Table) startingHand() []int {

	res := this.Poker[0:2]
	this.Poker = this.Poker[2:len(this.Poker)]

	return res
}

//添加玩家
func (this *Table) addPlayer(player Player) {

	this.Players = append(this.Players, player)
}

//当前玩家ID
func (this *Table) playerId(id int) {
	this.PlayerId = id
}

//增加本局的筹码
func (this *Table) addSumChip(num int) {
	this.SumChip += num
}

//增加本轮的筹码
func (this *Table) addMaxChip(num int) {
	this.MaxChip = num
}
