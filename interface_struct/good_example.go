package interface_struct

type dataSource interface {
	Data(query string) (string, error)
}

func NewService(ds dataSource) SomeService {
	return SomeService{
		dataSource: ds,
	}
}
