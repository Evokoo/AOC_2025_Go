package day10

import (
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// STATE
// ========================
type State struct {
	lights []int
	inputs int
}

func (s State) GenerateKey() int {
	key := 0
	for _, b := range s.lights {
		key = (key << 1) | b
	}
	return key
}
func (s State) UpdateState(seqeunce []int) *State {
	nextConfiguration := make([]int, len(s.lights))
	copy(nextConfiguration, s.lights)

	for _, index := range seqeunce {
		nextConfiguration[index] ^= 1
	}

	return &State{
		lights: nextConfiguration,
		inputs: s.inputs + 1,
	}
}
func (s State) IsMatch(target []int) bool {
	for i, value := range s.lights {
		if value != target[i] {
			return false
		}
	}
	return true
}

// ========================
// MACHINE
// ========================

type Machine struct {
	lights  []int
	buttons [][]int
	joltage []int
}

func (m Machine) Configure() int {
	queue := make(utils.Queue[*State], 0)
	queue.Push(&State{lights: make([]int, len(m.lights)), inputs: 0})

	visited := make(map[int]int)

	for !queue.IsEmpty() {
		cur := queue.Pop()
		key := cur.GenerateKey()

		if val, found := visited[key]; found && val >= cur.inputs {
			continue
		} else {
			visited[key] = cur.inputs
		}

		if cur.IsMatch(m.lights) {
			return cur.inputs
		}

		for _, seqeunce := range m.buttons {
			queue.Push(cur.UpdateState(seqeunce))
		}
	}

	panic("Configuration not found")
}

// ========================
// PART I
// ========================
func I(machines []*Machine) int {
	sum := 0
	for _, machine := range machines {
		sum += machine.Configure()
	}

	return sum
}

// ========================
// PARSER
// ========================
func ParseInput(file string) []*Machine {
	data := utils.ReadFile(file)
	lines := strings.Split(data, "\n")

	machines := make([]*Machine, len(lines))

	for i, line := range lines {
		sections := strings.Split(line, " ")
		length := len(sections)

		machines[i] = &Machine{
			lights:  ParseLights(sections[0]),
			buttons: ParseButtons(sections[1 : length-1]),
			joltage: ParseJoltage(sections[length-1]),
		}
	}

	return machines
}
func ParseLights(input string) []int {
	length := len(input)
	lights := make([]int, length-2)
	for i, r := range input[1 : length-1] {
		if r == '#' {
			lights[i] = 1
		}
	}
	return lights
}
func ParseJoltage(input string) []int {
	return utils.MatchInts(input)
}
func ParseButtons(input []string) [][]int {
	length := len(input)
	buttons := make([][]int, length)
	for i, s := range input {
		buttons[i] = utils.MatchInts(s)
	}
	return buttons
}
