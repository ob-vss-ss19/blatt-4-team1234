package handler

import (
	"context"
	"log"

	"github.com/ob-vss-ss19/blatt-4-team1234/commons"
	"github.com/ob-vss-ss19/blatt-4-team1234/showservice/proto/show"

	"github.com/ob-vss-ss19/blatt-4-team1234/hallservice/proto/hall"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HallHandler struct {
	Halls map[int64]hall.Hall
	NewID int64
}

func (handle *HallHandler) GetAllHalls(ctx context.Context, req *hall.GetAllHallsRequest,
	rsp *hall.GetAllHallsResponse) error {
	log.Printf("Received GetAllHallsRequest")
	protoHalls := make([]*hall.Hall, len(handle.Halls))
	i := 0
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
	log.Printf("Received GetHallRequest")
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
	log.Printf("Received RemoveHallRequest")
	if err := commons.CheckId(req.Id, "Hall"); err != nil {
		return err
	}
	_, found := handle.Halls[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Hall with the ID:%d does not Exist", req.Id)
	}
	if err := handle.RemoveShows(ctx, req.Id); err != nil {
		return status.Errorf(codes.Internal, "Error while calling ShowService, Error: "+err.Error())
	}
	delete(handle.Halls, req.Id)
	return nil
}

func (handle *HallHandler) AddHall(ctx context.Context, req *hall.AddHallRequest, rsp *hall.AddHallResponse) error {
	log.Printf("Received AddHallRequest")
	if req.Hall == nil {
		return status.Errorf(codes.InvalidArgument, "No Hall Submitted!")
	}
	if req.Hall.Rows <= 0 || req.Hall.Columns <= 0 || len(req.Hall.Name) < 1 {
		return status.Errorf(codes.InvalidArgument, "Please submit a name, the columns and rows of the hall!")
	}
	req.Hall.Id = handle.NewID
	handle.Halls[req.Hall.Id] = *req.Hall
	rsp.Id = handle.NewID
	handle.NewID++
	return nil
}

func (handle *HallHandler) RemoveShows(ctx context.Context, hallID int64) error {
	showRequest := show.RemoveShowsForHallRequest{HallId: hallID}
	showService := show.NewShowService(commons.GetShowServiceName(), nil)
	_, err := showService.RemoveShowsForHall(ctx, &showRequest)
	if err != nil {
		return err
	}
	return nil
}

func InitDB() *HallHandler {
	handler := HallHandler{Halls: make(map[int64]hall.Hall)}
	handler.Halls[1] = hall.Hall{Id: 1, Name: "Grosser-KinoSaal", Rows: 15, Columns: 15}
	handler.Halls[2] = hall.Hall{Id: 2, Name: "Kleiner-KinoSaal", Rows: 8, Columns: 10}
	handler.NewID = 3
	return &handler
}
