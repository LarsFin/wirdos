package managers

import (
	"github.com/wirdos/directors/input"
)

type Player struct {
	puppet Playable
	controller input.InputController

	requestsExit bool
}

func (p *Player) Update() {
	input := p.controller.Input()

	// TODO: currently only one input type, but the idea is that the player
	// knows how to decipher different input mechanisms to direct puppet
	p.puppet.FeedInput(input)
	p.requestsExit = input.Exit
}

func (p *Player) RequestsExit() bool {
	return p.requestsExit
}

func (p *Player) SetPuppet(puppet Playable) {
	p.puppet = puppet
}

func (p *Player) SetInput(input input.InputController) {
	p.controller = input
}

func NewPlayer() *Player {
	return &Player{
		requestsExit: false,
	}
}

type Playable interface {
	FeedInput(input *input.Input)
}
