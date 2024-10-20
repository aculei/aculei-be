package utils

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type gracefulExitConfig struct {
	onCloseForError  func(error)
	onCloseForSignal func(os.Signal)
}

func WithGracefulExitOnCloseForError(callback func(error)) func(*gracefulExitConfig) {
	return func(config *gracefulExitConfig) {
		config.onCloseForError = callback
	}
}

func WithGracefulExitOnCloseForSignal(callback func(os.Signal)) func(*gracefulExitConfig) {
	return func(config *gracefulExitConfig) {
		config.onCloseForSignal = callback
	}
}

type accessOnceMutex struct {
	sync.Mutex
	accessed bool
}

func (m *accessOnceMutex) tryAccess() bool {
	m.Lock()
	defer m.Unlock()

	accessed := m.accessed

	m.accessed = true

	return !accessed
}

// HandleGracefulExit allows you to wait for a given time when a SIGTERM or a SIGINT signal gets sent to the application,
// or when an error requires the application to shutdown, allowing for other resources to gradually complete their execution.
//
// This function does not wait when the context is forcibly closed from another point of the application.
//
// The first parameter needs to be the close function of the main context which drives the execution of the application.
//
// Example:
//
//	ctx, close := context.WithCancel(context.Background())
//	defer close()
//
//	// Use the `ctx` however you want
//
//	graceful, closeWithError := utils.HandleGracefulExit(close, 5 * time.Second)
//	defer graceful.Wait()
func HandleGracefulExit(close context.CancelFunc, duration time.Duration, opts ...func(*gracefulExitConfig)) (*sync.WaitGroup, func(error)) {

	config := &gracefulExitConfig{}

	for _, opt := range opts {
		opt(config)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	closeForErrorChan := make(chan error)

	var closing sync.WaitGroup

	var access accessOnceMutex

	go func() {
		select {
		case err := <-closeForErrorChan:
			if config.onCloseForError != nil {
				config.onCloseForError(err)
			}

			close()
			time.Sleep(duration)

			closing.Done()
		case sig := <-signalChan:

			canDispatchCloseMessage := access.tryAccess()

			if canDispatchCloseMessage {
				closing.Add(1)

				if config.onCloseForSignal != nil {
					config.onCloseForSignal(sig)
				}

				close()
				time.Sleep(duration)

				closing.Done()
			}
		}
	}()

	return &closing, func(err error) {

		canDispatchCloseMessage := access.tryAccess()

		if canDispatchCloseMessage {
			closing.Add(1)
			closeForErrorChan <- err
		}
	}
}
