package dataclass

type MetaData interface {
	GetMovie(title string, rating Rating)
	GetTitle() string
	SetTitle(newTitle string)
	GetRating() Rating
	SetRating(newRating Rating)
}

type Movie struct {
	Title  string
	rating Rating
}

type Rating string

const (
	R = "average"
	A = "good"
	D = "bad"
)

func (m *Movie) GetMovie(title string, rating Rating) {
	m.Title = title
	m.rating = rating
}

func (mov *Movie) GetTitle() string {
	return mov.Title
}

func (mov *Movie) SetTitle(newTitle string) {
	mov.Title = newTitle
}

func (mov *Movie) GetRating() Rating {
	return mov.rating
}

func (mov *Movie) SetRating(newRating Rating) {
	mov.rating = newRating
}
