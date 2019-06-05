package handler

import (
	"context"
	"github.com/ob-vss-ss19/blatt-4-team1234/hallservice/proto/hall"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HallHandler struct{
	Halls map[int64]Hall
}



type Hall struct {
	Name string
	Rows int64
	Columns int64
}

func (e *HallHandler) GetAllHalls(ctx context.Context, req *hall.GetAllHallsRequest, rsp *hall.GetAllHallsResponse) error {
	var protoHalls []*hall.Hall
	for i,h := range e.Halls {
		protoHalls = append(protoHalls, &hall.Hall{Id:i,Name:h.Name,Rows:h.Rows,Columns:h.Columns})
	}
	rsp.Halls = protoHalls
	return nil
}

func (e *HallHandler) GetHall(ctx context.Context,req *hall.GetHallRequest,rsp *hall.GetHallResponse) error {
	h,found := e.Halls[req.Id]
	if !found{
		return status.Errorf(codes.NotFound,"The Hall with the ID:%d does not Exist",req.Id)
	}
	rsp.Hall = &hall.Hall{Id:req.Id,Name:h.Name,Rows:h.Rows,Columns:h.Columns,}
	return nil
}

func (e *HallHandler) RemoveHall(ctx context.Context, req *hall.RemoveHallRequest, rsp *hall.RemoveHallResponse) error{
	_,found := e.Halls[req.Id]
	if !found{
		return status.Errorf(codes.NotFound,"The Hall with the ID:%d does not Exist",req.Id)
	}
	delete(e.Halls,req.Id)
	return nil
}

func (e *HallHandler) AddHall(ctx context.Context,req *hall.AddHallRequest,rsp *hall.AddHallResponse) error {
	e.Halls[int64(len(e.Halls)+1)] = Hall{Name:req.Hall.Name,Rows:req.Hall.Rows,Columns:req.Hall.Rows}
	return nil
}

func (e *HallHandler) InitDB(){
	e.Halls = make(map[int64]Hall)
	e.Halls[0]= Hall{Name:"Grosser-KinoSaal",Rows:15,Columns:15}
	e.Halls[1]= Hall{Name:"Kleiner-KinoSaal",Rows:8,Columns:10}
}
