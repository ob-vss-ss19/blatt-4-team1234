package handler

import (
	"context"
	"log"

	"github.com/ob-vss-ss19/blatt-4-team1234/showservice/proto/show"

	"github.com/ob-vss-ss19/blatt-4-team1234/commons"

	"github.com/ob-vss-ss19/blatt-4-team1234/movieservice/proto/movie"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MovieHandler struct {
	Movies map[int64]movie.Movie
	NewID  int64
}

func (handle *MovieHandler) GetAllMovies(ctx context.Context, req *movie.GetAllMoviesRequest,
	rsp *movie.GetAllMoviesResponse) error {
	log.Printf("Received GetAllMoviesRequest")
	protoMovies := make([]*movie.Movie, len(handle.Movies))
	i := 0
	for _, m := range handle.Movies {
		m := m
		protoMovies[i] = &m
		i++
	}
	rsp.Movies = protoMovies
	return nil
}

func (handle *MovieHandler) GetMovie(ctx context.Context, req *movie.GetMovieRequest,
	rsp *movie.GetMovieResponse) error {
	log.Printf("Received GetMovieRequest")
	if err := commons.CheckId(req.Id, "Movie"); err != nil {
		return err
	}
	m, found := handle.Movies[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Movie with the ID:%d does not Exist", req.Id)
	}
	rsp.Movie = &m
	return nil
}

func (handle *MovieHandler) AddMovie(ctx context.Context, req *movie.AddMovieRequest,
	rsp *movie.AddMovieResponse) error {
	log.Printf("Received AddMovieRequest")
	if req.Movie == nil {
		return status.Errorf(codes.InvalidArgument, "No Movie Submitted!")
	}
	if len(req.Movie.Title) < 1 || req.Movie.Fsk < 1 {
		return status.Errorf(codes.InvalidArgument, "Please Submit a Title and a FSK-Rating!")
	}
	req.Movie.Id = handle.NewID
	handle.Movies[req.Movie.Id] = *req.Movie
	rsp.Id = handle.NewID
	handle.NewID++
	return nil
}

func (handle *MovieHandler) RemoveMovie(ctx context.Context, req *movie.RemoveMovieRequest,
	rsp *movie.RemoveMovieResponse) error {
	log.Printf("Received RemoveMovieRequest")
	if err := commons.CheckId(req.Id, "Movie"); err != nil {
		return err
	}
	_, found := handle.Movies[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Hall with the ID:%d does not Exist", req.Id)
	}
	if err := handle.RemoveShows(ctx, req.Id); err != nil {
		return err
	}
	delete(handle.Movies, req.Id)
	return nil
}

func (handle *MovieHandler) RemoveShows(ctx context.Context, movieID int64) error {
	showRequest := show.RemoveShowsForMovieRequest{MovieId: movieID}
	showService := show.NewShowService(commons.GetShowServiceName(), nil)
	_, err := showService.RemoveShowsForMovie(ctx, &showRequest)
	if err != nil {
		return status.Errorf(codes.Internal, "Error while calling ShowService, Error: "+err.Error())
	}
	return nil
}

func InitDB() *MovieHandler {
	handler := MovieHandler{Movies: make(map[int64]movie.Movie)}
	handler.Movies[1] = movie.Movie{Id: 1, Title: "Der Schuh des Manitu", Fsk: 6}
	handler.Movies[2] = movie.Movie{Id: 2, Title: "Traumschiff Surprise", Fsk: 6}
	handler.Movies[3] = movie.Movie{Id: 3, Title: "Avengers: Endgame", Fsk: 12}
	handler.Movies[4] = movie.Movie{Id: 4, Title: "Avengers: Infinity War", Fsk: 12}
	handler.NewID = 5
	return &handler
}
