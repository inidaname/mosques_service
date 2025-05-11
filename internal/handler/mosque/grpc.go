package handler

import (
	"context"

	"github.com/inidaname/mosque/mosques-service/internal/service"
	"github.com/inidaname/mosque/protos"
	"google.golang.org/grpc"
)

type MosqueGrpcHandler struct {
	mosqueService service.MosqueService
	protos.UnimplementedMosqueServiceServer
}

func NewGrpcMosqueService(grpc *grpc.Server, mosqueService service.MosqueService) {
	grpcHandler := &MosqueGrpcHandler{
		mosqueService: mosqueService,
	}

	protos.RegisterMosqueServiceServer(grpc, grpcHandler)
}

func (h *MosqueGrpcHandler) CreateMosque(ctx context.Context, req *protos.CreateMosqueRequest) (*protos.CreateMosqueResponse, error) {
	resp, err := h.mosqueService.CreateMosque(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *MosqueGrpcHandler) ListMosque(ctx context.Context, req *protos.ListMosquesRequest) (*protos.ListMosquesResponse, error) {
	resp, err := h.mosqueService.ListMosque(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *MosqueGrpcHandler) UpdateMosque(ctx context.Context, req *protos.UpdateMosqueRequest) (*protos.UpdateMosqueResponse, error) {
	resp, err := h.mosqueService.UpdateMpsque(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
