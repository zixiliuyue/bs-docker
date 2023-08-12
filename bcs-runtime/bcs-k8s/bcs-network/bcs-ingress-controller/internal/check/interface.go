

package check

// Checker func Run() will be periodical executed by CheckRunner
type Checker interface {
	Run()
}
