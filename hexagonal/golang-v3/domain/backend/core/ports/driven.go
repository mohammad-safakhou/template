package ports

type HelloRepository interface {
	Get() string
	Save(string)
}
