package handler

import (
	"context"

	"github.com/ob-vss-ss19/blatt-4-team1234/showservice/proto/show"

	"github.com/ob-vss-ss19/blatt-4-team1234/commons"

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
	i := 1
	for _, h := range handle.Halls {
		h := h
		protoHalls[i] = &h
		i++
	}
	rsp.Halls = protoHalls
	return nil
}

func (handle *HallHandler) GetHall(ctx context.Context, req *hall.GetHallRequest,
	rsp *hall.GetHallResponse) error {
	if err := commons.CheckId(req.Id, "Hall"); err != nil {
		return err
	}
	h, found := handle.Halls[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Hall with the ID:%d does not Exist", req.Id)
	}
	rsp.Hall = &h
	return nil
}

func (handle *HallHandler) RemoveHall(ctx context.Context, req *hall.RemoveHallRequest,
	rsp *hall.RemoveHallResponse) error {
	if err := commons.CheckId(req.Id, "Hall"); err != nil {
		return err
	}
	_, found := handle.Halls[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Hall with the ID:%d does not Exist", req.Id)
	}
	if err := handle.RemoveReservations(ctx, req.Id); err != nil {
		return err
	}
	delete(handle.Halls, req.Id)
	return nil
}

func (handle *HallHandler) AddHall(ctx context.Context, req *hall.AddHallRequest, rsp *hall.AddHallResponse) error {
	if req.Hall.Rows <= 0 || req.Hall.Columns <= 0 || len(req.Hall.Name) < 1 {
		return status.Errorf(codes.InvalidArgument, "Please submit a name, the columns and rows of the hall!")
	}
	handle.Halls[int64(len(handle.Halls)+2)] = *req.Hall
	return nil
}

func (handle *HallHandler) RemoveReservations(ctx context.Context, hallId int64) error {
	showRequest := show.RemoveShowsForHallRequest{HallId: hallId}
	showService := show.NewShowService("go.micro.srv.showService", nil)
	_, err := showService.RemoveShowsForHall(ctx, &showRequest)
	if err != nil {
		return err
	}
	return nil
}

func (handle *HallHandler) InitDB() {
	handle.Halls = make(map[int64]hall.Hall)
	handle.Halls[1] = hall.Hall{Id: 1, Name: "Grosser-KinoSaal", Rows: 15, Columns: 15}
	handle.Halls[2] = hall.Hall{Id: 2, Name: "Kleiner-KinoSaal", Rows: 8, Columns: 10}
}
