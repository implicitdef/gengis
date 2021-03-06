package myerrors


type NotInCacheError struct {
	Msg string
}
func (e *NotInCacheError) Error() string {
	return e.Msg
}

type OtherTheMovieDbError struct {
	Msg string
}
func (e *OtherTheMovieDbError) Error() string {
	return e.Msg
}

type TooManyRequestsError struct {
	Msg string
}
func (e *TooManyRequestsError) Error() string {
	return e.Msg
}


type UnprocessableSeasonError struct {
	Msg string
}
func (e *UnprocessableSeasonError) Error() string {
	return e.Msg
}

