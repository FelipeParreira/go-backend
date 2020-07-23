package services

import (
	"github.com/montanaflynn/stats"
	"go-backend/domain/servers"
	"go-backend/utils/errors"
)

var (
	ServersService serversServiceInterface = &serversService{}
)

type serversService struct{}

type serversServiceInterface interface {
	GetSummaryInfo(servers.ServerInfoRequest) (*servers.SummaryInfo, *errors.RestErr)
	sumariseServerInfo(serverInfo []servers.Info) *servers.SummaryInfo
	summariseBaseInfo(info []servers.BaseInfo) servers.BaseSummaryInfo
}

func (s *serversService) GetSummaryInfo(serverRequestInfo servers.ServerInfoRequest) (*servers.SummaryInfo, *errors.RestErr) {
	DAO := &servers.Server{serverRequestInfo.Hostname}
	serverInfo, err := DAO.GetServerInfo()
	if err != nil {
		return nil, err
	}

	summaryInfo := s.sumariseServerInfo(serverInfo)
	summaryInfo.Hostname = serverRequestInfo.Hostname
	return summaryInfo, nil
}

func (s *serversService) sumariseServerInfo(serverInfo []servers.Info) *servers.SummaryInfo {
	CPULoadInfo := getBaseInfo(serverInfo, func(info servers.Info) servers.BaseInfo { return info.CPULoad })
	MemoryUsageInfo := getBaseInfo(serverInfo, func(info servers.Info) servers.BaseInfo { return info.MemoryUsage })
	DiskUsageInfo := getBaseInfo(serverInfo, func(info servers.Info) servers.BaseInfo { return info.DiskUsage })

	return &servers.SummaryInfo{
		CPULoad:     s.summariseBaseInfo(CPULoadInfo),
		MemoryUsage: s.summariseBaseInfo(MemoryUsageInfo),
		DiskUsage:   s.summariseBaseInfo(DiskUsageInfo),
	}
}

func (s *serversService) summariseBaseInfo(info []servers.BaseInfo) servers.BaseSummaryInfo {
	values := getValues(info)
	mean, _ := stats.Mean(values)
	median, _ := stats.Median(values)
	modes, _ := stats.Mode(values)

	var mode float64
	if len(modes) == 0 {
		mode = 0
	} else {
		mode = modes[0]
	}

	return servers.BaseSummaryInfo{
		Average: mean,
		Mode:    mode,
		Median:  median,
		Unit:    info[0].Unit,
	}
}
