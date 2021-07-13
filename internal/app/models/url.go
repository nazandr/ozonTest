package models

type URL struct {
	ID    int
	Long  string
	Short string
}

func NewURL() *URL {
	return &URL{
		ID:    0,
		Long:  "",
		Short: "",
	}
}
