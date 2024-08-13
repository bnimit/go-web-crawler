package datafetcher

func fetchHTML(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	return extractText(response.Body)
}
