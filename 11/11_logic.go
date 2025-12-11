package day11

import (
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// DEVICE
// ========================
type Device struct {
	id          string
	connections utils.Set[*Device]
}

func NewDevice(id string) *Device {
	return &Device{
		id:          id,
		connections: make(utils.Set[*Device]),
	}
}
func (d *Device) AddConnection(device *Device) {
	d.connections.Add(device)
}

// ========================
// SERVER
// ========================
type Server map[string]*Device

// Add Device, return new device
func (s *Server) GetDevice(id string) *Device {
	if device, found := (*s)[id]; found {
		return device
	}

	(*s)[id] = NewDevice(id)
	return (*s)[id]
}

// ========================
// PART I
// ========================
func I(server Server) int {
	return CountPaths(server.GetDevice("you"), "out", make(map[string]int))
}

// ========================
// PART II
// ========================

func II(server Server) int {
	routeA := [][2]string{{"svr", "fft"}, {"fft", "dac"}, {"dac", "out"}}
	routeB := [][2]string{{"svr", "dac"}, {"dac", "fft"}, {"fft", "out"}}

	productA := 1
	for _, pair := range routeA {
		productA *= CountPaths(server.GetDevice(pair[0]), pair[1], make(map[string]int))
	}
	productB := 1
	for _, pair := range routeB {
		productB *= CountPaths(server.GetDevice(pair[0]), pair[1], make(map[string]int))
	}

	return productA + productB
}

func CountPaths(node *Device, target string, cache map[string]int) int {
	if node.id == target {
		return 1
	}

	if val, ok := cache[node.id]; ok {
		return val
	}

	total := 0
	for neighbor := range node.connections {
		total += CountPaths(neighbor, target, cache)
	}

	cache[node.id] = total
	return total
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Server {
	data := utils.ReadFile(file)
	server := make(Server)

	for line := range strings.SplitSeq(data, "\n") {
		ids := utils.QuickMatch(line, `[a-z]+`)

		source := server.GetDevice(ids[0])
		for _, id := range ids[1:] {
			source.AddConnection(server.GetDevice(id))
		}
	}

	return server
}
