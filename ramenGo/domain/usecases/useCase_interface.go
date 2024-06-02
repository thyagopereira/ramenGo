package usecases

type Usecase interface {
	Execute() (interface{}, error)
}
