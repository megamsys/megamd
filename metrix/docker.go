package metrix

import (
	//"encoding/xml"
	//"github.com/megamsys/opennebula-go/metrics"
	"github.com/megamsys/vertice/carton"
	"io/ioutil"
	"time"
)

const DOCKER = "docker"

type Swarm struct {
	Url       string
	RawStatus []byte
}

func (s *Swarm) Prefix() string {
	return "docker"
}

func (s *Swarm) Collect(c *MetricsCollection) (e error) {
	 e = s.ReadStatus()
	if e != nil {
		return
	}
  //
	// s, e := s.ParseStatus(b)
	// if e != nil {
	// 	return
	// }
	// on.CollectMetricsFromStats(c, s)
	 return
}

func (s *Swarm) ReadStatus() (e error) {
		_, e = carton.ProvisionerMap[s.Prefix()].MetricEnvs(time.Now().Add(-10*time.Minute).Unix(),
			time.Now().Unix(),s.Url,ioutil.Discard)
		if e != nil {
			return
		}
  return
}
