//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/weaviate/weaviate/adapters/handlers/rest/state"
	"github.com/weaviate/weaviate/entities/dto"
	"github.com/weaviate/weaviate/entities/schema"
	pb "github.com/weaviate/weaviate/grpc"
	"github.com/weaviate/weaviate/usecases/auth/authentication/composer"
	schemaManager "github.com/weaviate/weaviate/usecases/schema"
	"github.com/weaviate/weaviate/usecases/traverser"
	"google.golang.org/grpc"
)

func CreateGRPCServer(state *state.State) *GRPCServer {
	s := grpc.NewServer()

	pb.RegisterWeaviateServer(s, &Server{
		traverser: state.Traverser,
		authComposer: composer.New(
			state.ServerConfig.Config.Authentication,
			state.APIKey, state.OIDC),
		allowAnonymousAccess: state.ServerConfig.Config.Authentication.AnonymousAccess.Enabled,
		schemaManager:        state.SchemaManager,
	})

	return &GRPCServer{s}
}

func StartAndListen(s *GRPCServer, state *state.State) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d",
		state.ServerConfig.Config.GRPC.Port))
	if err != nil {
		return err
	}
	state.Logger.WithField("action", "grpc_startup").
		Infof("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

type GRPCServer struct {
	*grpc.Server
}

type Server struct {
	pb.UnimplementedWeaviateServer
	traverser            *traverser.Traverser
	authComposer         composer.TokenFunc
	allowAnonymousAccess bool
	schemaManager        *schemaManager.Manager
}

func (s *Server) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchReply, error) {
	before := time.Now()

	principal, err := s.principalFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("extract auth: %w", err)
	}

	scheme := s.schemaManager.GetSchemaSkipAuth()

	searchParams, err := searchParamsFromProto(req, scheme)
	if err != nil {
		return nil, fmt.Errorf("extract params: %w", err)
	}

	if err := s.validateClassAndProperty(searchParams); err != nil {
		return nil, err
	}

	res, err := s.traverser.GetClass(ctx, principal, searchParams)
	if err != nil {
		return nil, err
	}

	return searchResultsToProto(res, before, searchParams)
}

func (s *Server) validateClassAndProperty(searchParams dto.GetParams) error {
	scheme := s.schemaManager.GetSchemaSkipAuth()
	class, err := schema.GetClassByName(scheme.Objects, searchParams.ClassName)
	if err != nil {
		return err
	}

	for _, prop := range searchParams.Properties {
		_, err := schema.GetPropertyByName(class, prop.Name)
		if err != nil {
			return err
		}
	}

	return nil
}
