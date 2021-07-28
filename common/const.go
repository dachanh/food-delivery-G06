package common

const (
	DbTypeUser = 2
)
const (
	CurrentUser = "user"
)

type Requester interface {
	GetRole() string
	GetEmail() string
	GetUsedID() int
}
