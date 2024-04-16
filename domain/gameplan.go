package domain

import "fmt"

type Gameplan struct {
	Name   string  `json:"name"`
	Action *Action `json:"firstAction"`
}

type Action struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	Reactions []*Reaction `json:"reactions"`
}

type Reaction struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Action *Action `json:"action"`
}

func NewGameplan(s string) *Gameplan {
	return &Gameplan{
		Name: s,
	}
}

func (g *Gameplan) CreateFirstAction(s string) *Action {
	g.Action = &Action{
		Id:   actionIdCounter,
		Name: s,
	}
	actionIdCounter++
	return g.Action
}

var reactionIdCounter = 0

func (a *Action) AddReaction(s string) *Reaction {
	reaction := &Reaction{Id: reactionIdCounter, Name: s}
	reactionIdCounter++
	a.Reactions = append(a.Reactions, reaction)
	return reaction
}

var actionIdCounter = 0

func (r *Reaction) AddAction(s string) *Action {
	action := &Action{Id: actionIdCounter, Name: s}
	actionIdCounter++
	r.Action = action
	return action
}

func (g *Gameplan) String() string {
	return fmt.Sprintf("Gameplan{ \n\t %s, \n\t %v }", g.Name, g.Action)
}

func (a *Action) String() string {
	return fmt.Sprintf("\n\tAction{ \n\t %s,   %v }", a.Name, a.Reactions)
}

func (r *Reaction) String() string {
	return fmt.Sprintf("\n\tReaction{ \n\t %s, \n\t %v }", r.Name, r.Action)
}
