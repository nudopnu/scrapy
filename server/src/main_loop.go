package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type LoopState struct {
	taskQueue   chan Task
	cancelQueue context.CancelFunc
	ctx         context.Context
}

type Task struct {
	name      string
	operation func() error
}

func mainLoop(state *State, ctx context.Context, taskQueue chan Task) {
	log.Println("Start main loop")
	for {
		select {
		case <-ctx.Done():
			log.Println("Exit main loop")
			return
		case task := <-taskQueue:
			log.Printf("[task]: %s\n", task.name)
			err := processTask(task, 3)
			if err != nil {
				log.Println(err)
			}
		case <-time.After(3 * time.Second):
			agent, err := state.db.GetNextAgentToUpdate(ctx)
			if err == nil {
				log.Println("Adding update task")
				go func() {
					state.ls.taskQueue <- *state.NewUpdateAgentTask(int(agent.ID))
					log.Printf("Added update task for agent %s\n", agent.Name)
				}()
			} else {
				log.Println("Idle...")
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func processTask(task Task, retries int) error {
	for numRetries := retries; numRetries > 0; numRetries-- {
		err := task.operation()
		if err == nil {
			return nil
		}
		log.Println(err)
		log.Printf("operation failed, retrying...\n")
	}
	return fmt.Errorf("operation '%s' failed 3 times", task.name)
}

func (state *State) StartMainLoop() {
	go mainLoop(state, state.ls.ctx, state.ls.taskQueue)
}

func (state *State) AddTask(task Task) {
	state.ls.taskQueue <- task
}
