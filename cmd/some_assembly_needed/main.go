package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	instruction string
	operands map[string]int
	inDegree int
	value int
	name string
}

var inDegree map[string]int
var edges map[string][]string
var nodes map[string]*node
var bfsQueue []*node

func isNum(operand string) (int,bool) {
	if num, err := strconv.Atoi(operand); err==nil{
		return num,true
	}
	return -1,false
}

func processLine(input string) {
	ops := strings.Split(input, " ")
	var name string
	var curNode node
	if len(ops) == 3 {
		//assignment operator
		name = ops[2]
		if num, err := strconv.Atoi(ops[0]); err == nil {
			// assigning numeric value to wire
			curNode = node{"", nil, 0, num, name}
			bfsQueue = append(bfsQueue, &curNode)
		} else {
			// assigning wire value to another wire
			curNode = node{ input, nil,1,-1, name}
			edges[ops[0]] = append(edges[ops[0]], ops[2])
		}
	} else if len(ops) == 4 {
		// unary NOT operator
		name = ops[3]
		curNode = node{input, nil, 1, -1,name }
		edges[ops[1]] = append(edges[ops[1]], ops[3])
	} else {
		// AND, OR, LSHIFT, RSHIFT
		name = ops[4]
		var inDeg int
		if _,err :=strconv.Atoi(ops[0]); err!=nil{
			// first operand is not a number
			inDeg++
			edges[ops[0]] = append(edges[ops[0]], ops[4])
		}
		if _,err := strconv.Atoi(ops[2]); err!=nil{
			inDeg++
			edges[ops[2]] = append(edges[ops[2]], ops[4])
		}
		if inDeg == 0{
			log.Fatal("Got indegree zero while processing a two operand instruction")
			os.Exit(-1)
		}
		curNode = node{ input, nil,inDeg,-1,name}
	}
	nodes[name] = &curNode
}

func (curNode *node) executeInstruction() {
	ops := strings.Split(curNode.instruction, " " )
	if len(ops) == 3{
		// wire assignment
		curNode.value = curNode.operands[ops[0]]
	} else if len(ops) == 4 {
		// Unary NOT
		curNode.value = ^curNode.operands[ops[1]]
		if curNode.value < 0 {
			curNode.value = 65536 + curNode.value
		}
	} else {
		var left, right int
		if num,err := strconv.Atoi(ops[0]); err == nil {
			left = num
		} else {
			left = curNode.operands[ops[0]]
		}

		if num,err := strconv.Atoi(ops[2]); err == nil {
			right = num
		}else {
			right = curNode.operands[ops[2]]
		}
		if ops[1] == "AND" {
			curNode.value = left & right
		} else if ops[1] == "OR" {
			curNode.value = left | right
		} else if ops[1] == "LSHIFT" {
			curNode.value = left << uint(right)
		} else {
			curNode.value = left >> uint(right)
		}
	}
}

func main(){
	f,_ := os.Open("/home/aarayambeth/go-stuff/advent_of_code/some_assembly_needed/input.txt")
	/*defer func(){
		if err := f.Close(); err != nil {
			log.Fatal("Could not close file")
			os.Exit(-1)
		}
	}()*/
	// Initialize data structures
	inDegree = make(map[string]int)
	edges = make(map[string][]string)
	nodes = make(map[string]*node)

	s := bufio.NewScanner(f)
	for s.Scan(){
		processLine(s.Text())
	}

	for len(bfsQueue) != 0{
		var curNode *node
		curNode = bfsQueue[0]
		if  curNode.instruction != "" {
			// we need to execute the instruction and get the result
			curNode.executeInstruction()
		}
		for _,node := range edges[curNode.name]{
			// put curNode value in each dependent node
			// reduce inDegree by 1
			// If inDegree of dependent node is 0, add to bfsQueue
			dependentNode := nodes[node]
			if dependentNode.operands == nil{
				dependentNode.operands = make(map[string]int)
			}
			dependentNode.operands[curNode.name] = curNode.value
			dependentNode.inDegree--
			if dependentNode.inDegree == 0{
				bfsQueue = append(bfsQueue, dependentNode)
			}
		}
		// remove first node from bfsQueue
		bfsQueue = bfsQueue[1:]
	}
	println(nodes["a"].value)
}
