package webif

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/p2pquake/userquake-aggregator/pkg/aggregate"
	"github.com/p2pquake/userquake-aggregator/pkg/epsp"
	"github.com/p2pquake/userquake-aggregator/pkg/evaluate"
)

type Server struct {
}

func (s Server) Start(port int) {
	r := gin.Default()
	r.POST("/calculate", func(c *gin.Context) {
		// parse
		epspRecords := []epsp.Record{}

		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.Unmarshal(body, &epspRecords); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		sort.Slice(epspRecords, func(i, j int) bool { return epspRecords[i].Time.Time.Before(*epspRecords[j].Time.Time) })

		// aggregate & evaluate
		aggregationResults := aggregate.CompatibleAggregator{}.Aggregate(epspRecords)

		evaluationResults := []evaluate.Result{}

		for _, r := range aggregationResults {
			result := evaluate.CompatibleEvaluator{}.Evaluate(r)
			evaluationResults = append(evaluationResults, result)
		}

		c.JSON(200, evaluationResults)
	})
	r.Run(":" + strconv.Itoa(port))
}
