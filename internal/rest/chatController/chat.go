package chatController

//
//import (
//	"chat-api/internal/model/dto/request/chatRequest"
//	"chat-api/internal/service"
//	_default "chat-api/internal/utilitis/default"
//	"github.com/labstack/echo/v4"
//	"net/http"
//)
//
//type ChatController interface {
//	NewRoamController(e echo.Context) error
//	GetBatchChatController(e echo.Context) error
//}
//type chatController struct {
//	chatService service.ChatService
//}
//
//func NewChatController(chatService service.ChatService) ChatController {
//	return &chatController{chatService: chatService}
//}
//func (c *chatController) NewRoamController(e echo.Context) error {
//	var req = new(chatRequest.NewRoamRequest)
//	err := e.Bind(&req)
//	if err != nil {
//		return e.JSON(http.StatusUnprocessableEntity, _default.InternalError(err))
//	}
//	res := c.chatService.NewRoamService(e.Request().Context(), req)
//	if res.Status != "success" {
//		return e.JSON(res.Code, res)
//	}
//	return e.JSON(http.StatusCreated, res)
//}
//func (c *chatController) GetBatchChatController(e echo.Context) error {
//	res := c.chatService.GetBatchChatService(e.Request().Context())
//	if res.Status != "success" {
//		return e.JSON(res.Code, res)
//	}
//	return e.JSON(http.StatusOK, res)
//}
