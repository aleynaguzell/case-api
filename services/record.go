package services

import (
	"case-api/model/record"
	repo "case-api/storage/repository"
)

type RecordService struct {
	repository *repo.RecordsRepository
}

func NewRecordService(repository *repo.RecordsRepository) *RecordService {
	return &RecordService{
		repository: repository,
	}
}

func (s *RecordService) GetRecords(request record.Request) ([]record.Record, error) {
	records, err := s.repository.Get(request)

	if err != nil {
		return nil, err
	}

	return records, nil
}
