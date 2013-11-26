package models

type Player struct {
	id      int
	poker   []int //扑克
	chip    int   //筹码
	cool    bool  //冻结
	betChip int   //下注的数
	fold    bool  //放弃
}

func (this *Player) init() {
	this.chip = 10000
}

//起手牌
func (this *Player) startingHand(poker []int) {
	this.poker = poker
}

//下注
func (this *Player) Bet(num int) {
	if this.chip < num {
		return
	}
	this.chip -= num
	this.betChip += num
}

func (this *Player) GetBet() int {
	return this.betChip
}

func (this *Player) GetChip() int {
	return this.chip
}

func (this *Player) GetFold() bool {
	return this.fold
}

func (this *Player) GetCool() bool {
	return this.cool
}

//设置行动
func (this *Player) SetCool(b bool) {
	this.cool = b
}

func (this *Player) Fold() {
	this.fold = true
}
