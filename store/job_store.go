package store

import (
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
)

type jobStore struct {
	db *gorm.DB
}

// Create bulk-inserts the job records
func (js *jobStore) Create(jobs []model.Job) error {
	return js.db.Create(jobs).Error
}

// Save updates the job record
func (js *jobStore) Save(job *model.Job) error {
	return js.db.Save(job).Error
}

// FindByHeight retrieves a job record for a given height
func (js *jobStore) FindByHeight(height int64) (*model.Job, error) {
	var job model.Job

	err := js.db.Where("height = ?", height).Take(&job).Error
	if err != nil {
		return nil, err
	}

	return &job, err
}

// FindAllUnfinished retrieves all of the unfinished jobs
func (js *jobStore) FindAllUnfinished() ([]model.Job, error) {
	var jobs []model.Job

	err := js.db.
		Where("finished_at IS NULL").
		Order("height ASC").
		Find(&jobs).
		Error

	if err != nil {
		return nil, err
	}

	return jobs, nil
}

// LastFinishedHeight returns the most recent finished height
func (js *jobStore) LastFinishedHeight() (int64, error) {
	var result int64

	err := js.db.
		Table("jobs").
		Where("finished_at IS NOT NULL").
		Select("MAX(height)").
		Scan(&result).
		Error

	if err != nil {
		return 0, err
	}

	return result, nil
}
