package domain

type Pod struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	GitUrl   string   `json:"git_url"`
	Branch   string   `json:"branch"`
	Location string   `json:"location"`
	Commands []string `json:"commands"`
	Started  bool     `json:"started"`
	Pid      int      `json:"pid"`
}
