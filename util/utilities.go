package util

//CreateIfNotExistsAndAddIfNotContained creates a slice if needed and adds if the value is not contained
func CreateIfNotExistsAndAddIfNotContained(a []string, b string) []string {
	if a == nil {
		a = make([]string, 0)
	}
	return AddIfNotContained(a, b)
}

//AddIfNotContained adds a string to the slice if its not contained and returns the slice
func AddIfNotContained(a []string, b string) []string {
	contains := false
	for _, v := range a {
		if v == b {
			contains = true
			break
		}
	}
	if !contains {
		a = append(a, b)
	}
	return a
}
