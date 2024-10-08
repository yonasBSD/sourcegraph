// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package subscriptionlicensechecksservice

import (
	"context"
	"sync"
	"time"

	subscriptions "github.com/sourcegraph/sourcegraph/cmd/enterprise-portal/internal/database/subscriptions"
	slack "github.com/sourcegraph/sourcegraph/internal/slack"
)

// MockStoreV1 is a mock implementation of the StoreV1 interface (from the
// package
// github.com/sourcegraph/sourcegraph/cmd/enterprise-portal/internal/subscriptionlicensechecksservice)
// used for unit testing.
type MockStoreV1 struct {
	// BypassAllLicenseChecksFunc is an instance of a mock function object
	// controlling the behavior of the method BypassAllLicenseChecks.
	BypassAllLicenseChecksFunc *StoreV1BypassAllLicenseChecksFunc
	// GetByLicenseKeyFunc is an instance of a mock function object
	// controlling the behavior of the method GetByLicenseKey.
	GetByLicenseKeyFunc *StoreV1GetByLicenseKeyFunc
	// GetByLicenseKeyHashFunc is an instance of a mock function object
	// controlling the behavior of the method GetByLicenseKeyHash.
	GetByLicenseKeyHashFunc *StoreV1GetByLicenseKeyHashFunc
	// GetSubscriptionFunc is an instance of a mock function object
	// controlling the behavior of the method GetSubscription.
	GetSubscriptionFunc *StoreV1GetSubscriptionFunc
	// NowFunc is an instance of a mock function object controlling the
	// behavior of the method Now.
	NowFunc *StoreV1NowFunc
	// PostToSlackFunc is an instance of a mock function object controlling
	// the behavior of the method PostToSlack.
	PostToSlackFunc *StoreV1PostToSlackFunc
	// SetDetectedInstanceFunc is an instance of a mock function object
	// controlling the behavior of the method SetDetectedInstance.
	SetDetectedInstanceFunc *StoreV1SetDetectedInstanceFunc
}

// NewMockStoreV1 creates a new mock of the StoreV1 interface. All methods
// return zero values for all results, unless overwritten.
func NewMockStoreV1() *MockStoreV1 {
	return &MockStoreV1{
		BypassAllLicenseChecksFunc: &StoreV1BypassAllLicenseChecksFunc{
			defaultHook: func() (r0 bool) {
				return
			},
		},
		GetByLicenseKeyFunc: &StoreV1GetByLicenseKeyFunc{
			defaultHook: func(context.Context, string) (r0 *subscriptions.SubscriptionLicense, r1 error) {
				return
			},
		},
		GetByLicenseKeyHashFunc: &StoreV1GetByLicenseKeyHashFunc{
			defaultHook: func(context.Context, string) (r0 *subscriptions.SubscriptionLicense, r1 error) {
				return
			},
		},
		GetSubscriptionFunc: &StoreV1GetSubscriptionFunc{
			defaultHook: func(context.Context, string) (r0 *subscriptions.Subscription, r1 error) {
				return
			},
		},
		NowFunc: &StoreV1NowFunc{
			defaultHook: func() (r0 time.Time) {
				return
			},
		},
		PostToSlackFunc: &StoreV1PostToSlackFunc{
			defaultHook: func(context.Context, *slack.Payload) (r0 error) {
				return
			},
		},
		SetDetectedInstanceFunc: &StoreV1SetDetectedInstanceFunc{
			defaultHook: func(context.Context, string, string) (r0 error) {
				return
			},
		},
	}
}

