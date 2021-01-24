package util

func StringSliceToSet(slice []string) []string {
	m := map[string]bool{}
	for _, s := range slice {
		m[s] = true
	}

	list := []string{}
	for item := range m {
		list = append(list, item)
	}
	return list
}
