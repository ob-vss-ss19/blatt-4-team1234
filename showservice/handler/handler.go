package handler

import (
	"context"
	"log"

	"github.com/ob-vss-ss19/blatt-4-team1234/reservationservice/proto/reservation"

	"github.com/ob-vss-ss19/blatt-4-team1234/commons"

	"github.com/ob-vss-ss19/blatt-4-team1234/hallservice/proto/hall"

	"github.com/ob-vss-ss19/blatt-4-team1234/movieservice/proto/movie"

	"github.com/ob-vss-ss19/blatt-4-team1234/showservice/proto/show"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShowHandler struct {
	Shows map[int64]show.Show
	NewID int64
}

func (handle *ShowHandler) RemoveShowsForHall(ctx context.Context, req *show.RemoveShowsForHallRequest,
	rsp *show.RemoveShowsForHallResponse) error {
	log.Printf("Received RemoveShowsForHallRequest")
	if err := commons.CheckId(req.HallId, "Hall"); err != nil {
		return err
	}
	var deleteKeys []int64
	for i, r := range handle.Shows {
		r := r
		i := i
		if r.HallId == req.HallId {
			deleteKeys = append(deleteKeys, i)
		}
	}
	for _, i := range deleteKeys {
		if err := handle.RemoveShow(ctx, &show.RemoveShowRequest{Id: i}, &show.RemoveShowResponse{}); err != nil {
			return status.Errorf(codes.Internal, "Error calling ReservationService, Error: "+err.Error())
		}
	}
	return nil
}

func (handle *ShowHandler) RemoveShowsForMovie(ctx context.Context, req *show.RemoveShowsForMovieRequest,
	rsp *show.RemoveShowsForMovieResponse) error {
	log.Printf("Received RemoveShowsForMovieRequest")
	if err := commons.CheckId(req.MovieId, "Movie"); err != nil {
		return err
	}
	var deleteKeys []int64
	for i, r := range handle.Shows {
		r := r
		i := i
		if r.MovieId == req.MovieId {
			deleteKeys = append(deleteKeys, i)
		}
	}
	for _, i := range deleteKeys {
		if err := handle.RemoveShow(ctx, &show.RemoveShowRequest{Id: i}, &show.RemoveShowResponse{}); err != nil {
			return status.Errorf(codes.Internal, "Error calling ReservationService, Error: "+err.Error())
		}
	}
	return nil
}

func (handle *ShowHandler) GetAllShows(ctx context.Context, req *show.GetAllShowsRequest,
	rsp *show.GetAllShowsResponse) error {
	log.Printf("Received GetAllShowsRequest")
	protoShows := make([]*show.Show, len(handle.Shows))
	i := 0
	for _, s := range handle.Shows {
		s := s
		protoShows[i] = &s
		i++
	}
	rsp.Shows = protoShows
	return nil
}

func (handle *ShowHandler) GetShow(ctx context.Context, req *show.GetShowRequest,
	rsp *show.GetShowResponse) error {
	log.Printf("Received GetShowRequest")
	if err := commons.CheckId(req.Id, "RequestId"); err != nil {
		return err
	}
	s, found := handle.Shows[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Show with the ID:%d does not Exist", req.Id)
	}
	rsp.Show = &s
	return nil
}

func (handle *ShowHandler) RemoveShow(ctx context.Context, req *show.RemoveShowRequest,
	rsp *show.RemoveShowResponse) error {
	log.Printf("Received RemoveShowRequest")
	if err := commons.CheckId(req.Id, "Show"); err != nil {
		return err
	}
	_, found := handle.Shows[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Show with the ID:%d does not Exist", req.Id)
	}
	if err := handle.RemoveReservations(ctx, req.Id); err != nil {
		return err
	}
	delete(handle.Shows, req.Id)
	return nil
}

func (handle *ShowHandler) RemoveReservations(ctx context.Context, showID int64) error {
	reservationRequest := reservation.RemoveReservationForShowRequest{ShowId: showID}
	reservationService := reservation.NewReservationService(commons.GetReservationServiceName(), nil)
	_, err := reservationService.RemoveReservationsForShow(ctx, &reservationRequest)
	if err != nil {
		return err
	}
	return nil
}

func (handle *ShowHandler) AddShow(ctx context.Context, req *show.AddShowRequest, rsp *show.AddShowResponse) error {
	if req.Show == nil {
		return status.Errorf(codes.InvalidArgument, "No Show was Provided!")
	}
	log.Printf("Received AddShowRequest")
	if err := commons.CheckId(req.Show.MovieId, "Movie"); err != nil {
		return err
	}
	if err := commons.CheckId(req.Show.HallId, "Hall"); err != nil {
		return err
	}
	if len(req.Show.DateTime) < 1 {
		return status.Errorf(codes.InvalidArgument, "No DateTime was Specified")
	}
	movieRequest := movie.GetMovieRequest{Id: req.Show.MovieId}
	movieService := movie.NewMovieService(commons.GetMovieServiceName(), nil)
	_, err := movieService.GetMovie(ctx, &movieRequest)
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "No Movie with the Id (%d) exists", req.Show.MovieId)
	}
	hallRequest := hall.GetHallRequest{Id: req.Show.HallId}
	hallService := hall.NewHallService(commons.GetHallServiceName(), nil)
	_, err = hallService.GetHall(ctx, &hallRequest)
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "No Hall with the Id (%d) exists", req.Show.HallId)
	}
	req.Show.Id = handle.NewID
	handle.Shows[req.Show.Id] = *req.Show
	rsp.Id = handle.NewID
	handle.NewID++
	return nil
}

func InitDB() *ShowHandler {
	handler := ShowHandler{Shows: make(map[int64]show.Show)}
	handler.Shows[1] = show.Show{MovieId: 1, HallId: 1, DateTime: "2019-06-05_20:15"}
	handler.Shows[2] = show.Show{MovieId: 2, HallId: 2, DateTime: "2019-06-05_23:15"}
	handler.Shows[3] = show.Show{MovieId: 3, HallId: 2, DateTime: "2019-06-06_14:00"}
	handler.Shows[4] = show.Show{MovieId: 4, HallId: 1, DateTime: "2019-06-06_18:30"}
	handler.NewID = 5
	return &handler
}
