package services

import "go-backend/domain/servers"

func getValues(info []servers.BaseInfo) []float64 {
	vsm := make([]float64, len(info))
	for i, v := range info {
		vsm[i] = v.Value
	}
	return vsm
}

func getBaseInfo(vs []servers.Info, f func(servers.Info) servers.BaseInfo) []servers.BaseInfo {
	vsm := make([]servers.BaseInfo, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
