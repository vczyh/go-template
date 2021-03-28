package demo

func Test(query string) (map[string]string, error) {
	l.Debug("test request")

	return map[string]string{
		"q": query,
	}, nil
}
