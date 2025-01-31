package storage_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/blckvia/go-db/internal/database/storage"
	"github.com/blckvia/go-db/mock"
)

type StorageTestSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	mockEngine *mock.MockEngine
	storage    *storage.Storage
}

func (suite *StorageTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockEngine = mock.NewMockEngine(suite.ctrl)
	suite.storage = storage.NewStorage(suite.mockEngine, zap.NewNop())
}

func (suite *StorageTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *StorageTestSuite) TestSet() {
	suite.T().Parallel()

	suite.mockEngine.EXPECT().Set("key1", "value1").Return(nil).Times(1)
	err := suite.storage.Set(context.Background(), "key1", "value1")
	suite.NoError(err)
}

func (suite *StorageTestSuite) TestGet() {
	suite.T().Parallel()

	suite.mockEngine.EXPECT().Get("key1").Return("value1", nil).Times(1)
	val, err := suite.storage.Get(context.Background(), "key1")
	suite.NoError(err)
	suite.Equal("value1", val)
}

func (suite *StorageTestSuite) TestDelete() {
	suite.T().Parallel()

	suite.mockEngine.EXPECT().Delete("key1").Return(nil).Times(1)
	err := suite.storage.Delete(context.Background(), "key1")
	suite.NoError(err)
}

func (suite *StorageTestSuite) TestDeleteNotFound() {
	suite.T().Parallel()

	suite.mockEngine.EXPECT().Delete("nonexistent_key").Return(fmt.Errorf("key: nonexistent_key not found")).Times(1)
	err := suite.storage.Delete(context.Background(), "nonexistent_key")
	suite.Error(err)
}

func TestStorageTestSuite(t *testing.T) {
	suite.Run(t, new(StorageTestSuite))
}
