package themoviedb



func GetSeasonNumbers(tvShowId string) ([]int, error) {
	tvs := tvshow{}
	err := doGetAndJsonUnmarshall("/tv/" + tvShowId, &tvs)
	if err != nil {
		return nil, err
	}
	numbers := make([]int, len(tvs.Seasons))
	for i, season := range tvs.Seasons {
		numbers[i] = season.SeasonNumber
	}
	return numbers, nil
}