// NewStrictMockStoreV1 creates a new mock of the StoreV1 interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockStoreV1() *MockStoreV1 {
	return &MockStoreV1{
		BypassAllLicenseChecksFunc: &StoreV1BypassAllLicenseChecksFunc{
			defaultHook: func() bool {
				panic("unexpected invocation of MockStoreV1.BypassAllLicenseChecks")
			},
		},
		GetByLicenseKeyFunc: &StoreV1GetByLicenseKeyFunc{
			defaultHook: func(context.Context, string) (*subscriptions.SubscriptionLicense, error) {
				panic("unexpected invocation of MockStoreV1.GetByLicenseKey")
			},
		},
		GetByLicenseKeyHashFunc: &StoreV1GetByLicenseKeyHashFunc{
			defaultHook: func(context.Context, string) (*subscriptions.SubscriptionLicense, error) {
				panic("unexpected invocation of MockStoreV1.GetByLicenseKeyHash")
			},
		},
		GetSubscriptionFunc: &StoreV1GetSubscriptionFunc{
			defaultHook: func(context.Context, string) (*subscriptions.Subscription, error) {
				panic("unexpected invocation of MockStoreV1.GetSubscription")
			},
		},
		NowFunc: &StoreV1NowFunc{
			defaultHook: func() time.Time {
				panic("unexpected invocation of MockStoreV1.Now")
			},
		},
		PostToSlackFunc: &StoreV1PostToSlackFunc{
			defaultHook: func(context.Context, *slack.Payload) error {
				panic("unexpected invocation of MockStoreV1.PostToSlack")
			},
		},
		SetDetectedInstanceFunc: &StoreV1SetDetectedInstanceFunc{
			defaultHook: func(context.Context, string, string) error {
				panic("unexpected invocation of MockStoreV1.SetDetectedInstance")
			},
		},
	}
}

// NewMockStoreV1From creates a new mock of the MockStoreV1 interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockStoreV1From(i StoreV1) *MockStoreV1 {
	return &MockStoreV1{
		BypassAllLicenseChecksFunc: &StoreV1BypassAllLicenseChecksFunc{
			defaultHook: i.BypassAllLicenseChecks,
		},
		GetByLicenseKeyFunc: &StoreV1GetByLicenseKeyFunc{
			defaultHook: i.GetByLicenseKey,
		},
		GetByLicenseKeyHashFunc: &StoreV1GetByLicenseKeyHashFunc{
			defaultHook: i.GetByLicenseKeyHash,
		},
		GetSubscriptionFunc: &StoreV1GetSubscriptionFunc{
			defaultHook: i.GetSubscription,
		},
		NowFunc: &StoreV1NowFunc{
			defaultHook: i.Now,
		},
		PostToSlackFunc: &StoreV1PostToSlackFunc{
			defaultHook: i.PostToSlack,
		},
		SetDetectedInstanceFunc: &StoreV1SetDetectedInstanceFunc{
			defaultHook: i.SetDetectedInstance,
		},
	}
}

// StoreV1BypassAllLicenseChecksFunc describes the behavior when the
// BypassAllLicenseChecks method of the parent MockStoreV1 instance is
// invoked.
type StoreV1BypassAllLicenseChecksFunc struct {
	defaultHook func() bool
	hooks       []func() bool
	history     []StoreV1BypassAllLicenseChecksFuncCall
	mutex       sync.Mutex
}

