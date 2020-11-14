package service

import (
	"testing"

	"github.com/diegoclair/best-route-travel/domain/contract"
)

func getService() *Service {
	svc := new(Service)
	return svc
}

func Test_validateInput(t *testing.T) {
	type fields struct {
		svc           *Service
		travelService contract.TravelService
	}
	tests := []struct {
		name   string
		fields fields
		input  string
		want   bool
	}{
		{
			name:  "Validated with success",
			input: "GRU-CDG",
			want:  true,
		},
		{
			name:  "Validated Failed - where to invalid",
			input: "GRU-",
			want:  false,
		},
		{
			name:  "Validated Failed - where from invalid",
			input: "-CDG",
			want:  false,
		},
		{
			name:  "Validated Failed - invalid input",
			input: "HEGR",
			want:  false,
		},
		{
			name:  "Validated Failed - More than 2 inputs",
			input: "GRU-CDG-BRL",
			want:  false,
		},
		{
			name:  "Validated Failed - blank input",
			input: "-",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &commandLine{
				svc: tt.fields.svc,
			}
			if got := cli.validateInput(tt.input); got != tt.want {
				t.Errorf("commandLine.validateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRoutes(t *testing.T) {
	svc := &Service{}
	cli := &commandLine{
		svc: svc,
	}

	whereFrom, whereTo := cli.getRoutes("GRU-CDG")
	if whereFrom != "GRU" {
		t.Errorf("getRoutes failed, expected %s, got %v", "GRU", whereFrom)
	}
	if whereTo != "CDG" {
		t.Errorf("getRoutes failed, expected %s, got %v", "CDG", whereTo)
	}
}
