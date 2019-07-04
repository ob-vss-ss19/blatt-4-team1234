package handler

import (
	"context"
	"strings"
	"testing"

	"gotest.tools/assert"

	"github.com/ob-vss-ss19/blatt-4-team1234/hallservice/proto/hall"
)

func TestHallHandler_GetAllHalls(t *testing.T) {
	hallHandler := InitDB()
	ctx := context.Background()
	req := hall.GetAllHallsRequest{}
	resp := hall.GetAllHallsResponse{}
	err := hallHandler.GetAllHalls(ctx, &req, &resp)
	if err != nil {
		t.Errorf("Error requesting all halls. ERR:" + err.Error())
	}
	assert.Assert(t, len(resp.Halls) == 2)
}

func TestHallHandler_GetHall(t *testing.T) {
	hallHandler := InitDB()
	ctx := context.Background()
	req := hall.GetHallRequest{Id: 1}
	resp := hall.GetHallResponse{}
	err := hallHandler.GetHall(ctx, &req, &resp)
	if err != nil {
		t.Errorf("Error requesting a hall. ERR:" + err.Error())
	}
	assert.Assert(t, resp.Hall.Id == 1)
	assert.Assert(t, resp.Hall.Name == "Grosser-KinoSaal")
	assert.Assert(t, resp.Hall.Rows == 15)
	assert.Assert(t, resp.Hall.Columns == 15)
}

func TestHallHandler_AddHall(t *testing.T) {
	hallHandler := InitDB()
	ctx := context.Background()
	req := hall.AddHallRequest{Hall: &hall.Hall{Id: -42, Name: "Ein-KinoSaal", Rows: 5, Columns: 5}}
	resp := hall.AddHallResponse{}
	err := hallHandler.AddHall(ctx, &req, &resp)
	if err != nil {
		t.Errorf("Error adding a hall. ERR:" + err.Error())
	}
	getReq := hall.GetHallRequest{Id: 3}
	getResp := hall.GetHallResponse{}
	err = hallHandler.GetHall(ctx, &getReq, &getResp)
	if err != nil {
		t.Errorf("Error requesting a hall. ERR:" + err.Error())
	}
	assert.Assert(t, getResp.Hall.Id == 3)
	assert.Assert(t, getResp.Hall.Name == "Ein-KinoSaal")
	assert.Assert(t, getResp.Hall.Rows == 5)
	assert.Assert(t, getResp.Hall.Columns == 5)
}

func TestHallHandler_RemoveHall(t *testing.T) {
	//Wont work without a mock
	hallHandler := InitDB()
	ctx := context.Background()
	req := hall.RemoveHallRequest{Id: 1}
	resp := hall.RemoveHallResponse{}
	err := hallHandler.RemoveHall(ctx, &req, &resp)
	if err != nil {
		assert.Assert(t, strings.Contains(err.Error(), "code = Internal"))
	} else {
		t.Error("An Error was expected but not received!")
	}

}
