package adapter

type RepoAdapter interface {
	// Close()
	FindByID(id string, data interface{}, query string) (interface{}, error)
	List(mdl interface{}, query string) (interface{}, error)
	Create(data interface{}, query string) (int64, error)
	Update(data interface{}, id, query string) (int64, error)
	Delete(id, query string) (int64, error)
	Auth(uname, pass, query string, data interface{}) (interface{}, error)
}
