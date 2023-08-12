

// Package concurrency xxx
package concurrency

// Concurrency for limited resources
type Concurrency struct {
	c chan struct{}
}

// NewConcurrency create object and init it
func NewConcurrency(num int) *Concurrency {
	conNum := &Concurrency{}
	conNum.c = make(chan struct{}, num)

	for i := 0; i < num; i++ {
		conNum.c <- struct{}{}
	}

	return conNum
}

// Add allocate 1 if there is 1 un-used resources
func (con *Concurrency) Add() {
	if con == nil {
		return
	}

	<-con.c
}

// Done release 1 resource to resourcePool
func (con *Concurrency) Done() {
	if con == nil {
		return
	}

	select {
	case con.c <- struct{}{}:
	default:
	}
}
