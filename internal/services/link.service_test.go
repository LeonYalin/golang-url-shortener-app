package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestLinkService(t *testing.T) {
	suite.Run(t, new(LinkServiceTestSuite))
}

type LinkServiceTestSuite struct {
	suite.Suite
}

func (this *LinkServiceTestSuite) TestCreateLink() {
	assert.True(this.T(), true)
}
