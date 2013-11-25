package models

type Player struct {
	Id    int
	Poker []int //扑克
	Chip  int   //筹码
	Cool  bool  //冻结
	Bet   int   //下注的数
}

func (this *Player) init() {
	this.Chip = 10000
}

//起手牌
func (this *Player) startingHand(poker []int) {
	this.Poker = poker
}

//下注
func (this *Player) bet(num int) {
	if this.Chip < num {
		return
	}
	this.Chip -= num
	this.Bet += num
}

//设置行动
func (this *Player) cool(b bool) {
	this.Cool = b
}

func (this *Player) getBet() int {
	return this.Bet
}
