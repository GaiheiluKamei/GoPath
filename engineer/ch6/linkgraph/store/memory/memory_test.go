package memory

import (
	gc "gopkg.in/check.v1"
	"testing"
)

var _ = gc.Suite(new(InMemoryGraph))

func Test(t *testing.T) {gc.TestingT(t)}

type InMemoryGraphTestSuite struct {
	// graphtest.SuiteBase
}

// Register our test-suite with go test.
func (s *InMemoryGraphTestSuite) SetUpTest(c *gc.C) {
	// s.SetGraph(NewInMemoryGraph())
}
