package backlogimpl

import (
	"cuddly-octo-palm-tree/config"
	bi_repo "cuddly-octo-palm-tree/repo/backlogimpl"

	log "github.com/sirupsen/logrus"
)

type BiServiceImpl struct {
	biRepo bi_repo.BiRepository
	// producer *kafka.Producer
	logger *log.Entry
}

func NewBiService(conf *config.Config, repo bi_repo.BiRepository) BiService {
	return &BiServiceImpl{
		biRepo: repo,
		// producer: producer,
		logger: conf.Logger.ContextLogger.WithFields(log.Fields{
			"type": "service:Bi",
		}),
	}
}
