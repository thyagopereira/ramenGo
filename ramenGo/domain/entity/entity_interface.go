package entity

type Entity interface {
	validate() (bool, error)
}
