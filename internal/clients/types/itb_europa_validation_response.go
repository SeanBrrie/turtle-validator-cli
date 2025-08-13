package types

type ValidationResponse struct {
	Valid   bool     `json:"valid"`
	Reports []string `json:"reports,omitempty"`
}
