package handler

import (
	"context"
	"log"

	"github.com/ob-vss-ss19/blatt-4-team1234/commons"
	"github.com/ob-vss-ss19/blatt-4-team1234/hallservice/proto/hall"
	"github.com/ob-vss-ss19/blatt-4-team1234/reservationservice/proto/reservation"
	"github.com/ob-vss-ss19/blatt-4-team1234/showservice/proto/show"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReservationConflict struct {
	Conflicting []int64
}

type ReservationHandler struct {
	Reservations         map[int64]reservation.Reservation
	ReservationConflicts map[int64]ReservationConflict
	NewID                int64
}

func (handle *ReservationHandler) RemoveReservationsForShow(ctx context.Context,
	req *reservation.RemoveReservationForShowRequest, rsp *reservation.RemoveReservationForShowResponse) error {
	log.Printf("Received RemoveReservationsForShowRequest")
	if err := commons.CheckId(req.ShowId, "Show"); err != nil {
		return err
	}
	var deleteKeys []int64
	for i, r := range handle.Reservations {
		r := r
		i := i
		if r.ShowId == req.ShowId {
			deleteKeys = append(deleteKeys, i)
		}
	}
	for _, i := range deleteKeys {
		delete(handle.Reservations, i)
	}
	return nil
}

func (handle *ReservationHandler) RequestReservation(ctx context.Context, req *reservation.RequestReservationRequest,
	rsp *reservation.RequestReservationResponse) error {
	log.Printf("Received RequestReservationRequest")
	if err := commons.CheckId(req.ShowId, "Show"); err != nil {
		return err
	}
	if len(req.Seats) < 1 {
		return status.Errorf(codes.InvalidArgument, "At least One Seat needs to be provided ")
	}
	if err := handle.CheckSeatValidity(ctx, req); err != nil {
		return err
	}
	if err := handle.SeatsAreFree(req); err != nil {
		return err
	} //+1 since map starts at 1, +1 since next item
	handle.Reservations[handle.NewID] = reservation.Reservation{Id: handle.NewID, ShowId: req.ShowId,
		Seats: req.Seats, UserId: req.UserId, Active: false}
	rsp.ReservationId = handle.NewID
	handle.NewID++
	return nil
}

func (handle *ReservationHandler) SeatsAreFree(req *reservation.RequestReservationRequest) error {
	var reservationsForShow []*reservation.Reservation
	for _, r := range handle.Reservations {
		r := r
		if r.ShowId == req.ShowId {
			reservationsForShow = append(reservationsForShow, &r)
		}
	}
	conflicts := make([]int64, 0)
	for _, r := range reservationsForShow {
		r := r
		for _, s := range r.Seats {
			if handle.ContainsSeat(r.Seats, s) {
				if r.Active {
					return status.Errorf(codes.AlreadyExists, "A Reservation for Seat (Row: %d, Column:%d),"+
						" already Exists!", s.Row, s.Column)
				}
				conflicts = append(conflicts, r.Id)
			}
		}
	}
	if len(conflicts) > 0 {
		handle.ReservationConflicts[handle.NewID] = ReservationConflict{Conflicting: conflicts}
	}
	return nil
}

func (handle *ReservationHandler) ContainsSeat(seats []*reservation.Seat, seat *reservation.Seat) bool {
	for _, r := range seats {
		if seat.Equal(r) {
			return false
		}
	}
	return true
}

func (handle *ReservationHandler) ActivateReservation(ctx context.Context, req *reservation.ActivateReservationRequest,
	rsp *reservation.ActivateReservationResponse) error {
	log.Printf("ReceivedActivateReservationRequest")
	if err := commons.CheckId(req.ReservationId, "Reservation"); err != nil {
		return err
	}
	if err := commons.CheckId(req.UserId, "User"); err != nil {
		return err
	}
	r, found := handle.Reservations[req.ReservationId]
	if !found {
		return status.Errorf(codes.NotFound, "The ReservationID (%d) was not found!", req.ReservationId)
	}
	if r.UserId != req.UserId {
		return status.Errorf(codes.FailedPrecondition, "The userid does not match the reservation's userid")
	}
	if _, found := handle.ReservationConflicts[r.Id]; found {
		for conflict := range (handle.ReservationConflicts[r.Id]).Conflicting {
			delete(handle.Reservations, int64(conflict))
		}
	}
	r.UserId = req.UserId
	handle.Reservations[req.ReservationId] = r
	return nil
}

func (handle *ReservationHandler) GetReservationsForUser(ctx context.Context,
	req *reservation.GetReservationsForUserRequest, rsp *reservation.GetReservationsForUserResponse) error {
	log.Printf("Received GetReservationsForUserRequest")
	if err := commons.CheckId(req.UserId, "User"); err != nil {
		return err
	}
	var userReservations []*reservation.Reservation
	for _, r := range handle.Reservations {
		r := r
		if req.UserId == r.Id {
			userReservations = append(userReservations, &r)
		}
	}
	rsp.Reservations = userReservations
	return nil
}

func (handle *ReservationHandler) GetAllReservations(ctx context.Context, req *reservation.GetAllReservationsRequest,
	rsp *reservation.GetAllReservationsResponse) error {
	log.Printf("Received GetAllReservationsRequest")
	protoReservations := make([]*reservation.Reservation, len(handle.Reservations))
	i := 0
	for _, r := range handle.Reservations {
		r := r
		protoReservations[i] = &r
		i++
	}
	rsp.Reservations = protoReservations
	return nil
}

func (handle *ReservationHandler) GetReservation(ctx context.Context, req *reservation.GetReservationRequest,
	rsp *reservation.GetReservationResponse) error {
	log.Printf("Received GetReservationRequest")
	if err := commons.CheckId(req.Id, "Reservation"); err != nil {
		return err
	}
	r, found := handle.Reservations[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Reservation with the ID:%d does not Exist", req.Id)
	}
	if !r.Active {
		return status.Errorf(codes.FailedPrecondition,"The Reservation is not yet Active!")
	}
	rsp.Reservation = &r
	return nil
}

func (handle *ReservationHandler) RemoveReservation(ctx context.Context, req *reservation.RemoveReservationRequest,
	rsp *reservation.RemoveReservationResponse) error {
	log.Printf("Received RemoveReservationRequest")
	if err := commons.CheckId(req.Id, "Reservation"); err != nil {
		return err
	}
	_, found := handle.Reservations[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Reservation with the ID:%d does not Exist", req.Id)
	}
	delete(handle.Reservations, req.Id)
	return nil
}

func InitDB() *ReservationHandler {
	handler := ReservationHandler{
		ReservationConflicts: make(map[int64]ReservationConflict),
		Reservations:         make(map[int64]reservation.Reservation),
		NewID:                5,
	}
	handler.Reservations[1] = reservation.Reservation{Id: 1, UserId: 1, ShowId: 1,
		Seats: []*reservation.Seat{{Row: 1, Column: 2}, {Row: 1, Column: 3}},Active:true}
	handler.Reservations[2] = reservation.Reservation{Id: 2, UserId: 2, ShowId: 2,
		Seats: []*reservation.Seat{{Row: 3, Column: 3}},Active:true}
	handler.Reservations[3] = reservation.Reservation{Id: 3, UserId: 3, ShowId: 3,
		Seats: []*reservation.Seat{{Row: 7, Column: 7}, {Row: 7, Column: 8}},Active:true}
	handler.Reservations[4] = reservation.Reservation{Id: 4, UserId: 4, ShowId: 4,
		Seats: []*reservation.Seat{{Row: 2, Column: 1}, {Row: 2, Column: 2}},Active:true}
	return &handler
}

func (handle *ReservationHandler) CheckSeatValidity(ctx context.Context,
	req *reservation.RequestReservationRequest) error {
	showRequest := show.GetShowRequest{Id: req.ShowId}
	showService := show.NewShowService(commons.GetShowServiceName(), nil)
	showResponse, err := showService.GetShow(ctx, &showRequest)
	if err != nil {
		return status.Errorf(codes.Internal, "An Error occurred while requesting the showService! Error:"+err.Error())
	}
	hallRequest := hall.GetHallRequest{Id: showResponse.Show.HallId}
	hallService := hall.NewHallService(commons.GetHallServiceName(), nil)
	hallResponse, err := hallService.GetHall(ctx, &hallRequest)
	if err != nil {
		return status.Errorf(codes.Internal, "An Error occurred while requesting the hallService! Error:"+err.Error())
	}
	for _, seat := range req.Seats {
		if seat.Column > hallResponse.Hall.Columns {
			return status.Errorf(codes.InvalidArgument, "A seats column exceeds the maximum columns of the hall")
		}
		if seat.Row > hallResponse.Hall.Rows {
			return status.Errorf(codes.InvalidArgument, "A seats row exceeds the maximum columns of the hall")
		}
	}
	return nil
}
