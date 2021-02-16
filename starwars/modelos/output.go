package modelos

type OutPut struct {
	Position `json:"position"`
	Message  string `json:"message"`
}

type Position struct {
	X float32
	Y float32
}
