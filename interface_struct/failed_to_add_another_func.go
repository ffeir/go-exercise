package interface_struct

import "testing"

type DataSource interface {
	Data(query string) (string, error)
	Data2(query string) (string, error)
}

func NewDataSource() DataSource {
	return MemoryDataSource{}
}

type MemoryDataSource struct{}

func (mds MemoryDataSource) Data(query string) (string, error) {
	return "MemoryDataSource", nil
}

func NewMemoryDataSource() DataSource {
	return MemoryDataSource{}
}

type SomeService struct {
	dataSource DataSource
}

func NewService(ds DataSource) SomeService {
	return SomeService{
		dataSource: ds,
	}
}

type mockDataSource struct{}

func (mockDataSource) Data(query string) (string, error) {
	return "mockDataSource", nil
}

func TestServiceAction(t *testing.T) {
	NewService(mockDataSource{})
}
