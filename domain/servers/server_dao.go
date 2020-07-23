package servers

import (
	"encoding/json"
	"fmt"
	"go-backend/logger"
	"go-backend/utils/errors"
	"io/ioutil"
	"os"
)

var serverInfo []Info

func (s *Server) GetServerInfo() ([]Info, *errors.RestErr) {
	serverInfo, err := ReadDataFile()
	if err != nil {
		logger.Error("error when retrieving server data", *err)
		return nil, errors.NewInternalServerError("error retrieving data")
	}

	filteredInfo := Filter(serverInfo, func(info Info) bool {
		return info.Hostname == s.Hostname
	})

	if len(filteredInfo) == 0 {
		err := errors.NewNotFoundError(fmt.Sprintf("No server with hostname: `%s`", s.Hostname))
		logger.Error("error when trying to get server info by hostname", err)
		return nil, err
	}

	return filteredInfo, nil
}

func ReadDataFile() ([]Info, *error) {
	if len(serverInfo) == 0 {
		dataFile, err := os.Open("./data.json")
		if err != nil {
			return nil, &err
		}

		byteValue, err := ioutil.ReadAll(dataFile)
		if err != nil {
			return nil, &err
		}

		json.Unmarshal(byteValue, &serverInfo)
	}

	return serverInfo, nil
}
