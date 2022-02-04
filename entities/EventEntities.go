package entities

type Event struct {
	Id          int    `json:"id"`
	Id_user     int    `json:"id_user"`
	Id_category int    `json:"id_category"`
	Title       string `json:"title"`
	Start_date  string `json:"star_date"`
	End_date    string `json:"end_date"`
	Location    string `json:"location"`
	Details     string `json:"details"`
	Photo       string `json:"photo"`
}
