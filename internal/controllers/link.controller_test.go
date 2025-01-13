package controllers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestLinkController(t *testing.T) {
	suite.Run(t, new(LinkControllerSuite))
}

type LinkControllerSuite struct {
	suite.Suite
}

func (this *LinkControllerSuite) TestCreate() {
	assert.True(this.T(), true)
}
