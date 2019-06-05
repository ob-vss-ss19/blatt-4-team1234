package handler

import (
	"context"

	"github.com/ob-vss-ss19/blatt-4-team1234/movieservice/proto/movie"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MovieHandler struct {
	Movies map[int64]Movie
}

func (handle *MovieHandler) GetAllMovies(ctx context.Context, req *movie.GetAllMoviesRequest,
	rsp *movie.GetAllMoviesResponse) error {
	protoMovies := make([]*movie.Movie, len(handle.Movies))
	for i, m := range handle.Movies {
		protoMovies = append(protoMovies, &movie.Movie{Id: i, Title: m.Title, Fsk: m.Fsk})
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
	rsp.Movie = &movie.Movie{Id: req.Id, Title: m.Title, Fsk: m.Fsk}
	return nil
}

func (handle *MovieHandler) AddMovie(ctx context.Context, req *movie.AddMovieRequest,
	rsp *movie.AddMovieResponse) error {
	handle.Movies[int64(len(handle.Movies)+1)] = Movie{Title: req.Movie.Title, Fsk: req.Movie.Fsk}
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
	handle.Movies = make(map[int64]Movie)
	handle.Movies[0] = Movie{Title: "Der Schuh des Manitu", Fsk: 6}
	handle.Movies[1] = Movie{Title: "Traumschiff Surprise", Fsk: 6}
	handle.Movies[2] = Movie{Title: "Avengers: Endgame", Fsk: 12}
	handle.Movies[3] = Movie{Title: "Avengers: Infinity War", Fsk: 12}
}

type Movie struct {
	Title string
	Fsk   int64
}
