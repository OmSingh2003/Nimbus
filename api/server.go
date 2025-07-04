package api

import (
	"fmt"
	"net/http"

	db "github.com/OmSingh2003/nimbus/db/sqlc"
	"github.com/OmSingh2003/nimbus/token"
	"github.com/OmSingh2003/nimbus/util"
	"github.com/OmSingh2003/nimbus/worker"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service
type Server struct {
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
	router          *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token master: %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setUpRouter()

	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	// User routes
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/token.renew_access", server.renewAccessToken)
	router.GET("/verify_email", server.verifyEmail)
	router.POST("/resend_verification", server.resendVerificationEmail)
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// Account routes
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)

	// Transfer routes
	authRoutes.POST("/transfers", server.createTransfer)
	authRoutes.GET("/transfers", server.listTransfers)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// ServeHTTP implements http.Handler interface for testing
func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}
