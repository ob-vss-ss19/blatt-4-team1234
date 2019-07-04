package handler

import (
	"context"
	"strings"
	"testing"

	"gotest.tools/assert"

	"github.com/ob-vss-ss19/blatt-4-team1234/movieservice/proto/movie"
)

func TestMovieHandler_GetAllMovies(t *testing.T) {
	movieHandler := InitDB()
	ctx := context.Background()
	req := movie.GetAllMoviesRequest{}
	resp := movie.GetAllMoviesResponse{}
	err := movieHandler.GetAllMovies(ctx, &req, &resp)
	if err != nil {
		t.Errorf("Error requesting all movies. ERR:" + err.Error())
	}
	assert.Assert(t, len(resp.Movies) == 4)
}

func TestMovieHandler_GetMovie(t *testing.T) {
	movieHandler := InitDB()
	ctx := context.Background()
	req := movie.GetMovieRequest{Id: 1}
	resp := movie.GetMovieResponse{}
	err := movieHandler.GetMovie(ctx, &req, &resp)
	if err != nil {
		t.Errorf("Error requesting a movie. ERR:" + err.Error())
	}
	assert.Assert(t, resp.Movie.Id == 1)
	assert.Assert(t, resp.Movie.Title == "Der Schuh des Manitu")
	assert.Assert(t, resp.Movie.Fsk == 6)
}

func TestMovieHandler_AddMovie(t *testing.T) {
	movieHandler := InitDB()
	ctx := context.Background()
	req := movie.AddMovieRequest{Movie: &movie.Movie{Id: -42, Title: "Ein-Film", Fsk: 21}}
	resp := movie.AddMovieResponse{}
	err := movieHandler.AddMovie(ctx, &req, &resp)
	if err != nil {
		t.Errorf("Error adding a movie. ERR:" + err.Error())
	}
	getReq := movie.GetMovieRequest{Id: 5}
	getResp := movie.GetMovieResponse{}
	err = movieHandler.GetMovie(ctx, &getReq, &getResp)
	if err != nil {
		t.Errorf("Error requesting a movie. ERR:" + err.Error())
	}
	assert.Assert(t, getResp.Movie.Id == 5)
	assert.Assert(t, getResp.Movie.Title == "Ein-Film")
	assert.Assert(t, getResp.Movie.Fsk == 21)
}

func TestMovieHandler_RemoveMovie(t *testing.T) {
	//Wont work without a mock
	movieHandler := InitDB()
	ctx := context.Background()
	req := movie.RemoveMovieRequest{Id: 1}
	resp := movie.RemoveMovieResponse{}
	err := movieHandler.RemoveMovie(ctx, &req, &resp)
	if err != nil {
		assert.Assert(t, strings.Contains(err.Error(), "code = Internal"))
	} else {
		t.Error("An Error was expected but not received!")
	}

}
