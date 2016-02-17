package main

import "fmt"

type person struct {
	pts    int
	armor  int
	damage int
	mana   int
	spent  int
}

type spell struct {
	name   string
	timer  int
	cost   int
	armor  int
	damage int
	health int
	mana   int
}

func main() {

	s := make(map[string]*spell)

	s["missile"] = &spell{
		timer:  0,
		cost:   53,
		damage: 4,
	}
	s["drain"] = &spell{
		timer:  0,
		cost:   73,
		damage: 2,
		health: 2,
	}
	s["shield"] = &spell{
		timer: 6,
		cost:  113,
		armor: 7,
	}
	s["poison"] = &spell{
		timer:  6,
		cost:   173,
		damage: 3,
	}
	s["recharge"] = &spell{
		timer: 5,
		cost:  229,
		mana:  101,
	}

	p := person{
		pts:  50,
		mana: 500,
	}
	b := person{
		pts:    58,
		damage: 9,
	}
	cast := make(map[string]int)

	fmt.Println("Solution 1:", p1(p, b, s, cast, 0, false))
	fmt.Println("Solution 2:", p1(p, b, s, cast, 0, true))
}

func p1(p person, b person, s map[string]*spell, cast map[string]int, spent int, hard bool) int {
	cast2 := make(map[string]int)
	for k, v := range cast {
		cast2[k] = v
	}

	p2 := person{
		pts:    p.pts,
		damage: p.damage,
		armor:  p.armor,
		mana:   p.mana,
		spent:  p.spent,
	}
	b2 := person{
		pts:    b.pts,
		damage: b.damage,
		armor:  b.armor,
		mana:   b.mana,
		spent:  b.spent,
	}

	p = p2
	b = b2

	if hard {
		p.pts -= 1
		if p.pts <= 0 {
			return 0
		}
	}

	processEffects(&p, &b, s, cast2)

	if b.pts <= 0 {
		return p.spent
	}
	if p.mana < 53 {
		return 0
	}

	for name, sp := range s {
		p2 = person{
			pts:    p.pts,
			damage: p.damage,
			armor:  p.armor,
			mana:   p.mana,
			spent:  p.spent,
		}
		b2 = person{
			pts:    b.pts,
			damage: b.damage,
			armor:  b.armor,
			mana:   b.mana,
			spent:  b.spent,
		}

		cast3 := make(map[string]int)
		for k, v := range cast2 {
			cast3[k] = v
		}

		if _, ok := cast3[name]; ok {
			continue
		}

		if p2.mana < sp.cost {
			continue
		}
		p2.mana -= sp.cost
		p2.spent += sp.cost

		if spent > 0 && p2.spent >= spent {
			continue
		}

		if sp.timer == 0 {
			if sp.damage > 0 {
				b2.pts -= sp.damage
			}
			if sp.health > 0 {
				p2.pts += sp.health
			}
			if b2.pts <= 0 {
				if spent == 0 || spent > p.spent {
					spent = p2.spent
				}
				continue
			}
		} else {
			cast3[name] = sp.timer
		}

		processEffects(&p2, &b2, s, cast3)
		if b2.pts <= 0 {
			if spent == 0 || spent > p.spent {
				spent = p2.spent
			}
			continue
		}

		d := b2.damage - p2.armor
		if d < 1 {
			d = 1
		}
		p2.pts -= d
		if p2.pts <= 0 {
			continue
		}

		spent2 := p1(p2, b2, s, cast3, spent, hard)
		if spent2 > 0 {
			if spent == 0 || spent > spent2 {
				spent = spent2
			}
		}
	}

	return spent
}

func processEffects(p *person, b *person, s map[string]*spell, cast map[string]int) {
	p.armor = 0
	for k, v := range cast {

		sp := s[k]

		if sp.damage > 0 {
			b.pts -= sp.damage
		}
		if sp.health > 0 {
			p.pts += sp.health
		}
		if sp.armor > 0 {
			p.armor = sp.armor
		}
		if sp.mana > 0 {
			p.mana += sp.mana
		}

		if v == 1 {
			delete(cast, k)
			continue
		}
		cast[k]--
	}
}
