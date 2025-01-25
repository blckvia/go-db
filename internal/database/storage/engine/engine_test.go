package engine_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/blckvia/go-db/internal/database/storage/engine"
	"github.com/blckvia/go-db/mock"
)

type EngineTestSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	mockEngine *mock.MockEngine
	engine     *engine.Engine
}

func (suite *EngineTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockEngine = mock.NewMockEngine(suite.ctrl)
	suite.engine = engine.NewEngine()
}

func (suite *EngineTestSuite) TestSet() {
	suite.T().Parallel()

	suite.mockEngine.EXPECT().Set("key1", "value1").Return(nil).Times(1)
	err := suite.mockEngine.Set("key1", "value1")
	suite.NoError(err)
}

func (suite *EngineTestSuite) TestGet() {
	suite.T().Parallel()

	suite.mockEngine.EXPECT().Get("key1").Return("value1", nil).Times(1)
	val, err := suite.mockEngine.Get("key1")
	suite.NoError(err)
	suite.Equal("value1", val)
}

func (suite *EngineTestSuite) TestDelete() {
	suite.T().Parallel()

	suite.mockEngine.EXPECT().Delete("key1").Return(nil).Times(1)
	err := suite.mockEngine.Delete("key1")
	suite.NoError(err)
}

func (suite *EngineTestSuite) TestDeleteNotFound() {
	suite.mockEngine.EXPECT().Delete("nonexistent_key").Return(fmt.Errorf("key: nonexistent_key not found")).Times(1)
	err := suite.mockEngine.Delete("nonexistent_key")
	suite.Error(err)
}

func TestEngineTestSuite(t *testing.T) {
	suite.Run(t, new(EngineTestSuite))
}
