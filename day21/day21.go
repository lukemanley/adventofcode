package main

import (
	"fmt"
	"sort"
)

type item struct {
	cost   int
	damage int
	armor  int
}

type person struct {
	cost   int
	damage int
	armor  int
	pts    int
}

type persons []*person

func main() {

	var weapons []*item
	var armor []*item
	var rings []*item

	weapons = append(weapons, &item{8, 4, 0})
	weapons = append(weapons, &item{10, 5, 0})
	weapons = append(weapons, &item{25, 6, 0})
	weapons = append(weapons, &item{40, 7, 0})
	weapons = append(weapons, &item{74, 8, 0})

	armor = append(armor, &item{0, 0, 0})
	armor = append(armor, &item{13, 0, 1})
	armor = append(armor, &item{31, 0, 2})
	armor = append(armor, &item{53, 0, 3})
	armor = append(armor, &item{75, 0, 4})
	armor = append(armor, &item{102, 0, 5})

	rings = append(rings, &item{0, 0, 0})
	rings = append(rings, &item{25, 1, 0})
	rings = append(rings, &item{50, 2, 0})
	rings = append(rings, &item{100, 3, 0})
	rings = append(rings, &item{20, 0, 1})
	rings = append(rings, &item{40, 0, 2})
	rings = append(rings, &item{80, 0, 3})

	var options persons

	for _, w := range weapons {
		for _, a := range armor {
			for i1, r1 := range rings {
				for i2, r2 := range rings {
					if i1 > 0 && i1 >= i2 {
						continue
					}
					p := &person{}
					p.cost = w.cost + a.cost + r1.cost + r2.cost
					p.damage = w.damage + a.damage + r1.damage + r2.damage
					p.armor = w.armor + a.armor + r1.armor + r2.armor
					options = append(options, p)
				}
			}
		}
	}

	sort.Sort(options)
	fmt.Println("Solution 1:", p1(options))

	sort.Sort(sort.Reverse(options))
	fmt.Println("Solution 2:", p2(options))
}

func (s persons) Len() int {
	return len(s)
}
func (s persons) Less(i, j int) bool {
	return s[i].cost < s[j].cost
}
func (s persons) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func p1(options persons) int {
	for _, p := range options {

		p.pts = 100

		boss := person{
			damage: 8,
			armor:  2,
			pts:    100,
		}

		for {
			d := p.damage - boss.armor
			if d < 1 {
				d = 1
			}

			boss.pts -= d
			if boss.pts <= 0 {
				return p.cost
			}

			d = boss.damage - p.armor
			if d < 1 {
				d = 1
			}
			p.pts -= d
			if p.pts <= 0 {
				break
			}
		}
	}
	return -1
}
func p2(options persons) int {
	for _, p := range options {

		p.pts = 100

		boss := person{
			damage: 8,
			armor:  2,
			pts:    100,
		}

		for {
			d := p.damage - boss.armor
			if d < 1 {
				d = 1
			}

			boss.pts -= d
			if boss.pts <= 0 {
				break
			}

			d = boss.damage - p.armor
			if d < 1 {
				d = 1
			}
			p.pts -= d
			if p.pts <= 0 {
				return p.cost
			}
		}
	}
	return -1
}
