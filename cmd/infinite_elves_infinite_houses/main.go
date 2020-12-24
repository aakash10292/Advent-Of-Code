package main

func main(){
	number := 33100000
	houses := make([]int,number)
	for elf:=1;elf<=number/10;elf++{
		maxHouse := number / elf / 10
		if maxHouse > 50 {
			maxHouse = 50
		}
		for house:=1;house<=maxHouse;house++{
			houses[elf*house] += elf*11
		}
	}

	for i:=0;i<len(houses);i++{
		if houses[i] >= number {
			println(i)
			break
		}
	}
}
