package testutils

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BaseSuite struct {
	suite.Suite
}

func (s BaseSuite) T() *testing.T {
	//TODO implement me
	panic("implement me")
}

func (s BaseSuite) SetT(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

func (s BaseSuite) SetS(suite suite.TestingSuite) {
	//TODO implement me
	panic("implement me")
}
