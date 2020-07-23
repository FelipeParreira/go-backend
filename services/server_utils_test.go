package services

import (
	"github.com/stretchr/testify/assert"
	"go-backend/domain/servers"
	"testing"
)

func TestGetValues(t *testing.T) {
	setOfinfos := [][]servers.BaseInfo{
		{
			{
				Value: 12,
				Unit:  "GB",
			},
			{
				Value: 23,
				Unit:  "%",
			},
			{
				Value: 90,
				Unit:  "GB",
			},
		},
		{},
	}
	setOfValues := [][]float64{{12, 23, 90}, {}}

	for i, infos := range setOfinfos {
		values := setOfValues[i]
		assert.EqualValues(t, getValues(infos), values, "should extract the values from base info")
	}
}

func TestGetBaseInfo(t *testing.T) {
	setOfinfos := [][]servers.Info{
		{
			{
				CPULoad: servers.BaseInfo{
					Value: 12,
					Unit:  "GB",
				},
			},
			{
				MemoryUsage: servers.BaseInfo{
					Value: 12,
					Unit:  "GB",
				},
			},
			{
				CPULoad: servers.BaseInfo{
					Value: 27,
					Unit:  "GB",
				},
			},
		},
		{},
	}
	setOfBaseInfo := [][]servers.BaseInfo{{
		servers.BaseInfo{
			Value: 12,
			Unit:  "GB",
		},
		servers.BaseInfo{},
		servers.BaseInfo{
			Value: 27,
			Unit:  "GB",
		},
	}, {}}

	for i, infos := range setOfinfos {
		values := setOfBaseInfo[i]
		assert.EqualValues(
			t,
			getBaseInfo(infos, func(info servers.Info) servers.BaseInfo { return info.CPULoad }),
			values,
			"should extract the values from base info")
	}
}
