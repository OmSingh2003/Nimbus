package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	xForwardedForHeader        = "x-forwarded-for"
	userAgentHeader            = "user-agent"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}
	
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("md: %+v\n", md)
		
		// Try to get user agent from gRPC Gateway first
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		} else if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			// Fallback to standard user-agent header for direct gRPC calls
			mtdt.UserAgent = userAgents[0]
		}
		
		// Try to get client IP from X-Forwarded-For header (from gRPC Gateway)
		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			mtdt.ClientIP = clientIPs[0]
		}
	}
	
	// If no client IP from headers, try to get it from peer info (direct gRPC)
	if mtdt.ClientIP == "" {
		if peer, ok := peer.FromContext(ctx); ok {
			mtdt.ClientIP = peer.Addr.String()
		}
	}
	
	return mtdt
}