// BypassAllLicenseChecks delegates to the next hook function in the queue
// and stores the parameter and result values of this invocation.
func (m *MockStoreV1) BypassAllLicenseChecks() bool {
	r0 := m.BypassAllLicenseChecksFunc.nextHook()()
	m.BypassAllLicenseChecksFunc.appendCall(StoreV1BypassAllLicenseChecksFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the
// BypassAllLicenseChecks method of the parent MockStoreV1 instance is
// invoked and the hook queue is empty.
func (f *StoreV1BypassAllLicenseChecksFunc) SetDefaultHook(hook func() bool) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// BypassAllLicenseChecks method of the parent MockStoreV1 instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *StoreV1BypassAllLicenseChecksFunc) PushHook(hook func() bool) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1BypassAllLicenseChecksFunc) SetDefaultReturn(r0 bool) {
	f.SetDefaultHook(func() bool {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1BypassAllLicenseChecksFunc) PushReturn(r0 bool) {
	f.PushHook(func() bool {
		return r0
	})
}

func (f *StoreV1BypassAllLicenseChecksFunc) nextHook() func() bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1BypassAllLicenseChecksFunc) appendCall(r0 StoreV1BypassAllLicenseChecksFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreV1BypassAllLicenseChecksFuncCall
// objects describing the invocations of this function.
func (f *StoreV1BypassAllLicenseChecksFunc) History() []StoreV1BypassAllLicenseChecksFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1BypassAllLicenseChecksFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1BypassAllLicenseChecksFuncCall is an object that describes an
// invocation of method BypassAllLicenseChecks on an instance of
// MockStoreV1.
type StoreV1BypassAllLicenseChecksFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 bool
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1BypassAllLicenseChecksFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1BypassAllLicenseChecksFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// StoreV1GetByLicenseKeyFunc describes the behavior when the
// GetByLicenseKey method of the parent MockStoreV1 instance is invoked.
type StoreV1GetByLicenseKeyFunc struct {
	defaultHook func(context.Context, string) (*subscriptions.SubscriptionLicense, error)
	hooks       []func(context.Context, string) (*subscriptions.SubscriptionLicense, error)
	history     []StoreV1GetByLicenseKeyFuncCall
	mutex       sync.Mutex
}

// GetByLicenseKey delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStoreV1) GetByLicenseKey(v0 context.Context, v1 string) (*subscriptions.SubscriptionLicense, error) {
	r0, r1 := m.GetByLicenseKeyFunc.nextHook()(v0, v1)
	m.GetByLicenseKeyFunc.appendCall(StoreV1GetByLicenseKeyFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetByLicenseKey
// method of the parent MockStoreV1 instance is invoked and the hook queue
// is empty.
func (f *StoreV1GetByLicenseKeyFunc) SetDefaultHook(hook func(context.Context, string) (*subscriptions.SubscriptionLicense, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetByLicenseKey method of the parent MockStoreV1 instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreV1GetByLicenseKeyFunc) PushHook(hook func(context.Context, string) (*subscriptions.SubscriptionLicense, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1GetByLicenseKeyFunc) SetDefaultReturn(r0 *subscriptions.SubscriptionLicense, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*subscriptions.SubscriptionLicense, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1GetByLicenseKeyFunc) PushReturn(r0 *subscriptions.SubscriptionLicense, r1 error) {
	f.PushHook(func(context.Context, string) (*subscriptions.SubscriptionLicense, error) {
		return r0, r1
	})
}

func (f *StoreV1GetByLicenseKeyFunc) nextHook() func(context.Context, string) (*subscriptions.SubscriptionLicense, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1GetByLicenseKeyFunc) appendCall(r0 StoreV1GetByLicenseKeyFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreV1GetByLicenseKeyFuncCall objects
// describing the invocations of this function.
func (f *StoreV1GetByLicenseKeyFunc) History() []StoreV1GetByLicenseKeyFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1GetByLicenseKeyFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1GetByLicenseKeyFuncCall is an object that describes an invocation
// of method GetByLicenseKey on an instance of MockStoreV1.
type StoreV1GetByLicenseKeyFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *subscriptions.SubscriptionLicense
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1GetByLicenseKeyFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1GetByLicenseKeyFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreV1GetByLicenseKeyHashFunc describes the behavior when the
// GetByLicenseKeyHash method of the parent MockStoreV1 instance is invoked.
type StoreV1GetByLicenseKeyHashFunc struct {
	defaultHook func(context.Context, string) (*subscriptions.SubscriptionLicense, error)
	hooks       []func(context.Context, string) (*subscriptions.SubscriptionLicense, error)
	history     []StoreV1GetByLicenseKeyHashFuncCall
	mutex       sync.Mutex
}

// GetByLicenseKeyHash delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStoreV1) GetByLicenseKeyHash(v0 context.Context, v1 string) (*subscriptions.SubscriptionLicense, error) {
	r0, r1 := m.GetByLicenseKeyHashFunc.nextHook()(v0, v1)
	m.GetByLicenseKeyHashFunc.appendCall(StoreV1GetByLicenseKeyHashFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetByLicenseKeyHash
// method of the parent MockStoreV1 instance is invoked and the hook queue
// is empty.
func (f *StoreV1GetByLicenseKeyHashFunc) SetDefaultHook(hook func(context.Context, string) (*subscriptions.SubscriptionLicense, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetByLicenseKeyHash method of the parent MockStoreV1 instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreV1GetByLicenseKeyHashFunc) PushHook(hook func(context.Context, string) (*subscriptions.SubscriptionLicense, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1GetByLicenseKeyHashFunc) SetDefaultReturn(r0 *subscriptions.SubscriptionLicense, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*subscriptions.SubscriptionLicense, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1GetByLicenseKeyHashFunc) PushReturn(r0 *subscriptions.SubscriptionLicense, r1 error) {
	f.PushHook(func(context.Context, string) (*subscriptions.SubscriptionLicense, error) {
		return r0, r1
	})
}

func (f *StoreV1GetByLicenseKeyHashFunc) nextHook() func(context.Context, string) (*subscriptions.SubscriptionLicense, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1GetByLicenseKeyHashFunc) appendCall(r0 StoreV1GetByLicenseKeyHashFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreV1GetByLicenseKeyHashFuncCall objects
// describing the invocations of this function.
func (f *StoreV1GetByLicenseKeyHashFunc) History() []StoreV1GetByLicenseKeyHashFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1GetByLicenseKeyHashFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1GetByLicenseKeyHashFuncCall is an object that describes an
// invocation of method GetByLicenseKeyHash on an instance of MockStoreV1.
type StoreV1GetByLicenseKeyHashFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *subscriptions.SubscriptionLicense
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1GetByLicenseKeyHashFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1GetByLicenseKeyHashFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreV1GetSubscriptionFunc describes the behavior when the
// GetSubscription method of the parent MockStoreV1 instance is invoked.
type StoreV1GetSubscriptionFunc struct {
	defaultHook func(context.Context, string) (*subscriptions.Subscription, error)
	hooks       []func(context.Context, string) (*subscriptions.Subscription, error)
	history     []StoreV1GetSubscriptionFuncCall
	mutex       sync.Mutex
}

// GetSubscription delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStoreV1) GetSubscription(v0 context.Context, v1 string) (*subscriptions.Subscription, error) {
	r0, r1 := m.GetSubscriptionFunc.nextHook()(v0, v1)
	m.GetSubscriptionFunc.appendCall(StoreV1GetSubscriptionFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetSubscription
// method of the parent MockStoreV1 instance is invoked and the hook queue
// is empty.
func (f *StoreV1GetSubscriptionFunc) SetDefaultHook(hook func(context.Context, string) (*subscriptions.Subscription, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetSubscription method of the parent MockStoreV1 instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreV1GetSubscriptionFunc) PushHook(hook func(context.Context, string) (*subscriptions.Subscription, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1GetSubscriptionFunc) SetDefaultReturn(r0 *subscriptions.Subscription, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*subscriptions.Subscription, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1GetSubscriptionFunc) PushReturn(r0 *subscriptions.Subscription, r1 error) {
	f.PushHook(func(context.Context, string) (*subscriptions.Subscription, error) {
		return r0, r1
	})
}

func (f *StoreV1GetSubscriptionFunc) nextHook() func(context.Context, string) (*subscriptions.Subscription, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1GetSubscriptionFunc) appendCall(r0 StoreV1GetSubscriptionFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreV1GetSubscriptionFuncCall objects
// describing the invocations of this function.
func (f *StoreV1GetSubscriptionFunc) History() []StoreV1GetSubscriptionFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1GetSubscriptionFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1GetSubscriptionFuncCall is an object that describes an invocation
// of method GetSubscription on an instance of MockStoreV1.
type StoreV1GetSubscriptionFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *subscriptions.Subscription
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1GetSubscriptionFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1GetSubscriptionFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreV1NowFunc describes the behavior when the Now method of the parent
// MockStoreV1 instance is invoked.
type StoreV1NowFunc struct {
	defaultHook func() time.Time
	hooks       []func() time.Time
	history     []StoreV1NowFuncCall
	mutex       sync.Mutex
}

// Now delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStoreV1) Now() time.Time {
	r0 := m.NowFunc.nextHook()()
	m.NowFunc.appendCall(StoreV1NowFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the Now method of the
// parent MockStoreV1 instance is invoked and the hook queue is empty.
func (f *StoreV1NowFunc) SetDefaultHook(hook func() time.Time) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Now method of the parent MockStoreV1 instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StoreV1NowFunc) PushHook(hook func() time.Time) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1NowFunc) SetDefaultReturn(r0 time.Time) {
	f.SetDefaultHook(func() time.Time {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1NowFunc) PushReturn(r0 time.Time) {
	f.PushHook(func() time.Time {
		return r0
	})
}

func (f *StoreV1NowFunc) nextHook() func() time.Time {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1NowFunc) appendCall(r0 StoreV1NowFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreV1NowFuncCall objects describing the
// invocations of this function.
func (f *StoreV1NowFunc) History() []StoreV1NowFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1NowFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1NowFuncCall is an object that describes an invocation of method
// Now on an instance of MockStoreV1.
type StoreV1NowFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 time.Time
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1NowFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1NowFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// StoreV1PostToSlackFunc describes the behavior when the PostToSlack method
// of the parent MockStoreV1 instance is invoked.
type StoreV1PostToSlackFunc struct {
	defaultHook func(context.Context, *slack.Payload) error
	hooks       []func(context.Context, *slack.Payload) error
	history     []StoreV1PostToSlackFuncCall
	mutex       sync.Mutex
}

// PostToSlack delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockStoreV1) PostToSlack(v0 context.Context, v1 *slack.Payload) error {
	r0 := m.PostToSlackFunc.nextHook()(v0, v1)
	m.PostToSlackFunc.appendCall(StoreV1PostToSlackFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the PostToSlack method
// of the parent MockStoreV1 instance is invoked and the hook queue is
// empty.
func (f *StoreV1PostToSlackFunc) SetDefaultHook(hook func(context.Context, *slack.Payload) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// PostToSlack method of the parent MockStoreV1 instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *StoreV1PostToSlackFunc) PushHook(hook func(context.Context, *slack.Payload) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1PostToSlackFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, *slack.Payload) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1PostToSlackFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, *slack.Payload) error {
		return r0
	})
}

func (f *StoreV1PostToSlackFunc) nextHook() func(context.Context, *slack.Payload) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1PostToSlackFunc) appendCall(r0 StoreV1PostToSlackFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreV1PostToSlackFuncCall objects
// describing the invocations of this function.
func (f *StoreV1PostToSlackFunc) History() []StoreV1PostToSlackFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1PostToSlackFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1PostToSlackFuncCall is an object that describes an invocation of
// method PostToSlack on an instance of MockStoreV1.
type StoreV1PostToSlackFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *slack.Payload
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1PostToSlackFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1PostToSlackFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// StoreV1SetDetectedInstanceFunc describes the behavior when the
// SetDetectedInstance method of the parent MockStoreV1 instance is invoked.
type StoreV1SetDetectedInstanceFunc struct {
	defaultHook func(context.Context, string, string) error
	hooks       []func(context.Context, string, string) error
	history     []StoreV1SetDetectedInstanceFuncCall
	mutex       sync.Mutex
}

// SetDetectedInstance delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStoreV1) SetDetectedInstance(v0 context.Context, v1 string, v2 string) error {
	r0 := m.SetDetectedInstanceFunc.nextHook()(v0, v1, v2)
	m.SetDetectedInstanceFunc.appendCall(StoreV1SetDetectedInstanceFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the SetDetectedInstance
// method of the parent MockStoreV1 instance is invoked and the hook queue
// is empty.
func (f *StoreV1SetDetectedInstanceFunc) SetDefaultHook(hook func(context.Context, string, string) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// SetDetectedInstance method of the parent MockStoreV1 instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreV1SetDetectedInstanceFunc) PushHook(hook func(context.Context, string, string) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1SetDetectedInstanceFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, string, string) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1SetDetectedInstanceFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, string, string) error {
		return r0
	})
}

func (f *StoreV1SetDetectedInstanceFunc) nextHook() func(context.Context, string, string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1SetDetectedInstanceFunc) appendCall(r0 StoreV1SetDetectedInstanceFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreV1SetDetectedInstanceFuncCall objects
// describing the invocations of this function.
func (f *StoreV1SetDetectedInstanceFunc) History() []StoreV1SetDetectedInstanceFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1SetDetectedInstanceFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1SetDetectedInstanceFuncCall is an object that describes an
// invocation of method SetDetectedInstance on an instance of MockStoreV1.
type StoreV1SetDetectedInstanceFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1SetDetectedInstanceFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1SetDetectedInstanceFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
