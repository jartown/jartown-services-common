package common

type AckReply struct {
	Success bool `json:"success"`
}

type ErrorReply struct {
	ErrorMsg string `json:"error"`
}
