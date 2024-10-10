package service

import (
	"chat-api/internal/repository/postgres"
	ws_resource "chat-api/internal/ws-resource"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/gommon/log"
	"sync"
)

type broadcastService struct {
	chatRepository        postgres.ChatRepository
	msgRepository         postgres.MessageRepository
	userRepository        postgres.UserRepository
	participantRepository postgres.ParticipantRepository
	referenceRepository   postgres.ReferenceRepository
	wsConn                *ws_resource.WsConn
}

func NewBroadcastService(
	chatRepository postgres.ChatRepository,
	msgRepository postgres.MessageRepository,
	userRepository postgres.UserRepository,
	participantRepository postgres.ParticipantRepository,
	referenceRepository postgres.ReferenceRepository,
) BroadcastService {
	return &broadcastService{
		chatRepository:        chatRepository,
		msgRepository:         msgRepository,
		userRepository:        userRepository,
		participantRepository: participantRepository,
		referenceRepository:   referenceRepository,
	}
}

type BroadcastService interface {
	BroadcastChatService(context.Context, *map[string]interface{})
}

func (s *broadcastService) BroadcastChatService(ctx context.Context, queue *map[string]interface{}) {
	var wg sync.WaitGroup
	var iteration int
	for key, value := range *queue {
		if ws, exist := s.wsConn.Connection[key]; exist {
			iteration++
			wg.Add(iteration)
			go func(conn *websocket.Conn, value interface{}, wg *sync.WaitGroup) {
				defer wg.Done()
				jsonValue, err := json.Marshal(value)
				if err != nil {
					log.Errorf("Err broadcast.json.marshal err: %v : %v", value, err)
				}
				err = conn.WriteMessage(websocket.TextMessage, jsonValue)
				if err != nil {
					log.Errorf("Err broadcast.ws.WriteMessage err: %v : %v", conn.RemoteAddr(), err)
				}
			}(ws, value, &wg)
		}
	}
	wg.Wait()
}
