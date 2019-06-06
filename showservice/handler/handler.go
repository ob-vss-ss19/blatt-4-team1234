package handler

import (
	"context"

	"github.com/ob-vss-ss19/blatt-4-team1234/hallservice/proto/hall"

	"github.com/ob-vss-ss19/blatt-4-team1234/movieservice/proto/movie"

	"github.com/ob-vss-ss19/blatt-4-team1234/showservice/proto/show"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShowHandler struct {
	Shows map[int64]show.Show
}

func (handle *ShowHandler) GetAllShows(ctx context.Context, req *show.GetAllShowsRequest,
	rsp *show.GetAllShowsResponse) error {
	protoShows := make([]*show.Show, len(handle.Shows))
	i := 1
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
	s, found := handle.Shows[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Show with the ID:%d does not Exist", req.Id)
	}
	rsp.Show = &s
	return nil
}

func (handle *ShowHandler) RemoveShow(ctx context.Context, req *show.RemoveShowRequest,
	rsp *show.RemoveShowResponse) error {
	_, found := handle.Shows[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Show with the ID:%d does not Exist", req.Id)
	}
	delete(handle.Shows, req.Id)
	return nil
}

func (handle *ShowHandler) AddShow(ctx context.Context, req *show.AddShowRequest, rsp *show.AddShowResponse) error {
	movieRequest := movie.GetMovieRequest{Id: req.Show.MovieId}
	movieService := movie.NewMovieService("go.micro.srv.movieservice", nil)
	_, err := movieService.GetMovie(ctx, &movieRequest)
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "No Movie with the Id (%d) exists", req.Show.MovieId)
	}
	hallRequest := hall.GetHallRequest{Id: req.Show.HallId}
	hallService := hall.NewHallService("go.micro.srv.hallservice", nil)
	_, err = hallService.GetHall(ctx, &hallRequest)
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "No Hall with the Id (%d) exists", req.Show.HallId)
	}
	handle.Shows[int64(len(handle.Shows)+2)] = *req.Show
	return nil
}

func (handle *ShowHandler) InitDB() {
	handle.Shows = make(map[int64]show.Show)
	handle.Shows[1] = show.Show{MovieId: 1, HallId: 1, DateTime: "2019-06-05_20:15"}
	handle.Shows[2] = show.Show{MovieId: 2, HallId: 2, DateTime: "2019-06-05_23:15"}
	handle.Shows[3] = show.Show{MovieId: 3, HallId: 2, DateTime: "2019-06-06_14:00"}
	handle.Shows[4] = show.Show{MovieId: 4, HallId: 1, DateTime: "2019-06-06_18:30"}
}
