// Example is provided with help by Gabriel Aszalos.
// Package runner manages the running and lifetime of a process.
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner runs a set of tasks within a given timeout and can be
// shut down on an operating system interrupt.
type Runner struct {
	// interrupt channel reports a signal from the
	// operating system.
	interrupt chan os.Signal

	// complete channel reports that processing is done.
	complete chan error

	// timeout reports that time has run out.
	timeout <-chan time.Time

	// tasks holds a set of functions that are executed
	// synchronously in index order.
	tasks []func(int)
}

// ErrTimeout is returned when a value is received on the timeout channel.
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt is returned when an event from the OS is received.
var ErrInterrupt = errors.New("received interrupt")

// New returns a new ready-to-use Runner.
func New(d time.Duration) *Runner {
	return &Runner{
		// 通道 interrupt 被初始化为缓冲区容量为 1 的通道。这可以保证通道至少能接收一个来自
		//语言运行时的 os.Signal 值，确保语言运行时发送这个事件的时候不会被阻塞。如果 goroutine
		//没有准备好接收这个值，这个值就会被丢弃
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add attaches tasks to the Runner. A task is a function that
// takes an int ID.
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all tasks and monitors channel events.
func (r *Runner) Start() error {
	// We want to receive all interrupt based signals.
	signal.Notify(r.interrupt, os.Interrupt)

	// Run the different tasks on a different goroutine.
	go func() {
		r.complete <- r.run()
	}()

	select {
	// Signaled when processing is done.
	case err := <-r.complete:
		return err

	// Signaled when we run out of time.
	case <-r.timeout:
		return ErrTimeout
	}
}

// run executes each registered task.
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// Check for an interrupt signal from the OS.
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// Execute the registered task.
		task(id + 1)
	}

	return nil
}

// gotInterrupt verifies if the interrupt signal has been issued.
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件被触发时发出的信号
	// Signaled when an interrupt event is sent.
	case <-r.interrupt:
		// Stop receiving any further signals.
		signal.Stop(r.interrupt)
		return true

	// Continue running as normal.
	default:
		return false
	}
}
