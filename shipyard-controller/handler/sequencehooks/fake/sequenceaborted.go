// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"github.com/keptn/keptn/shipyard-controller/models"
	"sync"
)

// ISequenceAbortedHookMock is a mock implementation of sequencehooks.ISequenceAbortedHook.
//
// 	func TestSomethingThatUsesISequenceAbortedHook(t *testing.T) {
//
// 		// make and configure a mocked sequencehooks.ISequenceAbortedHook
// 		mockedISequenceAbortedHook := &ISequenceAbortedHookMock{
// 			OnSequenceAbortedFunc: func(event models.Event)  {
// 				panic("mock out the OnSequenceAborted method")
// 			},
// 		}
//
// 		// use mockedISequenceAbortedHook in code that requires sequencehooks.ISequenceAbortedHook
// 		// and then make assertions.
//
// 	}
type ISequenceAbortedHookMock struct {
	// OnSequenceAbortedFunc mocks the OnSequenceAborted method.
	OnSequenceAbortedFunc func(event models.Event)

	// calls tracks calls to the methods.
	calls struct {
		// OnSequenceAborted holds details about calls to the OnSequenceAborted method.
		OnSequenceAborted []struct {
			// Event is the event argument value.
			Event models.Event
		}
	}
	lockOnSequenceAborted sync.RWMutex
}

// OnSequenceAborted calls OnSequenceAbortedFunc.
func (mock *ISequenceAbortedHookMock) OnSequenceAborted(event models.Event) {
	if mock.OnSequenceAbortedFunc == nil {
		panic("ISequenceAbortedHookMock.OnSequenceAbortedFunc: method is nil but ISequenceAbortedHook.OnSequenceAborted was just called")
	}
	callInfo := struct {
		Event models.Event
	}{
		Event: event,
	}
	mock.lockOnSequenceAborted.Lock()
	mock.calls.OnSequenceAborted = append(mock.calls.OnSequenceAborted, callInfo)
	mock.lockOnSequenceAborted.Unlock()
	mock.OnSequenceAbortedFunc(event)
}

// OnSequenceAbortedCalls gets all the calls that were made to OnSequenceAborted.
// Check the length with:
//     len(mockedISequenceAbortedHook.OnSequenceAbortedCalls())
func (mock *ISequenceAbortedHookMock) OnSequenceAbortedCalls() []struct {
	Event models.Event
} {
	var calls []struct {
		Event models.Event
	}
	mock.lockOnSequenceAborted.RLock()
	calls = mock.calls.OnSequenceAborted
	mock.lockOnSequenceAborted.RUnlock()
	return calls
}
