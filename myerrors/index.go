package myerrors


type UnprocessableSeasonError struct {
	Msg string
}
func (e *UnprocessableSeasonError) Error() string {
	return e.Msg
}