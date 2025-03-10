package service

import (
	"context"
	"github.com/PiotrIzw/webstore-grcp/internal/middleware/authorizer"
	"github.com/PiotrIzw/webstore-grcp/internal/pb"
	"github.com/PiotrIzw/webstore-grcp/internal/preferences"
	"github.com/PiotrIzw/webstore-grcp/internal/repository"
)

type PreferencesService struct {
	repo       *repository.PreferencesRepository
	authorizer *authorizer.Authorizer
	pb.UnimplementedPreferencesServiceServer
}

func NewPreferencesService(repo *repository.PreferencesRepository, authorizer *authorizer.Authorizer) *PreferencesService {
	return &PreferencesService{repo: repo, authorizer: authorizer}
}

func (s *PreferencesService) UpdatePreferences(ctx context.Context, req *pb.UpdatePreferencesRequest) (*pb.UpdatePreferencesResponse, error) {
	if err := s.authorizer.Authorize(ctx, "preferences:write"); err != nil {
		return nil, err
	}

	pref := &preferences.Preferences{
		UserID:        req.UserId,
		Theme:         req.Theme,
		Notifications: req.Notifications,
		Locale:        req.Locale,
	}
	err := s.repo.UpdatePreferences(pref)
	if err != nil {
		return nil, err
	}
	return &pb.UpdatePreferencesResponse{Success: true}, nil
}

func (s *PreferencesService) GetPreferences(ctx context.Context, req *pb.GetPreferencesRequest) (*pb.GetPreferencesResponse, error) {
	if err := s.authorizer.Authorize(ctx, "preferences:read"); err != nil {
		return nil, err
	}

	pref, err := s.repo.GetPreferences(req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetPreferencesResponse{
		Theme:         pref.Theme,
		Notifications: pref.Notifications,
		Locale:        pref.Locale,
	}, nil
}
