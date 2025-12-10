package event

type Interface interface {
	Search(opt Option) ([]Object, error)
}
