package apiserver

//API server
type APIServer struct {
}

//New
func New() *APIServer {
	return &APIServer{}
}

func (s *APIServer) Start() error {
	return nil
}
