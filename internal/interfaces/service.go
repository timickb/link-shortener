package interfaces

type Shortener interface {
	CreateLink(link string) (string, error)
	RestoreLink(shortened string) (string, error)
}
