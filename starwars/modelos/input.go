package modelos

type SateliteInput struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type SatelitesInput struct {
	SatelitesInput []SateliteInput `json:"satelites"`
}
