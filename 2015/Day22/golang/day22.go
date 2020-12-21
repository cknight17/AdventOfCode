package main

import (
	"fmt"
)

func main() {
	r := Round{
		player: Player{
			hp:50,
			armor:0,
			mana:500,
			manaSpent:0,
			poisonLeft:0,
			shieldLeft:0,
			rechargeLeft:0,
		},
		boss:Boss{hp:71,damage:10},
		action:"Magic Missle",
		complete:false,
		valid:true,
		win:false,
	}
	results := PlayGame([]Round{r})
	for i, ro := range results {
		lr := ro[len(ro)-1]
		if lr.complete && lr.valid && lr.win {
			fmt.Println(i,": ",lr)
		}
	}
}

type Player struct {
	hp int64
	armor int64
	mana int64
	manaSpent int64
	poisonLeft int64
	shieldLeft int64
	rechargeLeft int64
}

type Boss struct {
	hp int64
	damage int64
}

type Round struct {
	player Player
	boss Boss
	action string
	complete bool
	valid bool
	win bool
}

var actions []string = []string{"Magic Missile","Drain","Shield","Poison","Recharge"}

func ProcessRound(round Round) Round {
	//fmt.Println(round)
	// Process player action
	switch round.action {
		case "Magic Missile":
			round.boss.hp = round.boss.hp - 4
			round.player.mana = round.player.mana - 53
			round.player.manaSpent += 53
			//fmt.Println("MM: ",round)
		case "Drain":
			round.boss.hp = round.boss.hp - 2
			round.player.hp += 2
			round.player.mana = round.player.mana - 73
			round.player.manaSpent += 73
			//fmt.Println("D: ",round)
		case "Shield":
			if round.player.shieldLeft == 0 {
				round.player.shieldLeft = 6
				round.player.mana = round.player.mana - 113
				round.player.manaSpent += 113
				//fmt.Println("SH: ",round)
			} else {
				round.complete = true
				round.valid = false
				return round
			}
		case "Poison":
			if round.player.poisonLeft == 0 {
				round.player.poisonLeft = 6
				round.player.mana = round.player.mana - 173
				round.player.manaSpent += 173
				//fmt.Println("P: ",round)
			} else {
				round.complete = true
				round.valid = false
				return round
			}
		case "Recharge":
			if round.player.rechargeLeft == 0 {
				round.player.rechargeLeft = 5
				round.player.mana = round.player.mana - 229
				round.player.manaSpent += 229
				//fmt.Println("R: ",round)
			} else {
				round.complete = true
				round.valid = false
				return round
			}
	}
	// Process effects first
	if round.player.poisonLeft > 0 {
		round.boss.hp = round.boss.hp - 3
		round.player.poisonLeft--
	}
	if round.player.shieldLeft > 0 {
		round.player.armor = 7
		round.player.shieldLeft--
	} else {
		round.player.armor = 0
	}
	if round.player.rechargeLeft > 0 {
		round.player.mana += 101
		round.player.rechargeLeft--
	}
	// Did player run out of mana and die
	if round.player.mana <= 0 {
		round.complete = true
		round.valid = true
		round.win = false
		//fmt.Println("NO MANA")
		return round
	}
	// Is boss dead
	if round.boss.hp <= 0 {
		round.complete = true
		round.valid = true
		round.win = true
		//fmt.Println("WINNER")
		return round
	}
	// Process boss action
	round.player.hp = round.player.hp - (round.boss.damage - round.player.armor)
	// Did player die
	if round.player.hp <= 0 {
		round.complete = true
		round.valid = true
		round.win = false
		//fmt.Println("BLARGH")
		return round
	}
	return round
}

func PlayGame(progress []Round) [][]Round {
	allPaths := make([][]Round,0)
	for _, action := range actions {
		nprogress := make([]Round,len(progress))
		copy(nprogress,progress)
		newRound := nprogress[len(nprogress)-1]
		newRound.action = action
		newRound = ProcessRound(newRound)
		//fmt.Println("AFTER: ",newRound)
		nprogress = append(nprogress,newRound)
		if newRound.complete && newRound.valid {
			allPaths = append(allPaths,nprogress)
		} else if newRound.valid {
			allPaths = append(allPaths,PlayGame(nprogress)...)
		} else {
			//fmt.Println("INVALID: ",newRound,nprogress)
		}
	}
	return allPaths
}