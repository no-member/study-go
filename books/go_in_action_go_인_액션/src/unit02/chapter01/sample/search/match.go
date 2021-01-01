package search

import (
	"log"
)

// 검색 결과를 저장할 Result 구조체
type Result struct {
	Field   string
	Content string
}

// Matcher 인터페이스는 새로운 검색 타입을 구혈할 떄 필요한 동작을 정의한다.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}
