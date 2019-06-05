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

func (handle *ReservationHandler) GetAllReservations(ctx context.Context, req *reservation.GetAllReservationsRequest,
	rsp *reservation.GetAllReservationsResponse) error {
	protoReservations := make([]*reservation.Reservation, len(handle.Reservations))
	for _, r := range handle.Reservations {
		protoReservations = append(protoReservations, &r)
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

func (handle *ReservationHandler) AddReservation(ctx context.Context, req *reservation.AddReservationRequest,
	rsp *reservation.AddReservationResponse) error {
	handle.Reservations[int64(len(handle.Reservations)+1)] = *req.Reservation
	return nil
}

func (handle *ReservationHandler) InitDB() {
	handle.Reservations = make(map[int64]reservation.Reservation)
	handle.Reservations[0] = reservation.Reservation{UserId: 0, ShowId: 0,
		Seats: []*reservation.Seat{{Row: 1, Column: 2}, {Row: 1, Column: 3}}}
	handle.Reservations[1] = reservation.Reservation{UserId: 1, ShowId: 1,
		Seats: []*reservation.Seat{{Row: 3, Column: 3}}}
	handle.Reservations[2] = reservation.Reservation{UserId: 2, ShowId: 2,
		Seats: []*reservation.Seat{{Row: 7, Column: 7}, {Row: 7, Column: 8}}}
	handle.Reservations[3] = reservation.Reservation{UserId: 3, ShowId: 3,
		Seats: []*reservation.Seat{{Row: 2, Column: 1}, {Row: 2, Column: 2}}}
}
