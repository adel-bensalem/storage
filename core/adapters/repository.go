package adapters

type Repository interface {
	get(key string) (string, error)
}
