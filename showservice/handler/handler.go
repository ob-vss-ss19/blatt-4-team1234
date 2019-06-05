package handler

import (
	"context"
	"github.com/ob-vss-ss19/blatt-4-team1234/showservice/proto/show"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShowHandler struct{
	Shows map[int64]Show
}



type Show struct {
	MovieId int64
	HallId int64
	DateTime string
}

func (handle *ShowHandler) GetAllShows(ctx context.Context, req *show.GetAllShowsRequest, rsp *show.GetAllShowsResponse) error {
	var protoShows []*show.Show
	for i, s := range handle.Shows {
		protoShows = append(protoShows, &show.Show{Id:i,MovieId:s.MovieId,HallId:s.HallId,DateTime:s.DateTime})
	}
	rsp.Shows = protoShows
	return nil
}

func (handle *ShowHandler) GetShow(ctx context.Context,req *show.GetShowRequest,rsp *show.GetShowResponse) error {
	s,found := handle.Shows[req.Id]
	if !found{
		return status.Errorf(codes.NotFound,"The Show with the ID:%d does not Exist",req.Id)
	}
	rsp.Show = &show.Show{Id:req.Id,MovieId:s.MovieId,HallId:s.HallId,DateTime:s.DateTime}
	return nil
}

func (handle *ShowHandler) RemoveShow(ctx context.Context, req *show.RemoveShowRequest, rsp *show.RemoveShowResponse) error{
	_,found := handle.Shows[req.Id]
	if !found{
		return status.Errorf(codes.NotFound,"The Show with the ID:%d does not Exist",req.Id)
	}
	delete(handle.Shows,req.Id)
	return nil
}

func (handle *ShowHandler) AddShow(ctx context.Context,req *show.AddShowRequest,rsp *show.AddShowResponse) error {
	handle.Shows[int64(len(handle.Shows)+1)] = Show{MovieId:req.Show.MovieId,HallId:req.Show.HallId,DateTime:req.Show.DateTime}
	return nil
}

func (handle *ShowHandler) InitDB(){
	handle.Shows = make(map[int64]Show)
	handle.Shows[0]= Show{MovieId:0,HallId:0,DateTime:"2019-06-05_20:15"}
	handle.Shows[0]= Show{MovieId:1,HallId:1,DateTime:"2019-06-05_23:15"}
	handle.Shows[0]= Show{MovieId:2,HallId:1,DateTime:"2019-06-06_14:00"}
	handle.Shows[0]= Show{MovieId:3,HallId:0,DateTime:"2019-06-06_18:30"}
}
