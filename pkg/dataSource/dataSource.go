package dataSource

type DataSource interface {
	Open(string) error
	Close()
	InsertLinks(string, string) error
	FindLongLink(string) (string, error)
	DoesExistShortname(string) (bool, error)
}
