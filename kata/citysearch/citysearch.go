package citysearch

import "strings"

var cities = []string{"Vienna", "Berlin", "Paris", "Budapest", "London", "Skopje", "Rotterdam", "Valencia", "Vancouver", "Amsterdam", "Sydney", "New York City", "Bangkok", "Hong Kong", "Dubai", "Rome", "Istanbul"}

func Search(input string) []string {
	results := []string{}

	if len(input) < 2 {
		return results
	}

	for _, c := range cities {
		if strings.Contains(strings.ToLower(c), strings.ToLower(input)) {
			results = append(results, c)
		}
	}

	return results
}
