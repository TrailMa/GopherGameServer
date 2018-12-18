package helpers

import (

)

type GopherError struct {
	Message string
	ID      int
}

// Client response message error IDs
const (
	// Package gopher
	Error_Gopher_Invalid_Action = iota + 1001
	Error_Gopher_Incorrect_Format
	Error_Gopher_Incorrect_Custom_Action
	Error_Gopher_Not_Logged_In
	Error_Gopher_Logged_In
	Error_Gopher_Status_Change
	Error_Gopher_Feature_Disabled
	Error_Gopher_Columns_Format
	Error_Gopher_Name_Format
	Error_Gopher_Password_Format
	Error_Gopher_Remember_Format
	Error_Gopher_Guest_Format
	Error_Gopher_New_Password_Format
	Error_Gopher_Room_Name_Format
	Error_Gopher_Room_Type_Format
	Error_Gopher_Private_Format
	Error_Gopher_Max_Room_Format
	Error_Gopher_Room_Control
	Error_Gopher_Server_Room
	Error_Gopher_Not_Owner
	Error_Gopher_Login
	Error_Gopher_Sign_Up
	Error_Gopher_Delete_Account_Error
	Error_Gopher_Password_Change
	Error_Gopher_Info_Change
	Error_Gopher_Join
	Error_Gopher_Leave
	Error_Gopher_Create_Room
	Error_Gopher_Delete_Room
	Error_Gopher_Invite
	Error_Gopher_Revoke_Invite
	Error_Gopher_Friend_Request
	Error_Gopher_Friend_Accept
	Error_Gopher_Friend_Decline
	Error_Gopher_Friend_Remove
)

func NewError(message string, id int) GopherError {
	return GopherError{Message: message, ID: id}
}