package api 

import (
  "github.com/OmSingh2003/simple-bank/db/sqlc"
)
// server serves http request for our banking service  
type Server struct {
	 store     *db.Store 
	 router    *gin.Engine
}

//NewServer creates a new HTTP server and etup routing 
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.POST("/accounts",server.createAccount())
	server.router= router 
	return server 
}

func errorResponse(err error) gin.H
