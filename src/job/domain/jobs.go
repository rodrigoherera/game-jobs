package domain

type Jobs struct {
	Tasks []Tasks `json:"tasks"`
}

type Tasks struct {
	Name string `json:"name"`
	Cron string `json:"cron"`
}
