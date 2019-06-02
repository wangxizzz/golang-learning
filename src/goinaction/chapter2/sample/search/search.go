package search

import (
	"log"
	"sync"
)

// 注册用于搜索的匹配器的映射。这个变量没有定义在任何函数作用域内，所以会被当成包级变量(因为是小写变量)
var matchers = make(map[string]Matcher)

// Run performs the search logic.
func Run(searchTerm string) {
	//  获取需要搜索的数据源列表
	// Retrieve the list of feeds to search through.
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results to display.
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		// Retrieve a matcher for the search.
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search.
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed) //持有matcher, feed的副本.如果这里使用全局变量，那么就会matcher, feed每次的值都是for的最后一次值
	}

	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for everything to be processed.
		// waitGroup,results是一个对外部变量的引用.是一个闭包
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program.
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	Display(results)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
