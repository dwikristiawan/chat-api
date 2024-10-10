package websocket

//
//import (
//	"chat-api/internal/model/dto/response/chatResponse"
//	"chat-api/internal/model/entity"
//	"encoding/json"
//	"fmt"
//	"github.com/gorilla/websocket"
//	"github.com/labstack/echo/v4"
//	"github.com/labstack/gommon/log"
//	"time"
//)
//
//type WebSocketController struct {
//	connection     map[string]*websocket.Conn
//	BroadcastQueue map[string]bool
//	QueMessages    map[string]chatResponse.BroadcastResponse
//
//	//broadcastService service.BroadcastService
//
//}
//
//type Chat struct {
//	ID        uint
//	UpdatedAt time.Time
//	Name      string
//	Messages  *entity.Messages
//}
//
//func NewWebSocketController() WebSocketController {
//	return WebSocketController{
//		connection:     make(map[string]*websocket.Conn),
//		BroadcastQueue: make(map[string]bool),
//		QueMessages:    make(map[string]chatResponse.BroadcastMessage),
//	}
//}
//
//func HandlerWsInit(e *echo.Echo, controller WebSocketController) {
//	r := e.Group("api/v1")
//	r.GET("/ws/:id", controller.AddConn)
//}
//
//var upgrader = websocket.Upgrader{}
//
//func (ctr *WebSocketController) AddConn(e echo.Context) error {
//	ws, err := upgrader.Upgrade(e.Response(), e.Request(), nil)
//	if err != nil {
//		return err
//	}
//	defer ws.Close()
//	ctr.connection[e.Param("id")] = ws
//	for {
//		// Write
//		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
//		if err != nil {
//			if err.Error() == "websocket: close sent" {
//				//
//			} else {
//				e.Logger().Error(err)
//			}
//		}
//
//		// Read
//		_, msg, err := ws.ReadMessage()
//		if err != nil {
//			e.Logger().Error(err)
//		}
//		fmt.Printf("%+v\n", ws.RemoteAddr())
//		fmt.Printf("%s\n", msg)
//	}
//
//}
//
//func (svc *WebSocketController) WSBroadcastMassage() {
//
//	for {
//		if len(svc.QueMessages) == 0 {
//			time.Sleep(5 * time.Second)
//		} else {
//			for key, queData := range svc.QueMessages {
//				if ws, exist := svc.connection[key]; exist {
//					queData.Mutex.Lock()
//					go func(wsConn *websocket.Conn, data *chatResponse.BroadcastData, keyMap string) {
//						svc.broadcastMassage(wsConn, data)
//						defer func() {
//							queData.Mutex.Unlock()
//							delete(svc.QueMessages, keyMap)
//						}()
//					}(ws, &queData.Data, key)
//				}
//				delete(svc.QueMessages, key)
//				continue
//			}
//		}
//	}
//}
//func (svc *WebSocketController) broadcastMassage(ws *websocket.Conn, msg *chatResponse.BroadcastData) {
//	byteMsg, _ := json.Marshal(*msg)
//	err := ws.WriteMessage(websocket.TextMessage, byteMsg)
//	if err != nil {
//		log.Errorf("Err broadcastMassage.ws.Write err: %v", err)
//		return
//	}
//}
