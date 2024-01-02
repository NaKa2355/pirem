package utils

func Map[S any, D any](src []S, callback func(S) D) []D {
	dist := make([]D, len(src))
	for i := range src {
		dist[i] = callback(src[i])
	}
	return dist
}
