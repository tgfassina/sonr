// @title Highway API
// @version v0.23.0
// @description Manage your Sonr Powered services and blockchain registered types with the Highway API.
// @contact.name Sonr Inc.
// @contact.url https://sonr.io
// @contact.email team@sonr.io
// @license.name OpenGLv3
// @host localhost:8080
// @BasePath /v1

package core

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"

	ctv1 "github.com/sonr-io/sonr/internal/blockchain/x/channel/types"
	"github.com/sonr-io/sonr/internal/highway/x/ipfs"
	"github.com/sonr-io/sonr/internal/highway/x/matrix"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/config"
	hn "github.com/sonr-io/sonr/pkg/host"
	v1 "go.buf.build/grpc/go/sonr-io/highway/v1"
	"google.golang.org/grpc"
)

// Error Definitions
var (
	logger                 = golog.Default.Child("node/highway")
	ErrEmptyQueue          = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery        = errors.New("No SName or PeerID provided.")
	ErrMissingParam        = errors.New("Paramater is missing.")
	ErrProtocolsNotSet     = errors.New("Node Protocol has not been initialized.")
	ErrMethodUnimplemented = errors.New("Method is not implemented.")
)

// HighwayServer is the RPC Service for the Custodian Node.
type HighwayServer struct {
	// Config
	v1.HighwayServer
	ctx    context.Context
	Config *config.Config

	// Clients
	Host     hn.SonrHost
	Cosmos   *client.Cosmos
	Webauthn *client.WebAuthn

	// Http Properties
	Router     *gin.Engine
	HTTPServer *http.Server

	// Grpc Properties
	GRPCServer   *grpc.Server
	GRPCConn     *grpc.ClientConn
	GRPCClient   v1.HighwayClient
	GRPCListener net.Listener

	// Protocols
	channels       map[string]ctv1.Channel
	ipfsProtocol   *ipfs.IPFSProtocol
	matrixProtocol *matrix.MatrixProtocol
}

// setupBaseStub creates the base Highway Server.
func CreateStub(ctx context.Context, c *config.Config) (*HighwayServer, error) {
	node, err := hn.NewMachineHost(ctx, c)
	if err != nil {
		return nil, err
	}

	// Get the Listener for the Host
	lst, err := net.Listen(c.HighwayGRPCNetwork, c.HighwayGRPCEndpoint)
	if err != nil {
		return nil, err
	}

	// Create a new Cosmos Client for Sonr Blockchain
	cosmos, err := client.NewCosmos(ctx, c)
	if err != nil {
		return nil, err
	}

	// Create a new WebAuthn Client for Sonr Blockchain
	webauthn, err := client.NewWebauthn(ctx, c)
	if err != nil {
		return nil, err
	}

	// Create the IPFS Protocol
	ipfs, err := ipfs.New(ctx, node)
	if err != nil {
		return nil, err
	}

	// Create the Matrix Protocol
	matrix, err := matrix.New(ctx, node)
	if err != nil {
		return nil, err
	}

	// Create the RPC Service
	stub := &HighwayServer{
		Cosmos:     cosmos,
		Host:       node,
		ctx:        ctx,
		Router:     gin.Default(),
		GRPCServer: grpc.NewServer(),
		Config:     c,

		GRPCListener:   lst,
		Webauthn:       webauthn,
		ipfsProtocol:   ipfs,
		matrixProtocol: matrix,
		channels:       make(map[string]ctv1.Channel),
	}
	return stub, nil
}

// Serve starts the RPC Service.
func (s *HighwayServer) Serve() {
	// Print the Server Address's
	logger.Infof("Serving RPC Server on %s", s.GRPCListener.Addr().String())
	logger.Infof("Serving HTTP Server on %s", s.Config.HighwayHTTPEndpoint)

	// Start the gRPC Server
	go func() {
		// Start gRPC server (and proxy calls to gRPC server endpoint)
		if err := s.GRPCServer.Serve(s.GRPCListener); err != nil {
			logger.Errorf("%s - Failed to start HTTP server", err)
		}
	}()

	// Start HTTP server on a separate goroutine
	go func() {
		// Start HTTP server (and proxy calls to gRPC server endpoint)
		if err := s.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("%s - Failed to start HTTP server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	<-quit
	logger.Warn("Shutting down HTTP server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.HTTPServer.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown: ", err)
	}

	// Close the gRPC server
	logger.Warn("Shutting down gRPC server...")
	s.GRPCServer.GracefulStop()
	logger.Info("Goodbye!")
}
