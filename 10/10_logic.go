package day10

import (
	"fmt"
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// LIGHT STATE
// ========================
type LightState struct {
	lights []int
	inputs int
}

func (s LightState) GenerateKey() int {
	key := 0
	for _, b := range s.lights {
		key = (key << 1) | b
	}
	return key
}
func (s LightState) UpdateState(seqeunce []int) *LightState {
	nextConfiguration := make([]int, len(s.lights))
	copy(nextConfiguration, s.lights)

	for _, index := range seqeunce {
		nextConfiguration[index] ^= 1
	}

	return &LightState{
		lights: nextConfiguration,
		inputs: s.inputs + 1,
	}
}
func (s LightState) IsMatch(target []int) bool {
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

func (m Machine) ConfigureLights() int {
	queue := make(utils.Queue[*LightState], 0)
	queue.Push(&LightState{lights: make([]int, len(m.lights)), inputs: 0})

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
// VOLTAGE STATE
// ========================
type VoltageState struct {
	target    []int
	values    []int
	gState    int
	hState    int
	fState    int
	maxButton int
}

func NewVoltageState(target []int, buttons [][]int) *VoltageState {
	remaining := GetRemaining(target)
	maxButton := 0

	for _, button := range buttons {
		maxButton = max(maxButton, len(button))
	}

	gState := 0
	hState := (remaining + maxButton - 1) / maxButton
	fState := gState + hState

	return &VoltageState{
		target:    append([]int(nil), target...),
		values:    make([]int, len(target)),
		gState:    gState,
		hState:    hState,
		fState:    fState,
		maxButton: maxButton,
	}
}
func GetRemaining(values []int) int {
	rem := 0
	for _, value := range values {
		rem += value
	}
	return rem
}
func StateSorter(a, b *VoltageState) bool {
	if a.fState == b.fState {
		return a.gState < b.gState
	}
	return a.fState < b.fState
}

func (v VoltageState) GenerateKey() string {
	return fmt.Sprint(v.values)
}
func (v VoltageState) IsComplete() bool {
	for _, value := range v.target {
		if value != 0 {
			return false
		}
	}
	return true
}

func (v VoltageState) NewState(sequence []int) (*VoltageState, bool) {
	nextTarget := make([]int, len(v.target))
	copy(nextTarget, v.target)
	nextValues := make([]int, len(v.values))
	copy(nextValues, v.values)

	for _, index := range sequence {
		nextTarget[index]--

		if nextTarget[index] < 0 {
			return &VoltageState{}, false
		}
		nextValues[index]++
	}

	nextH := (GetRemaining(nextTarget) + v.maxButton - 1) / v.maxButton
	nextG := v.gState + 1
	nextF := nextH + nextG

	return &VoltageState{
		target:    nextTarget,
		values:    nextValues,
		gState:    nextG,
		hState:    nextH,
		fState:    nextF,
		maxButton: v.maxButton,
	}, true
}
func (m Machine) ConfigureVoltage() int {
	queue := utils.NewPriorityQueue(StateSorter)
	queue.Push(NewVoltageState(m.joltage, m.buttons))

	visited := make(map[string]int)

	for !queue.IsEmpty() {
		cur := queue.Remove()
		key := cur.GenerateKey()

		if val, found := visited[key]; found && val >= cur.gState {
			continue
		} else {
			visited[key] = cur.gState
		}

		if cur.IsComplete() {
			return cur.gState
		}

		for _, seqeunce := range m.buttons {
			if nextState, isValid := cur.NewState(seqeunce); isValid {
				queue.Push(nextState)
			}
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
		sum += machine.ConfigureLights()
	}
	return sum
}

// ========================
// PART II
// ========================
func II(machines []*Machine) int {
	sum := 0
	for _, machine := range machines {
		sum += machine.ConfigureVoltage()
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
