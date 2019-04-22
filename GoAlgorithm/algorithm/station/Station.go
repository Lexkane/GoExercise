package station

import (
	"encoding/csv"
	"fmt"
	mapset "github.com/fatih/set"
	"os"
)

type Station struct {
	Name        string
	Connections Connection
}
type Connection mapset.Set

var allConnections map[string]Connection

//initialisation of Connections
func init() {

	csvFile, err := os.Open("../resource/stations.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneStation Station
	//var allconnections map[string]Connection

	for _, value := range csvData {
		var stationlist []string
		oneStation.Name = value[0]
		stationlist = value[1:]
		for _, value := range stationlist {
			oneStation.Connections.Add(value)
		}
		allConnections[oneStation.Name] = oneStation.Connections
	}

}

//evaluate whether two stations are connected
func isConnected(stationInput Station, stationOutput Station) bool {
	if stationInput.Connections.Has(stationOutput) {
		return true
	}

	var visitedList Connection

	for _, value := range allConnections {
		visitedList.Add(value)
		return isMultipleConnected(stationInput, stationOutput, visitedList)
	}

	return false

}

//evaluate whether station is connected to set of stations
func isMultipleConnected(stationInput Station, stationOutput Station, ConnectList Connection) bool {
	if ConnectList.Has(stationOutput) {
		return true
	}

	for _, value := range allConnections {
		if !ConnectList.Has(value) {
			ConnectList.Add(value)
			return isConnected(stationInput, stationOutput)
		}
	}
	return false
}
