package themoviedb


type tvshow struct {
	Seasons []season
}
type season struct {
	SeasonNumber int `json:"season_number"`
}