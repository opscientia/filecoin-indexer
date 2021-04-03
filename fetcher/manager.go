package fetcher

import (
	"time"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/pipeline"
	"github.com/figment-networks/filecoin-indexer/store"
	"github.com/figment-networks/filecoin-indexer/worker"
)

// Manager represents a fetcher manager
type Manager struct {
	cfg    *config.Config
	pool   *worker.Pool
	store  *store.Store
	client *client.Client
}

// NewManager creates a fetcher manager
func NewManager(cfg *config.Config, pool *worker.Pool, store *store.Store, client *client.Client) (*Manager, error) {
	manager := Manager{
		cfg:    cfg,
		pool:   pool,
		store:  store,
		client: client,
	}

	return &manager, nil
}

// Run starts the fetcher manager
func (m *Manager) Run() error {
	defer m.pool.Stop()

	m.pool.Run(m.handleResponse)

	for {
		jobs, err := m.getJobs()
		if err != nil {
			return err
		}

		for _, job := range jobs {
			if err := m.scheduleJob(&job); err != nil {
				return err
			}
		}

		m.pool.Wait()

		time.Sleep(m.cfg.SyncIntervalDuration)
	}
}

func (m *Manager) getJobs() ([]model.Job, error) {
	jobs, err := m.store.Job.FindAllUnfinished()
	if err != nil {
		return nil, err
	}

	if len(jobs) == 0 {
		return m.createJobs()
	}

	return jobs, nil
}

func (m *Manager) createJobs() ([]model.Job, error) {
	var jobs []model.Job

	hr, err := m.getHeightRange()
	if err != nil {
		return nil, err
	}

	for h := hr.StartHeight(); h <= hr.EndHeight(); h++ {
		height := h
		jobs = append(jobs, model.Job{Height: &height})
	}

	err = m.store.Job.Create(jobs)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (m *Manager) getHeightRange() (*pipeline.HeightRange, error) {
	latestHeight, err := m.client.Epoch.GetLatestHeight()
	if err != nil {
		return nil, err
	}

	lastHeight, err := m.store.Job.LastFinishedHeight()
	if err != nil {
		lastHeight = -1
	}

	hr := pipeline.HeightRange{
		LatestHeight:  latestHeight,
		LastHeight:    lastHeight,
		InitialHeight: m.cfg.InitialHeight,
		BatchSize:     m.cfg.BatchSize,
	}

	if err := hr.Validate(false); err != nil {
		return nil, err
	}

	return &hr, nil
}

func (m *Manager) scheduleJob(job *model.Job) error {
	m.pool.Process(*job.Height)

	now := time.Now()

	job.StartedAt = &now
	job.RunCount++

	err := m.store.Job.Save(job)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) handleResponse(res worker.Response) {
	now := time.Now()

	job, err := m.store.Job.FindByHeight(res.Height)
	if err != nil {
		panic(err)
	}

	if res.Success {
		job.FinishedAt = &now
	} else {
		job.LastError = res.Error
	}

	err = m.store.Job.Save(job)
	if err != nil {
		panic(err)
	}
}
