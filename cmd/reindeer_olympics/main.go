package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Reindeer struct {
	flightSpeed int
	flightCapacity int
	restRequired int
	totalPoints int
	totalDistance int
	flightTimeRemaining int
	restTimeRemaining int
}

func (r *Reindeer) processSecond(){
	if r.restTimeRemaining > 0 {
		r.restTimeRemaining--
		if r.restTimeRemaining == 0 {
			r.flightTimeRemaining = r.flightCapacity
		}
	} else {
		r.totalDistance += r.flightSpeed
		r.flightTimeRemaining--
		if r.flightTimeRemaining == 0 {
			r.restTimeRemaining = r.restRequired
		}
	}
}

func main(){
	f, _ := os.Open("/Users/aarayambeth/go_stuff/advent_of_code/cmd/reindeer_olympics/input.txt")
	s := bufio.NewScanner(f)
	var reindeers []*Reindeer
	for s.Scan(){
		ops := strings.Split(s.Text()," ")
		// getting rid of the full stop
		ops[len(ops)-1] = ops[len(ops)-1][:len(ops[len(ops)-1])-1]
		var reindeer Reindeer
		num,_ := strconv.Atoi(ops[3])
		reindeer.flightSpeed = num

		num,_ = strconv.Atoi(ops[6])
		reindeer.flightCapacity = num
		reindeer.flightTimeRemaining = num

		num,_ = strconv.Atoi(ops[13])
		reindeer.restRequired = num

		reindeers = append(reindeers,&reindeer)
	}
	seconds := 2503
	maxDistance := -1
	maxPoints := -1
	var leader *Reindeer
	for seconds > 0 {
		// Tick one second for each reindeer
		for _, reindeer := range reindeers {
			reindeer.processSecond()
			if reindeer.totalDistance > maxDistance {
				leader = reindeer
				maxDistance = reindeer.totalDistance
			}
		}
		leader.totalPoints++
		if leader.totalPoints > maxPoints {
			maxPoints = leader.totalPoints
		}
		seconds--
	}
	println(maxPoints)
}