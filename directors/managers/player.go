package managers

import (
	"github.com/wirdos/directors/input"
)

type Player struct {
	// for now we assume the player can only control one playable entity at a time
	puppet Playable
	controller input.InputController

	requestsExit bool
}

func (p *Player) Update() {
	input := p.controller.Input()
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
