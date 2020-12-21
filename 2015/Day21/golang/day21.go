package main

import (
	"github.com/mxschmitt/golang-combinations"
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"sort"
)

type FightList []FightResults

func (e FightList) Len() int {
    return len(e)
}

func (e FightList) Less(i, j int) bool {
    return e[i].cost > e[j].cost
}

func (e FightList) Swap(i, j int) {
    e[i], e[j] = e[j], e[i]
}

func main() {
	players := AllPlayerCombos()
	for _, player := range players {
		fmt.Println(player)
	}
	test := LoadPlayer("TestPlayer",ReadFile("test_player.txt"))
	boss := LoadPlayer("BossPlayer",ReadFile("test.txt"))
	result := Fight([]Player{test,boss})
	fmt.Println(test.name," VS ",boss.name," wins? ",result)
	fmt.Println(LoadPlayer("Boss",ReadFile("prod.txt")))
	allGames := SimulateAllGames()
	winners := Winners(allGames)
	sort.Sort(FightList(winners))
	losers := Losers(allGames)
	sort.Sort(FightList(losers))
	fmt.Println(len(winners)-1, ": ",winners[len(winners)-1])
	fmt.Println(0, ": ",losers[0])
	fmt.Println(len(losers)-1, ": ",losers[len(losers)-1])
}

type Player struct {
	name string
	hp int64
	damage int64
	ac int64
	weapon Item
	armor Item
	rings []Item
	cost int64
}

type Item struct {
	name string
	cost int64
	damage int64
	armor int64
}

type FightResults struct {
	name string
	vs string
	result bool
	cost int64
}

var Weapons []Item = []Item {
	Item{
		name:"Dagger",
		cost:8,
		damage:4,
		armor:0,
	},
	Item{
		name:"Shortsword",
		cost:10,
		damage:5,
		armor:0,
	},
	Item{
		name:"Warhammer",
		cost:25,
		damage:6,
		armor:0,
	},
	Item{
		name:"Longsword",
		cost:40,
		damage:7,
		armor:0,
	},
	Item{
		name:"Greataxe",
		cost:74,
		damage:8,
		armor:0,
	},
}

var Armors []Item = []Item {
	Item{
		name:"Leather",
		cost:13,
		damage:0,
		armor:1,
	},
	Item{
		name:"Chainmail",
		cost:31,
		damage:0,
		armor:2,
	},
	Item{
		name:"Splintmail",
		cost:53,
		damage:0,
		armor:3,
	},
	Item{
		name:"Bandedmail",
		cost:75,
		damage:0,
		armor:4,
	},
	Item{
		name:"Platemail",
		cost:102,
		damage:0,
		armor:5,
	},
}

var Rings []Item = []Item {
	Item{
		name:"Damage +1",
		cost:25,
		damage:1,
		armor:0,
	},
	Item{
		name:"Damage +2",
		cost:50,
		damage:2,
		armor:0,
	},
	Item{
		name:"Damage +3",
		cost:100,
		damage:3,
		armor:0,
	},
	Item{
		name:"Defense +1",
		cost:20,
		damage:0,
		armor:1,
	},
	Item{
		name:"Defense +2",
		cost:40,
		damage:0,
		armor:2,
	},
	Item{
		name:"Defense +3",
		cost:80,
		damage:0,
		armor:3,
	},
}

func BasePlayer(name string, hp int64) Player {
	return Player{
		name:name,
		hp:hp,
		damage:0,
		ac:0,
		rings: make([]Item,0),
		cost: 0,
	}
}

func CopyPlayer(p Player) Player {
	nrings := make([]Item,len(p.rings))
	copy(nrings,p.rings)
	return Player {
		name:p.name,
		hp:p.hp,
		damage:p.damage,
		ac:p.ac,
		rings:nrings,
		cost:p.cost,
	}
}

