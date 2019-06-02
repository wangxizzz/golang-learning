package search

// defaultMatcher implements the default matcher.
// 定义了一个空结构。空结构在创建实例时，不会分配任何内存
type defaultMatcher struct{}

// init registers the default matcher with the program.
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
