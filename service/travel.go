package service

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/RyanCarrier/dijkstra"
	"github.com/diegoclair/best-route-travel/domain/contract"
	"github.com/diegoclair/best-route-travel/domain/entity"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type travelService struct {
	svc *Service
}

//newTravelService return a new instance of the service
func newTravelService(svc *Service) contract.TravelService {
	return &travelService{
		svc: svc,
	}
}

func (s *travelService) readFile() (routes []entity.Filedata, err error) {

	file, err := os.Open("input-file.csv")
	if err != nil {
		return routes, err
	}
	rows, err := csv.NewReader(file).ReadAll()
	file.Close()
	if err != nil {
		return routes, err
	}

	for i := 0; i < len(rows); i++ {
		price, _ := strconv.Atoi(rows[i][2])
		row := entity.Filedata{
			WhereFrom: rows[i][0],
			WhereTo:   rows[i][1],
			Price:     int64(price),
		}

		routes = append(routes, row)
	}
	return routes, nil
}

func (s *travelService) GetBestRoute(whereFrom, whereTo string) (bestRoute entity.BestRoute, restErr resterrors.RestErr) {

	filedata, err := s.readFile()
	if err != nil {
		logger.Error("Error to read the file: ", err)
		return bestRoute, resterrors.NewInternalServerError("Internal server error")
	}

	graph := dijkstra.NewGraph()
	placeIDs, placeValues := s.AddVertexAndArcs(filedata, graph)

	best, err := graph.Shortest(placeIDs[whereFrom], placeIDs[whereTo])
	if err != nil {
		logger.Error("Error to find the shortest route: ", err)
		return bestRoute, resterrors.NewInternalServerError("Internal server error")
	}

	bestRoute.Price = best.Distance
	for i := 0; i < len(best.Path); i++ {
		if i == 0 {
			bestRoute.Route = placeValues[best.Path[i]]
			continue
		}
		bestRoute.Route += " - " + placeValues[best.Path[i]]
	}

	return bestRoute, nil
}

func (s *travelService) AddVertexAndArcs(routes []entity.Filedata, graph *dijkstra.Graph) (placeIDs map[string]int, placeValues map[int]string) {

	placeIDs = make(map[string]int, 0)
	placeValues = make(map[int]string, 0)

	for i := 0; i < len(routes); i++ {

		_, ok := placeIDs[routes[i].WhereFrom]
		if !ok {
			placeID := len(placeIDs)
			placeIDs[routes[i].WhereFrom] = placeID
			placeValues[placeID] = routes[i].WhereFrom
			graph.AddVertex(placeID)
		}
		_, ok = placeIDs[routes[i].WhereTo]
		if !ok {
			placeID := len(placeIDs)
			placeIDs[routes[i].WhereTo] = placeID
			placeValues[placeID] = routes[i].WhereTo
			graph.AddVertex(placeID)
		}
	}

	for i := 0; i < len(routes); i++ {
		graph.AddArc(placeIDs[routes[i].WhereFrom], placeIDs[routes[i].WhereTo], routes[i].Price)
	}

	return
}
