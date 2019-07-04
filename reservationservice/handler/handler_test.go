package handler

import (
	"context"
	"strings"
	"testing"

	"gotest.tools/assert"

	"github.com/ob-vss-ss19/blatt-4-team1234/reservationservice/proto/reservation"
)

func TestReservationHandler_GetAllReservations(t *testing.T) {
	reservationHandler := InitDB()
	ctx := context.Background()
	req := reservation.GetAllReservationsRequest{}
	resp := reservation.GetAllReservationsResponse{}
	err := reservationHandler.GetAllReservations(ctx, &req, &resp)
	if err != nil {
		t.Errorf("Error requesting all reservations. ERR:" + err.Error())
	}
	assert.Assert(t, len(resp.Reservations) == 4)
}

func TestReservationHandler_GetReservation(t *testing.T) {
	reservationHandler := InitDB()
	ctx := context.Background()
	req := reservation.GetReservationRequest{Id: 1}
	resp := reservation.GetReservationResponse{}
	err := reservationHandler.GetReservation(ctx, &req, &resp)
	if err != nil {
		t.Errorf("Error requesting a reservation. ERR:" + err.Error())
	}
	assert.Assert(t, resp.Reservation.Id == 1)
	assert.Assert(t, resp.Reservation.UserId == 1)
	assert.Assert(t, resp.Reservation.ShowId == 1)
	assert.Assert(t, len(resp.Reservation.Seats) == 2)
}

func TestReservationHandler_RequestReservation(t *testing.T) {
	reservationHandler := InitDB()
	ctx := context.Background()
	seats := []*reservation.Seat{{Row: 1, Column: 1}}
	req := reservation.RequestReservationRequest{ShowId: 1, UserId: 1, Seats: seats}
	resp := reservation.RequestReservationResponse{}
	err := reservationHandler.RequestReservation(ctx, &req, &resp)
	if err != nil {
		assert.Assert(t, strings.Contains(err.Error(), "code = Internal"))
	} else {
		t.Errorf("An Error was Expected here!")
	}

}

func TestReservationHandler_RemoveReservation(t *testing.T) {
	//Wont work without a mock
	reservationHandler := InitDB()
	ctx := context.Background()
	req := reservation.RemoveReservationRequest{Id: 1}
	resp := reservation.RemoveReservationResponse{}
	err := reservationHandler.RemoveReservation(ctx, &req, &resp)
	if err != nil {
		t.Error("No Error was expected here!")
	}
	newReq := reservation.GetReservationRequest{Id: 1}
	newResp := reservation.GetReservationResponse{}
	err = reservationHandler.GetReservation(ctx, &newReq, &newResp)
	if err != nil {
		assert.Assert(t, strings.Contains(err.Error(), "code = NotFound"))
	} else {
		t.Error("An Error was expected here!")
	}
}
