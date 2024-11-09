package managers

import (
	"github.com/gopxl/pixel/v2"
	"github.com/wirdos/directors/input"
)

type Player struct {
	puppet PlayableActor
	input input.Input

	requestsExit bool
}

func (p *Player) Update() {
	// TODO: currently only one input type, but the idea is that the player
	// knows how to decipher different input mechanisms to direct puppet
	p.puppet.FeedDirection(p.input.Direction())
	p.puppet.FeedInteract(p.input.Interact())

	p.requestsExit = p.input.Exit()
}

func (p *Player) RequestsExit() bool {
	return p.requestsExit
}

func (p *Player) SetPuppet(puppet PlayableActor) {
	p.puppet = puppet
}

func (p *Player) SetInput(input input.Input) {
	p.input = input
}

func NewPlayer() *Player {
	return &Player{
		requestsExit: false,
	}
}

type PlayableActor interface {
	FeedDirection(d pixel.Vec)
	FeedInteract(interact bool)
}
