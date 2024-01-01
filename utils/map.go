package utils

func Map[S any, D any](src []S, callback func(S) D) []D {
	dist := []D{}
	for _, s := range src {
		dist = append(dist, callback(s))
	}
	return dist
}
