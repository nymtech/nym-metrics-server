package metrics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nymtech/nym-directory/models"
)

// Config for this controller
type Config struct {
	Sanitizer Sanitizer
	Service   IService
}

// controller is the metrics controller
type controller struct {
	service   IService
	sanitizer Sanitizer
}

const MaxDesiredRequests = 50 // per second
const MinReportDelay uint64 = 5 // seconds

var nextReportDelay = MinReportDelay


// we don't care about structure itself, we just want to know the count
type Topology struct {
	Gateways []interface{} `json:"gateways"`
	MixNodes []interface{} `json:"mixNodes"`
}

func nodesCount(validatorAddress string) int64 {
	resp, err := http.Get(validatorAddress + "/api/mixmining/topology")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to obtain network topology - %v", err)
		return - 1
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to obtain network topology - %v", err)
		return -1
	}

	var topology Topology
	err = json.Unmarshal(body, &topology)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to obtain network topology - %v", err)
		return - 1
	}


	return int64(len(topology.MixNodes))
}



func DynamicallyUpdateReportDelay(validatorAddress string) {
	updateTicker := time.NewTicker(time.Minute)
	for {
		<-updateTicker.C
		onlineNodes := nodesCount(validatorAddress)

		if onlineNodes > 0 {
			newNextReportDelay := uint64(onlineNodes / MaxDesiredRequests)
			if newNextReportDelay < MinReportDelay {
				// no point in sending it SO often
				newNextReportDelay = MinReportDelay
			}

			atomic.StoreUint64(&nextReportDelay, newNextReportDelay)
		}
	}
}

// Controller ...
type Controller interface {
	CreateMixMetric(c *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

// New returns a new metrics.Controller...
func New(cfg Config) Controller {
	return &controller{cfg.Service, cfg.Sanitizer}
}

func (controller *controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/api/metrics/mixes", controller.CreateMixMetric)
	router.GET("/api/metrics/mixes", controller.ListMixMetrics)
}

// CreateMixMetric ...
// @Summary Create a metric detailing how many messages a given mixnode sent and received
// @Description For demo and debug purposes it gives us the ability to generate useful visualisations of network traffic.
// @ID createMixMetric
// @Accept  json
// @Produce  json
// @Tags metrics
// @Param   object      body   models.MixMetric     true  "object"
// @Success 201 {object} models.MixMetricInterval
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/metrics/mixes [post]
func (controller *controller) CreateMixMetric(c *gin.Context) {
	var metric models.MixMetric
	if err := c.ShouldBindJSON(&metric); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sanitized := controller.sanitizer.Sanitize(metric)
	controller.service.CreateMixMetric(sanitized)

	nextReportDelay := atomic.LoadUint64(&nextReportDelay)

	interval := models.MixMetricInterval{
		NextReportIn: nextReportDelay,
	}

	c.JSON(http.StatusCreated, interval)
}

// ListMixMetrics lists mixnode activity
// @Summary Lists mixnode activity in the past 3 seconds
// @Description For demo and debug purposes it gives us the ability to generate useful visualisations of network traffic.
// @ID listMixMetrics
// @Accept  json
// @Produce  json
// @Tags metrics
// Param   object      body   models.ObjectRequest     true  "object"
// @Success 200 {array} models.MixMetric
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/metrics/mixes [get]
func (controller *controller) ListMixMetrics(c *gin.Context) {
	metrics := controller.service.List()
	c.JSON(http.StatusOK, metrics)
}
