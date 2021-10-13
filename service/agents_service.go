package service

import (
	"job-master/config"
	"job-master/data"
)

func AddAgent(agent *data.Agent)  {
	if err := config.DB.Create(agent); err != nil {

	}
}