func LoadPlayer(name string, inputs []string) Player {
	p := BasePlayer(name,0)
	for _, input := range inputs {
		elements := strings.Split(input,": ")
		//fmt.Println(inputs,elements)
		if (len(elements) >= 2) {
			switch elements[0] {
				case "Hit Points":
					hp, _ := strconv.ParseInt(elements[1],10,64)
					p.hp = int64(hp)
				case "Damage":
					d, _ := strconv.ParseInt(elements[1],10,64)
					p.damage = int64(d)
				case "Armor":
					ac, _ := strconv.ParseInt(elements[1],10,64)
					p.ac = int64(ac)
			}

		} else {
			fmt.Println("ERROR ",input,elements)
		}
	}
	return p
}

func AllPlayersNoArmor() []Player {
	allPlayers := make([]Player,0)
	for _, weapon := range Weapons {
		allPlayers = append(allPlayers,Player{
			name: "Player " + weapon.name,
			hp: 100,
			damage: weapon.damage,
			ac: 0,
			weapon: weapon,
			cost:weapon.cost,
		})
	}
	return allPlayers
}

func AllPlayersNoRings() []Player {
	allPlayers := make([]Player,0)
	for _, weapon := range Weapons {
		for _, armor := range Armors {
			allPlayers = append(allPlayers,Player{
				name: "Player " + weapon.name + "," + armor.name,
				hp: 100,
				damage: weapon.damage,
				ac: armor.armor,
				weapon: weapon,
				armor: armor,
				cost:weapon.cost+armor.cost,
			})
		}
	}
	return allPlayers
}

func AllPlayerCombos() []Player {
	allPlayers := make([]Player,0)
	rings := make(map[string]Item,0)
	for _, ring := range Rings {
		rings[ring.name] = ring
	}
	ringCombos := AllRingCombos()
	noRingPlayers := AllPlayersNoArmor()
	noRingPlayers = append(noRingPlayers,AllPlayersNoRings()...)
	allPlayers = append(allPlayers,noRingPlayers...)
	for _, noRingPlayer := range noRingPlayers {
		for _, ringCombo := range ringCombos {
			p := CopyPlayer(noRingPlayer)
			for _, ringName := range ringCombo {
				ring := rings[ringName]
				p.rings = append(p.rings,ring)
				p.name += "," + ring.name
				p.cost += ring.cost
				p.ac += ring.armor
				p.damage += ring.damage
			}
			allPlayers = append(allPlayers,p)
		}
	}
	return allPlayers
}

func AllRingCombos() [][]string {
	rings := make([]string,0)
	for _, ring := range Rings {
		rings = append(rings,ring.name)
	}
	combos := combinations.Combinations(rings,1)
	combos = append(combos,combinations.Combinations(rings,2)...)
	return combos
}

func ReadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println(err)
	}
	return strings.Split(string(b),"\n")
}

func Fight(players []Player) FightResults {
	if len(players) != 2 {
		return FightResults{ name:players[0].name,vs:"ERROR",result:false,cost:players[0].cost }
	}

	at := 0
	for {
		
		attacker := &players[at % 2]
		defender := &players[(at + 1) % 2]
		fmt.Println("The "+ attacker.name +" deals ",attacker.damage,"-",defender.ac," = ",attacker.damage - defender.ac," damage; the ",defender.name," goes down from ",defender.hp," to ",defender.hp - (attacker.damage - defender.ac)," hit points.")
		defender.hp = defender.hp - (attacker.damage - defender.ac)
		if defender.hp <= 0 {
			fmt.Println("DONE")
			break
		}
		at++
	}
	return FightResults{ name:players[0].name,vs:players[1].name,result:(at % 2 == 0),cost:players[0].cost }
}

func Winners(frs []FightResults) []FightResults {
	winners := make([]FightResults,0)
	for _, fr := range frs {
		if fr.result {
			winners = append(winners,fr)
		}
	}
	return winners
}

func Losers(frs []FightResults) []FightResults {
	winners := make([]FightResults,0)
	for _, fr := range frs {
		if !fr.result {
			winners = append(winners,fr)
		}
	}
	return winners
}


func SimulateAllGames() []FightResults {
	fr := make([]FightResults,0)
	allPlayers := AllPlayerCombos()
	boss := LoadPlayer("Boss",ReadFile("prod.txt"))
	for _, player := range allPlayers {
		fr = append(fr,Fight([]Player{player,CopyPlayer(boss)}))
	}
	return fr
}