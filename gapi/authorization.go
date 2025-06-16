package gapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/OmSingh2003/vaultguard-api/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) authorizeUser(ctx context.Context, accessibleRoles []string) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	// Since we don't have role-based access control yet, we'll skip the role check
	// TODO: Implement role-based access control when User model includes roles
	// if !hasPermission(payload.Username, accessibleRoles) {
	//	return nil, fmt.Errorf("permission denied")
	// }

	return payload, nil
}

// getAuthPayload extracts the authentication payload from context
func (server *Server) getAuthPayload(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	return payload, nil
}

// AuthorizationInterceptor is a gRPC unary interceptor for authorization
func (server *Server) AuthorizationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// List of methods that don't require authentication
	publicMethods := map[string]bool{
		"/pb.VaultguardAPI/CreateUser": true,
		"/pb.VaultguardAPI/LoginUser":  true,
	}

	// Check if this method requires authentication
	if !publicMethods[info.FullMethod] {
		// Verify the token
		_, err := server.getAuthPayload(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
	}

	// Call the actual handler
	return handler(ctx, req)
}

func hasPermission(userRole string, accessibleRoles []string) bool {
	for _, role := range accessibleRoles {
		if userRole == role {
			return true
		}
	}
	return false
}
