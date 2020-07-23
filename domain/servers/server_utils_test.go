package servers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	setsOfInfos := [][]Info{
		{
			Info{
				Hostname: "server7",
				CPULoad: BaseInfo{
					Value: 90,
				},
			},
			Info{
				Hostname: "server2",
				CPULoad: BaseInfo{
					Value: 90,
				},
			},
			Info{
				Hostname: "server7",
				CPULoad: BaseInfo{
					Value: 80,
				},
			},
		},
		{},
	}

	setsOfFilteredInfos := [][]Info{
		{
			Info{
				Hostname: "server7",
				CPULoad: BaseInfo{
					Value: 90,
				},
			},
			Info{
				Hostname: "server7",
				CPULoad: BaseInfo{
					Value: 80,
				},
			},
		},
		{},
	}

	for i, info := range setsOfInfos {
		filteredInfo := setsOfFilteredInfos[i]
		assert.EqualValues(t, Filter(info, func(info Info) bool { return info.Hostname == "server7" }), filteredInfo,
			"should filter the infos according to a predicate function")
	}
}
