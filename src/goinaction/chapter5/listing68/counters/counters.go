// Package counters provides alert counter support.
package counters

// alertCounter is an unexported type that
// contains an integer counter for alerts.
type alertCounter int

// New creates and returns values of the unexported
// 将工厂函数命名为 New 是 Go 语言的一个习惯
// type alertCounter.
func New(value int) alertCounter {
	// 利用工厂模式来使包外面访问alertCounter
	return alertCounter(value)
}
