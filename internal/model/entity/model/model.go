package model

var (
	TypeText TypeMessage = "TEXT"
	TypeFile TypeMessage = "FILE"

	PersonalChat TypeChat = "PERSONAL"
	GroupChat    TypeChat = "GROUP"

	StatusRead   ReadStatus = "READ"
	StatusUnread ReadStatus = "UNREAD"

	NewRoamResp ServiceType = "NEW_ROAM"
	NewMessage  ServiceType = "NEW_MESSAGE"
	AddContact  ServiceType = "ADD_CONTACT"
	GetChatById ServiceType = "GET_CHAT_BY_ID"
)

type TypeMessage string

type TypeChat string

type ReadStatus string

type ServiceType string
