package server

import (
	"context"
	"errors"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/wrrn/field-mask-poc/pkg/config"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errBadMask = errors.New("invalid mask")

type Server struct {
	confs map[string]config.Config
}

func (s *Server) configs() map[string]config.Config {
	if s.confs == nil {
		s.confs = make(map[string]config.Config)
	}
	return s.confs
}

// maskConfig returns a config where all of the fields in the mask are intentionally set.
func maskConfig(conf config.Config, mask field_mask.FieldMask) (config.Config, error) {

	// Start out with an empty config and "update" it with the mask. In other
	// words only the fields in the mask will be populated.
	return updateConfig(conf, config.Config{}, mask)
}

// List all of the configs.
func (s *Server) ListConfigs(_ context.Context, req *config.ListConfigsRequest) (*config.ListConfigsResponse, error) {

	respConfigs := make([]*config.Config, 0, len(s.configs()))
	for _, c := range s.configs() {

		masked, err := maskConfig(c, *req.FieldMask)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		respConfigs = append(respConfigs, &masked)
	}

	return &config.ListConfigsResponse{Configs: respConfigs}, nil
}

// Create a new config.
func (s *Server) CreateConfig(_ context.Context, req *config.CreateConfigRequest) (*config.CreateConfigResponse, error) {

	req.Config.Name = uuid.New().String()
	s.configs()[req.Config.Name] = *req.Config
	return &config.CreateConfigResponse{Config: req.Config}, nil
}

func updateConfig(updates, curr config.Config, updateMask field_mask.FieldMask) (config.Config, error) {

	publishers := map[string]func(*config.Config, *config.Config){
		"name":        func(updates, stored *config.Config) { stored.Name = updates.Name },
		"configname":  func(updates, stored *config.Config) { stored.ConfigName = updates.ConfigName },
		"config_name": func(updates, stored *config.Config) { stored.ConfigName = updates.ConfigName },
		"ip":          func(updates, stored *config.Config) { stored.Ip = updates.Ip },
		"tags":        func(updates, stored *config.Config) { stored.Tags = updates.Tags },
	}

	// Get a copy of the current config and only update the fields in the field mask
	updated := curr
	for _, path := range updateMask.Paths {
		publisher := publishers[strings.ToLower(path)]
		if publisher == nil {
			return config.Config{}, errBadMask
		}
		publisher(&updates, &updated)
	}

	return updated, nil

}

// Update a config.
func (s *Server) UpdateConfig(_ context.Context, req *config.UpdateConfigRequest) (*config.UpdateConfigResponse, error) {

	c, ok := s.configs()[req.Config.GetName()]
	if !ok {
		return nil, status.Error(codes.NotFound, "config not found")
	}

	updated, err := updateConfig(*req.Config, c, *req.UpdateMask)
	switch err {
	case errBadMask:
		return nil, status.Error(codes.InvalidArgument, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &config.UpdateConfigResponse{Config: &updated}, nil
}

// Delete a config.
func (s *Server) DeleteConfig(_ context.Context, req *config.DeleteConfigRequest) (*empty.Empty, error) {
	delete(s.configs(), req.Name)
	return &empty.Empty{}, nil
}
