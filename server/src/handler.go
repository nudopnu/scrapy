package main

// import (
// 	"log"
// )

type Command struct {
	name string
	args []string
}

//	func handleCommands(state *State, cmd Command) error {
//		handlers := map[string]func(*State, Command) error{
//			"locations": HandlerListLocations,
//			"location":  HandlerGetLocationId,
//			"users":     HandlerUsers,
//			"login":     HandlerLogin,
//			"register":  HandlerRegister,
//			"agents":    HandlerListSearchAgents,
//			"agent":     middlewareLoggedIn(HandlerAddSearchAgent),
//			"ads":       HandlerAds,
//			"run":       middlewareLoggedIn(HandlerRun),
//			"reset":     HandlerReset,
//		}
//		handler, ok := handlers[cmd.name]
//		if !ok {
//			log.Fatalf("command '%s' not found", cmd.name)
//		}
//		return handler(state, cmd)
//	}
