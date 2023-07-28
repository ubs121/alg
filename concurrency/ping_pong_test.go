package concurrency

import (
	"fmt"
	"testing"
	"time"
)

type Ball struct{ hits int }

func player(name string, table chan *Ball) {
	for {
		// Receive the ball from the table
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)

		// Send the ball back to the table
		table <- ball
	}
}

func TestPingPong(t *testing.T) {
	// Create a channel for the ping-pong table
	table := make(chan *Ball)

	// Start two goroutines for two players
	go player("ping", table)
	go player("pong", table)

	// Start the game by sending the ball to the table
	table <- new(Ball)

	// Wait for a few moments to let the players play
	time.Sleep(2 * time.Second)

	<-table // game over; grab the ball
}
