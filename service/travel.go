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

// FiledataPath is the file name that are used to control the routes list
const FiledataPath string = "possible_routes.csv"

func (s *travelService) GetBestRoute(whereFrom, whereTo string) (bestRoute entity.BestRoute, restErr resterrors.RestErr) {

	filedata, err := s.readFile()
	if err != nil {
		logger.Error("Error to read the file: ", err)
		return bestRoute, resterrors.NewInternalServerError("Internal server error")
	}

	graph := dijkstra.NewGraph()
	placeIDs, placeValues := s.addVertexAndArcs(filedata, graph)

	if !s.parametersIsValid(filedata, whereFrom, whereTo) {
		logger.Error("Some route don't exists: from: "+whereFrom+" - to: "+whereTo, nil)
		return bestRoute, resterrors.NewBadRequestError("Some route don't exists")
	}

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

func (s *travelService) addVertexAndArcs(routes []entity.Route, graph *dijkstra.Graph) (placeIDs map[string]int, placeValues map[int]string) {

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

func (s *travelService) readFile() (routes []entity.Route, err error) {

	file, err := os.Open(FiledataPath)
	if err != nil {
		return routes, err
	}
	defer file.Close()
	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return routes, err
	}

	return s.parseRowsToStruct(rows)
}

func (s *travelService) parseRowsToStruct(rows [][]string) (routes []entity.Route, err error) {
	for i := 0; i < len(rows); i++ {
		price, err := strconv.Atoi(rows[i][2])
		if err != nil {
			return routes, err
		}
		row := entity.Route{
			WhereFrom: rows[i][0],
			WhereTo:   rows[i][1],
			Price:     int64(price),
		}

		routes = append(routes, row)
	}
	return routes, nil
}

func (s *travelService) parametersIsValid(routes []entity.Route, whereFrom, whereTo string) bool {
	var whereFromIsValid, whereToIsValid bool

	for i := 0; i < len(routes); i++ {
		if routes[i].WhereFrom == whereFrom {
			whereFromIsValid = true
		}
		if routes[i].WhereTo == whereTo {
			whereToIsValid = true
		}
		if whereFromIsValid && whereToIsValid {
			break
		}
	}
	return whereFromIsValid && whereToIsValid
}

func (s *travelService) AddNewRoute(route entity.Route) (restErr resterrors.RestErr) {

	file, err := os.OpenFile(FiledataPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		logger.Error("Error to create the file: ", err)
		return resterrors.NewInternalServerError("Error to add new route")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	//writer.Write([]string{})
	err = writer.Write([]string{route.WhereFrom, route.WhereTo, strconv.Itoa(int(route.Price))})
	if err != nil {
		logger.Error("Error to write the new route into file: ", err)
		return resterrors.NewInternalServerError("Error to add new route")
	}

	return nil
}
