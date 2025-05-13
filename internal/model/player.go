package model

import (
	"sync"
)

type Player struct {
	sync.Mutex
	ID          int64       `json:"id"`
	Bank        uint8       `json:"bank"`
	Builds      []Card      `json:"builds"`
	MainBuilds  []MainBuild `json:"main_builds"`
	QueueNotiff chan Nottif `json:"-"`
}

type PlayerCard struct {
	Card
	Player map[int64]PlayerOfCard
}

type MainBuild struct {
	Card
	Builded bool `json:"builded"`
}

func (p *Player) PlayerInit(id int64) {

	p.ID = id
	p.Bank = 0
	p.MainBuilds = InitBuildsCards()
	p.Builds = InitDeckPlayerCards()
}

type PlayerOfCard struct {
	*Player
	CountOfCard uint8
}
