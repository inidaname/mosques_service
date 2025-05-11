package service

import (
	"context"

	db "github.com/inidaname/mosque/mosques-service/internal/db/models"
	"github.com/inidaname/mosque/mosques-service/internal/helpers"
	"github.com/inidaname/mosque/mosques-service/internal/types"
	"github.com/inidaname/mosque/protos"
	"github.com/jackc/pgx/v5/pgtype"
)

type MosqueService struct {
	*types.Application
}

func NewMosqueService(cfg *types.Application) *MosqueService {
	return &MosqueService{cfg}
}

func (s *MosqueService) CreateMosque(ctx context.Context, req *protos.CreateMosqueRequest) (*protos.CreateMosqueResponse, error) {
	eidTime, err := helpers.ConvertStringToTimestamp(req.EidTime.String())
	if err != nil {
		// Handle error
	}
	jummahTime, err := helpers.ConvertStringToTimestamp(req.JummahTime.String())
	if err != nil {
		// Handle error
	}

	Lat, err := helpers.ConvertToPgNumeric(req.Lat)
	if err != nil {
		// Handle error
	}

	Lng, err := helpers.ConvertToPgNumeric(req.Lng)
	if err != nil {
		// Handle error
	}
	mosque, err := s.Store.CreateMosque(ctx, db.CreateMosqueParams{
		Name:       req.Name,
		Address:    req.Address,
		EidTime:    eidTime,
		JummahTime: jummahTime,
		Lat:        Lat,
		Lng:        Lng,
	})

	return &protos.CreateMosqueResponse{
		Mosques: &protos.Mosque{
			Id:         mosque.ID.String(),
			Name:       mosque.Name,
			Address:    mosque.Address,
			EidTime:    req.EidTime,
			JummahTime: req.JummahTime,
			Lat:        req.Lat,
			Lng:        req.Lng,
		},
	}, nil
}

func (s *MosqueService) ListMosque(ctx context.Context, req *protos.ListMosquesRequest) (*protos.ListMosquesResponse, error) {
	dbMosques, err := s.Store.GetAllMosque(ctx)
	if err != nil {
		return nil, err
	}

	var mosque []*protos.Mosque

	for _, v := range dbMosques {
		mosque = append(mosque, &protos.Mosque{
			Id:         v.ID.String(),
			Name:       v.Name,
			Address:    v.Address,
			EidTime:    helpers.ConvertToTime(v.EidTime),
			JummahTime: helpers.ConvertToTime(v.JummahTime),
			Lat:        float64(v.Lat.Exp),
			Lng:        float64(v.Lng.Exp),
		})
	}

	return &protos.ListMosquesResponse{
		Mosques: mosque,
	}, nil
}

func (s *MosqueService) UpdateMpsque(ctx context.Context, req *protos.UpdateMosqueRequest) (*protos.UpdateMosqueResponse, error) {
	eidTime, err := helpers.ConvertStringToTimestamp(req.EidTime.String())
	if err != nil {
		// Handle error
	}
	jummahTime, err := helpers.ConvertStringToTimestamp(req.JummahTime.String())
	if err != nil {
		// Handle error
	}

	Lat, err := helpers.ConvertToPgNumeric(req.Lat)
	if err != nil {
		// Handle error
	}

	Lng, err := helpers.ConvertToPgNumeric(req.Lng)
	if err != nil {
		// Handle error
	}

	mosque, err := s.Store.UpdateMosque(ctx, db.UpdateMosqueParams{
		Name:       pgtype.Text{String: req.Name, Valid: true},
		Address:    pgtype.Text{String: req.Address, Valid: true},
		EidTime:    eidTime,
		JummahTime: jummahTime,
		Lat:        Lat,
		Lng:        Lng,
	})

	return &protos.UpdateMosqueResponse{Mosques: &protos.Mosque{
		Id:         mosque.ID.String(),
		Name:       mosque.Name,
		Address:    mosque.Address,
		EidTime:    req.EidTime,
		JummahTime: req.JummahTime,
		Lat:        req.Lat,
		Lng:        req.Lng,
	},
	}, nil
}
