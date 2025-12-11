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
	queue := make(utils.Queue[*Device], 0)
	queue.Push(server.GetDevice("you"))

	paths := 0
	for !queue.IsEmpty() {
		cur := queue.Pop()

		for device := range cur.connections {
			if device.id == "out" {
				paths++
				continue
			} else {
				queue.Push(device)
			}
		}
	}

	return paths
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
