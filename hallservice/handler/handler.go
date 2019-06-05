package handler

import (
	"context"

	"github.com/ob-vss-ss19/blatt-4-team1234/hallservice/proto/hall"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HallHandler struct {
	Halls map[int64]hall.Hall
}

func (handle *HallHandler) GetAllHalls(ctx context.Context, req *hall.GetAllHallsRequest,
	rsp *hall.GetAllHallsResponse) error {
	protoHalls := make([]*hall.Hall, len(handle.Halls))
	for _, h := range handle.Halls {
		protoHalls = append(protoHalls, &h)
	}
	rsp.Halls = protoHalls
	return nil
}

func (handle *HallHandler) GetHall(ctx context.Context, req *hall.GetHallRequest,
	rsp *hall.GetHallResponse) error {
	h, found := handle.Halls[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Hall with the ID:%d does not Exist", req.Id)
	}
	rsp.Hall = &h
	return nil
}

func (handle *HallHandler) RemoveHall(ctx context.Context, req *hall.RemoveHallRequest,
	rsp *hall.RemoveHallResponse) error {
	_, found := handle.Halls[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Hall with the ID:%d does not Exist", req.Id)
	}
	delete(handle.Halls, req.Id)
	return nil
}

func (handle *HallHandler) AddHall(ctx context.Context, req *hall.AddHallRequest, rsp *hall.AddHallResponse) error {
	handle.Halls[int64(len(handle.Halls)+1)] = *req.Hall
	return nil
}

func (handle *HallHandler) InitDB() {
	handle.Halls = make(map[int64]hall.Hall)
	handle.Halls[0] = hall.Hall{Name: "Grosser-KinoSaal", Rows: 15, Columns: 15}
	handle.Halls[1] = hall.Hall{Name: "Kleiner-KinoSaal", Rows: 8, Columns: 10}
}
