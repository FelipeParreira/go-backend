package servers

import (
	"github.com/gin-gonic/gin"
	"go-backend/domain/servers"
	"go-backend/services"
	"go-backend/utils/errors"
	"net/http"
)

// Servers godoc
// @Tags servers
// @Summary Get server info
// @Description Get summary info concerning a server.
// @Accept  json
// @Produce json
// @Body
// @Param hostname body servers.HostName true "server hostname" server7
// @Success 200 {object} servers.SummaryInfo
// @Failure 400 {object} errors.RestErr "Invalid JSON body. Body should have a property called `hostname` of type string."
// @Failure 404 {object} errors.RestErr "No server with hostname: `hostname`."
// @Router /servers/summary [get]
func HandleGetSummary(c *gin.Context) {
	var serverInfoRequest servers.ServerInfoRequest
	if err := c.ShouldBindJSON(&serverInfoRequest); err != nil || serverInfoRequest.Hostname == "" {
		restErr := errors.NewBadRequestError("Invalid JSON body. Body should have a non-empty property called `hostname` of type string.")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.ServersService.GetSummaryInfo(serverInfoRequest)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
