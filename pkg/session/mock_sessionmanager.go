// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package session

import (
	"sync"
)

var (
	lockManagerMockCreate sync.RWMutex
	lockManagerMockGet    sync.RWMutex
	lockManagerMockSetCtx sync.RWMutex
)

// Ensure, that ManagerMock does implement Manager.
// If this is not the case, regenerate this file with moq.
var _ Manager = &ManagerMock{}

// ManagerMock is a mock implementation of Manager.
//
//     func TestSomethingThatUsesManager(t *testing.T) {
//
//         // make and configure a mocked Manager
//         mockedManager := &ManagerMock{
//             CreateFunc: func(token string, username string, organization string) error {
// 	               panic("mock out the Create method")
//             },
//             GetFunc: func() (*Session, error) {
// 	               panic("mock out the Get method")
//             },
//             SetCtxFunc: func(ctx string) error {
// 	               panic("mock out the SetCtx method")
//             },
//         }
//
//         // use mockedManager in code that requires Manager
//         // and then make assertions.
//
//     }
type ManagerMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(token string, username string, organization string) error

	// GetFunc mocks the Get method.
	GetFunc func() (*Session, error)

	// SetCtxFunc mocks the SetCtx method.
	SetCtxFunc func(ctx string) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Token is the token argument value.
			Token string
			// Username is the username argument value.
			Username string
			// Organization is the organization argument value.
			Organization string
		}
		// Get holds details about calls to the Get method.
		Get []struct {
		}
		// SetCtx holds details about calls to the SetCtx method.
		SetCtx []struct {
			// Ctx is the ctx argument value.
			Ctx string
		}
	}
}

// Create calls CreateFunc.
func (mock *ManagerMock) Create(token string, username string, organization string) error {
	if mock.CreateFunc == nil {
		panic("ManagerMock.CreateFunc: method is nil but Manager.Create was just called")
	}
	callInfo := struct {
		Token        string
		Username     string
		Organization string
	}{
		Token:        token,
		Username:     username,
		Organization: organization,
	}
	lockManagerMockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	lockManagerMockCreate.Unlock()
	return mock.CreateFunc(token, username, organization)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedManager.CreateCalls())
func (mock *ManagerMock) CreateCalls() []struct {
	Token        string
	Username     string
	Organization string
} {
	var calls []struct {
		Token        string
		Username     string
		Organization string
	}
	lockManagerMockCreate.RLock()
	calls = mock.calls.Create
	lockManagerMockCreate.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *ManagerMock) Get() (*Session, error) {
	if mock.GetFunc == nil {
		panic("ManagerMock.GetFunc: method is nil but Manager.Get was just called")
	}
	callInfo := struct {
	}{}
	lockManagerMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockManagerMockGet.Unlock()
	return mock.GetFunc()
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedManager.GetCalls())
func (mock *ManagerMock) GetCalls() []struct {
} {
	var calls []struct {
	}
	lockManagerMockGet.RLock()
	calls = mock.calls.Get
	lockManagerMockGet.RUnlock()
	return calls
}

// SetCtx calls SetCtxFunc.
func (mock *ManagerMock) SetCtx(ctx string) error {
	if mock.SetCtxFunc == nil {
		panic("ManagerMock.SetCtxFunc: method is nil but Manager.SetCtx was just called")
	}
	callInfo := struct {
		Ctx string
	}{
		Ctx: ctx,
	}
	lockManagerMockSetCtx.Lock()
	mock.calls.SetCtx = append(mock.calls.SetCtx, callInfo)
	lockManagerMockSetCtx.Unlock()
	return mock.SetCtxFunc(ctx)
}

// SetCtxCalls gets all the calls that were made to SetCtx.
// Check the length with:
//     len(mockedManager.SetCtxCalls())
func (mock *ManagerMock) SetCtxCalls() []struct {
	Ctx string
} {
	var calls []struct {
		Ctx string
	}
	lockManagerMockSetCtx.RLock()
	calls = mock.calls.SetCtx
	lockManagerMockSetCtx.RUnlock()
	return calls
}
