package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func ChekcWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}
	resultsChannel := make(chan result)

	for _, url := range(urls) {
		go func(u string) {
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <- resultsChannel
		results[r.string] = r.bool
	}
	
	return results
}