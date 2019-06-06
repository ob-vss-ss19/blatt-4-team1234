package handler

import (
	"context"

	"github.com/ob-vss-ss19/blatt-4-team1234/reservationservice/proto/reservation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReservationHandler struct {
	Reservations map[int64]reservation.Reservation
}

func (handle *ReservationHandler) RemoveReservationsForShow(ctx context.Context,
	req *reservation.RemoveReservationForShowRequest, rsp *reservation.RemoveReservationForShowResponse) error {
	var deleteKeys []int64
	if req.ShowId == -1 {
		return status.Errorf(codes.InvalidArgument, "No ShowId was ")
	}
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
	if err := handle.SeatsAreFree(req); err != nil {
		return err
	}
	newId := int64(len(handle.Reservations) + 2) //+1 since map starts at 1, +1 since next item
	handle.Reservations[newId] = reservation.Reservation{Id: newId, ShowId: req.ShowId, Seats: req.Seats, UserId: -1}
	rsp.ReservationId = newId
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
	for _, r := range reservationsForShow {
		r := r
		for _, s := range r.Seats {
			if ContainsSeat(r.Seats, s) {
				return status.Errorf(codes.AlreadyExists, "A Reservation for Seat (Row: %d, Column:%d),"+
					" already Exists!", s.Row, s.Column)
			}
		}
	}
	return nil
}

func ContainsSeat(seats []*reservation.Seat, seat *reservation.Seat) bool {
	for _, r := range seats {
		if seat.Equal(r) {
			return false
		}
	}
	return true
}

func (handle *ReservationHandler) ActivateReservation(ctx context.Context, req *reservation.ActivateReservationRequest, rsp *reservation.ActivateReservationResponse) error {
	r, found := handle.Reservations[req.ReservationId]
	if !found {
		return status.Errorf(codes.NotFound, "The ReservationID (%d) was not found!", req.ReservationId)
	}
	r.UserId = req.UserId
	handle.Reservations[req.ReservationId] = r
	return nil
}

func (handle *ReservationHandler) GetReservationsForUser(ctx context.Context,
	req *reservation.GetReservationsForUserRequest, rsp *reservation.GetReservationsForUserResponse) error {
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
	r, found := handle.Reservations[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Reservation with the ID:%d does not Exist", req.Id)
	}
	rsp.Reservation = &r
	return nil
}

func (handle *ReservationHandler) RemoveReservation(ctx context.Context, req *reservation.RemoveReservationRequest,
	rsp *reservation.RemoveReservationResponse) error {
	_, found := handle.Reservations[req.Id]
	if !found {
		return status.Errorf(codes.NotFound, "The Reservation with the ID:%d does not Exist", req.Id)
	}
	delete(handle.Reservations, req.Id)
	return nil
}

func (handle *ReservationHandler) InitDB() {
	handle.Reservations = make(map[int64]reservation.Reservation)
	handle.Reservations[0] = reservation.Reservation{Id: 1, UserId: 1, ShowId: 1,
		Seats: []*reservation.Seat{{Row: 1, Column: 2}, {Row: 1, Column: 3}}}
	handle.Reservations[1] = reservation.Reservation{Id: 2, UserId: 2, ShowId: 2,
		Seats: []*reservation.Seat{{Row: 3, Column: 3}}}
	handle.Reservations[2] = reservation.Reservation{Id: 3, UserId: 3, ShowId: 3,
		Seats: []*reservation.Seat{{Row: 7, Column: 7}, {Row: 7, Column: 8}}}
	handle.Reservations[3] = reservation.Reservation{Id: 4, UserId: 4, ShowId: 4,
		Seats: []*reservation.Seat{{Row: 2, Column: 1}, {Row: 2, Column: 2}}}
}
