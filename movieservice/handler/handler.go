package handler

import (
	"context"

	"github.com/ob-vss-ss19/blatt-4-team1234/movieservice/proto/movie"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MovieHandler struct {
	Movies map[int64]movie.Movie
}

func (handle *MovieHandler) GetAllMovies(ctx context.Context, req *movie.GetAllMoviesRequest,
	rsp *movie.GetAllMoviesResponse) error {
	protoMovies := make([]*movie.Movie, len(handle.Movies))
	i := 0
	for _, m := range handle.Movies {
		protoMovies[i] = &m
		i++
	}
	rsp.Movies = protoMovies
	return nil
}

func (handle *MovieHandler) GetMovie(ctx context.Context, req *movie.GetMovieRequest,
	rsp *movie.GetMovieResponse) error {
	m, found := handle.Movies[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Movie with the ID:%d does not Exist", req.Id)
	}
	rsp.Movie = &m
	return nil
}

func (handle *MovieHandler) AddMovie(ctx context.Context, req *movie.AddMovieRequest,
	rsp *movie.AddMovieResponse) error {
	handle.Movies[int64(len(handle.Movies)+1)] = *req.Movie
	return nil
}

func (handle *MovieHandler) RemoveMovie(ctx context.Context, req *movie.RemoveMovieRequest,
	rsp *movie.RemoveMovieResponse) error {
	_, found := handle.Movies[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Hall with the ID:%d does not Exist", req.Id)
	}
	delete(handle.Movies, req.Id)
	return nil
}

func (handle *MovieHandler) InitDB() {
	handle.Movies = make(map[int64]movie.Movie)
	handle.Movies[0] = movie.Movie{Title: "Der Schuh des Manitu", Fsk: 6}
	handle.Movies[1] = movie.Movie{Title: "Traumschiff Surprise", Fsk: 6}
	handle.Movies[2] = movie.Movie{Title: "Avengers: Endgame", Fsk: 12}
	handle.Movies[3] = movie.Movie{Title: "Avengers: Infinity War", Fsk: 12}
}
