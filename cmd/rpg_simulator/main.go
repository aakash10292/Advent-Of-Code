package main

import "math"

type item struct{
	cost int
	damage int
	armor int
}

func getShotsNeeded(hitPoints int, damage int, armor int) int {
	if (damage - armor) <= 0 {
		// can never defeat
		return math.MaxInt32
	}
	shotsNeeded := hitPoints/(damage-armor)
	if shotsNeeded % (damage-armor) != 0 {
		shotsNeeded++
	}
	return shotsNeeded
}

func main(){
	var Weapons []item
	Weapons = append(Weapons, item{8,4,0})
	Weapons = append(Weapons, item{10,5,0})
	Weapons = append(Weapons, item{25,6,0})
	Weapons = append(Weapons, item{40,7,0})
	Weapons = append(Weapons, item{74,8,0})

	var Armor []item
	Armor = append(Armor, item{13,0,1})
	Armor = append(Armor, item{31,0,2})
	Armor = append(Armor, item{53,0,3})
	Armor = append(Armor, item{75,0,4})
	Armor = append(Armor, item{102,0,5})

	var Rings []item
	Rings = append(Rings, item{25,1,0})
	Rings = append(Rings, item{50,2,0})
	Rings = append(Rings, item{100,3,0})
	Rings = append(Rings, item{20,0,1})
	Rings = append(Rings, item{40,0,2})
	Rings = append(Rings, item{80,0,3})

	maxCost := math.MinInt32
	userHitPoints := 100
	enemyHitPoints := 109
	enemyDamage := 8
	enemyArmor := 2
	for i:=0;i<len(Weapons);i++{
		var cost int
		var damage int
		var armor int
		var userShotsNeeded int
		var enemyShotsNeeded int
		cost = Weapons[i].cost
		damage = Weapons[i].damage
		armor = Weapons[i].armor
		userShotsNeeded = getShotsNeeded(enemyHitPoints,damage,enemyArmor)
		enemyShotsNeeded = getShotsNeeded(userHitPoints,enemyDamage,armor)
		//println(userShotsNeeded)
		//println(enemyShotsNeeded)
		if userShotsNeeded >= enemyShotsNeeded {
			if cost > maxCost {
				println("--------")
				println(i)
				maxCost = cost
			}
			// If we can defeat enemy using just a weapon
			// no need to try other combinations with this weapon
			// as this would only increase the cost
			// continue
		}

		for j:=0;j<len(Armor);j++ {
			costArmor := cost + Armor[j].cost
			damageArmor := damage + Armor[j].damage
			armorArmor  := armor + Armor[j].armor
			userShotsNeeded = getShotsNeeded(enemyHitPoints,damageArmor,enemyArmor)
			enemyShotsNeeded = getShotsNeeded(userHitPoints,enemyDamage,armorArmor)

			if userShotsNeeded >= enemyShotsNeeded {
				if costArmor > maxCost {
					println("-----")
					println(i)
					println(j)
					maxCost = costArmor
				}
				// If we can defeat an enemy using this combination of
				// weapon and armor, no need to try other combinations
				// as the cost would only be higher
				// continue
			}
			for k:=0;k<len(Rings)-1;k++{
				costRing1 := costArmor + Rings[k].cost
				damageRing1 := damageArmor + Rings[k].damage
				armorRing1  := armorArmor + Rings[k].armor
				userShotsNeeded = getShotsNeeded(enemyHitPoints, damageRing1, enemyArmor)
				enemyShotsNeeded = getShotsNeeded(userHitPoints,enemyDamage,armorRing1)
				if userShotsNeeded >= enemyShotsNeeded {
					if costRing1 > maxCost {
						println("-----")
						println(i)
						println(j)
						println(k)
						maxCost = costRing1
					}
					// If we can defeat an enemy using just this ring
					// No need to consider a second ring, as the cost would
					// be higher
					// continue
				}
				for l:=k+1;l<len(Rings);l++{
					costRing2  := costRing1 +  Rings[l].cost
					damageRing2 := damageRing1 + Rings[l].damage
					armorRing2  := armorRing1 + Rings[l].armor
					userShotsNeeded = getShotsNeeded(enemyHitPoints, damageRing2, enemyArmor)
					enemyShotsNeeded = getShotsNeeded(userHitPoints, enemyDamage, armorRing2 )
					if userShotsNeeded >= enemyShotsNeeded {
						if costRing2 > maxCost {
							println("-----")
							println(i)
							println(j)
							println(k)
							println(l)
							maxCost = costRing2
						}
					}
				}
			}
		}
	}
	println(maxCost)
}