package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Mathcher)

func Run(searchTerm string) {
	// 검색할 피드의 목록을 조회한다.
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 버퍼가 없는 채널을 생성하여 화면에 표시할 검색 결과를 전달받는다.
	results := make(chan *Result)

	// 모든 피드를 처리할 때까지 기다릴 대기 그룹(Wait group)을 설정한다.
	var waitGroup sync.WaitGroup

	// 개별 피드를 처리하는 동안 대기해야 할 고루틴의 개수를 설정한다.
	waitGroup.Add(len(feeds))

	// 각기 다른 종류의 피드를 처리할 고루틴을 실행한다.
	for _, feed := range feeds {
		// 검색을 위해 검색기를 조회한다.
		matcher, exits := matchers[feed.Type]
		if !exits {
			matcher = matchers["default"]
		}

		// 검색을 실행하기 위해 고루틴을 실행한다.
		go func(matcher Matcher, feed *Feed) {
			Matcher(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 모든 작업이 완료되었지를 모니터링할 고루틴을 실행한다.
	go func() {
		// 모든 작업이 처리될 때까지 기다린다.
		waitGroup.Wait()

		// Display 함수에게 프로그램을 종료할 수 있을을 알리기 위해 채널을 닫는다.
		close(results)
	}()

	Display(results)
}
