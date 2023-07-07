package parts

import "context"

type IController interface {
	Fetch(c context.Context) error
}

type IDao interface {
	Fetch(c context.Context, id int) (string, error)
}

type IService interface {
	Fetch(c context.Context, id string) (string, error)
}
