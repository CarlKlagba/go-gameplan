package domain

import (
	"log"
	"testing"
)

func TestGameplanName(t *testing.T) {
	var gameplan = Gameplan{
		Name: "passing to north south",
		Action: &Action{
			Name:      "get angle",
			Reactions: []*Reaction{},
		},
	}
	if gameplan.Name != "passing to north south" {
		t.Errorf("GameplaneName = %s; want works", gameplan.Name)
	}
	if gameplan.Action.Name != "get angle" {
		t.Errorf("GameplaneName = %s; want get angl", gameplan.Action.Name)
	}
	//if len(gameplan.firstAction.Reactions) != 0 {
	//	t.Errorf("GameplaneName = %s; to be empty", gameplan.firstAction.Reactions)
	//}
}

func TestCreateGameplanFirstAction(t *testing.T) {
	var gameplan = Gameplan{
		Name: "passing to north south",
	}

	var action = gameplan.CreateFirstAction("get angle")
	action.AddReaction("use bottom leg")

	log.Println("gameplan", gameplan)

	if gameplan.Action.Name != "get angle" {
		t.Errorf("GameplaneName = %s; want 'get angle'", gameplan.Action.Name)
	}
	if len(gameplan.Action.Reactions) == 1 && gameplan.Action.Reactions[0].Name != "use bottom leg" {
		t.Error(gameplan.Action.Reactions)
	}
}

func TestCreateGameplanMultipleReactions(t *testing.T) {
	var gameplan = Gameplan{
		Name: "passing to north south",
	}

	var action = gameplan.CreateFirstAction("get angle")
	action.AddReaction("use bottom leg")
	action.AddReaction("use top leg")

	log.Println("gameplan", gameplan)

	if gameplan.Action.Name != "get angle" {
		t.Errorf("GameplaneName = %s; want 'get angle'", gameplan.Action.Name)
	}
	if len(gameplan.Action.Reactions) == 2 &&
		gameplan.Action.Reactions[0].Name != "use bottom leg" &&
		gameplan.Action.Reactions[1].Name != "use top leg" {
		t.Error(gameplan.Action.Reactions)
	}
}

func TestCreateGameplanChainReactions(t *testing.T) {
	var gameplan = Gameplan{
		Name: "passing to north south",
	}

	var action = gameplan.CreateFirstAction("get angle")
	reaction := action.AddReaction("use bottom leg")
	action2 := reaction.AddAction("switch lead arm to top hip and go back to the legs")
	action2.AddReaction("heist up")
	log.Println("gameplan", gameplan)
	log.Println("action2", action2)

	if gameplan.Action.Reactions[0].Action.Name != "switch lead arm to top hip and go back to the legs" &&
		gameplan.Action.Reactions[0].Action.Id != 1 {
		t.Error("error creating chain reaction")
	}
	if gameplan.Action.Reactions[0].Action.Reactions[0].Name != "heist up" &&
		gameplan.Action.Reactions[0].Action.Reactions[0].Id != 2 {
		t.Error("error creating chain reaction")
	}
}
