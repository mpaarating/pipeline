package main

type newMatcheResponse struct {
	Id        string   `json:"id"`
	StartedAt int64    `json:"started_at"`
	GridSize  int      `json:"gridsize"`
	Players   []player `json:"plaayers"`
}

type player struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}
