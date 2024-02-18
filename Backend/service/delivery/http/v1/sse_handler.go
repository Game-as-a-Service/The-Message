package http

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

type Event struct {
	Message       chan map[string]any
	NewClients    chan chan string
	ClosedClients chan chan string
	TotalClients  map[chan string]any
}

type ClientChan chan string

func NewSSEServer() (event *Event) {
	event = &Event{
		Message:       make(chan map[string]any),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]any),
	}

	go event.listen()

	return
}

func (stream *Event) listen() {
	for {
		select {
		case client := <-stream.NewClients:
			stream.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

		case client := <-stream.ClosedClients:
			delete(stream.TotalClients, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(stream.TotalClients))

		case eventMsg := <-stream.Message:
			jsonMsg, err := json.Marshal(eventMsg)
			if err != nil {
				continue
			}

			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- string(jsonMsg)
			}
		}
	}
}

func (stream *Event) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientChan := make(ClientChan)

		stream.NewClients <- clientChan

		defer func() {
			stream.ClosedClients <- clientChan
		}()

		c.Set("clientChan", clientChan)

		c.Next()
	}
}

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}
