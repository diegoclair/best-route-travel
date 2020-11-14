package service

import (
	"reflect"
	"testing"

	"github.com/RyanCarrier/dijkstra"
	"github.com/diegoclair/best-route-travel/domain/entity"
)

func Test_parseRowsToStruct(t *testing.T) {
	type fields struct {
		svc *Service
	}
	type args struct {
		rows [][]string
	}

	var routes []entity.Route
	tests := []struct {
		name       string
		fields     fields
		rows       [][]string
		wantRoutes []entity.Route
		wantErr    bool
	}{
		{
			name: "Parse with success",
			rows: [][]string{{"GRU", "BRL", "20"}, {"CDG", "ORL", "5"}},
			wantRoutes: []entity.Route{
				{
					WhereFrom: "GRU",
					WhereTo:   "BRL",
					Price:     20,
				},
				{
					WhereFrom: "CDG",
					WhereTo:   "ORL",
					Price:     5,
				},
			},
			wantErr: false,
		},
		{
			name:       "Parse with error",
			rows:       [][]string{{"GRU", "BRL", "Randon"}, {"CDG", "ORL", "5"}},
			wantRoutes: routes,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &travelService{
				svc: tt.fields.svc,
			}
			gotRoutes, err := s.parseRowsToStruct(tt.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRowsToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRoutes, tt.wantRoutes) {
				t.Errorf("parseRowsToStruct() = %v, want %v", gotRoutes, tt.wantRoutes)
			}
		})
	}
}

func Test_parametersIsValid(t *testing.T) {
	type fields struct {
		svc *Service
	}

	tests := []struct {
		name      string
		fields    fields
		routes    []entity.Route
		whereFrom string
		whereTo   string
		want      bool
	}{
		{
			name: "Validated",
			routes: []entity.Route{
				{
					WhereFrom: "GRU",
					WhereTo:   "CDG",
					Price:     50,
				},
				{
					WhereFrom: "BRC",
					WhereTo:   "ORL",
					Price:     50,
				},
			},
			whereFrom: "GRU",
			whereTo:   "ORL",
			want:      true,
		},
		{
			name: "Invalid where from",
			routes: []entity.Route{
				{
					WhereFrom: "GRU",
					WhereTo:   "CDG",
					Price:     50,
				},
				{
					WhereFrom: "BRC",
					WhereTo:   "ORL",
					Price:     50,
				},
			},
			whereFrom: "SCL",
			whereTo:   "BRC",
			want:      false,
		},
		{
			name: "Invalid where to",
			routes: []entity.Route{
				{
					WhereFrom: "GRU",
					WhereTo:   "CDG",
					Price:     50,
				},
				{
					WhereFrom: "BRC",
					WhereTo:   "ORL",
					Price:     50,
				},
			},
			whereFrom: "CDG",
			whereTo:   "BRA",
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &travelService{
				svc: tt.fields.svc,
			}
			if got := s.parametersIsValid(tt.routes, tt.whereFrom, tt.whereTo); got != tt.want {
				t.Errorf("parametersIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_travelService_addVertexAndArcs(t *testing.T) {
	type fields struct {
		svc *Service
	}
	type args struct {
		routes []entity.Route
		graph  *dijkstra.Graph
	}

	sucessArgs := args{
		routes: []entity.Route{
			{
				WhereFrom: "BRC",
				WhereTo:   "SCL",
				Price:     55,
			},
			{
				WhereFrom: "CDG",
				WhereTo:   "GRU",
				Price:     55,
			},
		},
		graph: dijkstra.NewGraph(),
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		wantPlaceIDs    map[string]int
		wantPlaceValues map[int]string
	}{
		{
			name: "Added with success",
			args: sucessArgs,
			wantPlaceIDs: map[string]int{
				"BRC": 0,
				"SCL": 1,
				"CDG": 2,
				"GRU": 3,
			},
			wantPlaceValues: map[int]string{
				0: "BRC",
				1: "SCL",
				2: "CDG",
				3: "GRU",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &travelService{
				svc: tt.fields.svc,
			}
			gotPlaceIDs, gotPlaceValues := s.addVertexAndArcs(tt.args.routes, tt.args.graph)
			if !reflect.DeepEqual(gotPlaceIDs, tt.wantPlaceIDs) {
				t.Errorf("travelService.addVertexAndArcs() gotPlaceIDs = %v, want %v", gotPlaceIDs, tt.wantPlaceIDs)
			}
			if !reflect.DeepEqual(gotPlaceValues, tt.wantPlaceValues) {
				t.Errorf("travelService.addVertexAndArcs() gotPlaceValues = %v, want %v", gotPlaceValues, tt.wantPlaceValues)
			}
		})
	}
}
