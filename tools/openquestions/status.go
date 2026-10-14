package main

func countStatus(m manifest, want string) int {
	count := 0
	for _, q := range m.Questions {
		if q.Status == want {
			count++
		}
	}
	return count
}

func status(m manifest) string {
	if countStatus(m, "open") > 0 {
		return "advisory_open_questions"
	}
	return "verified"
}
