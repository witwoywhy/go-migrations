package domain

type Request struct {
	Action Action
	Schema *Migrate `json:"schema"`
	Data   *Migrate `json:"data"`
}
