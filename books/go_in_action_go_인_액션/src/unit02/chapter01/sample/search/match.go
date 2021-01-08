package search

import "log"

// 검색 결과를 저장할 Result 구조체
type Result struct {
	Field   string
	Content string
}

// Matcher 인터페이스는 새로운 검색 타입을 구혈할 떄 필요한 동작을 정의한다.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 함수는 고루틴으로서 호출되며
// 개별 피드 타입에 대한 검색을 동시에 숭행한다.
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 지정된 검색기를 이용해 검색을 수행한다.
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 검색 결과를 채널에 기록한다.
	for _, result := range searchResults {
		results <- result
	}
}
