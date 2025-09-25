package user

type Service struct {}

func New() Service {
    return Service{}
}

func (us Service) Exists(un UserName) bool {
    // TODO: Implement
    return false
}