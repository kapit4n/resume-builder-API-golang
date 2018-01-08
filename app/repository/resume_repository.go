package repository

import (
	"errors"
	"rb-rgo/app/models"
)

type ResumeRepository interface {
	GetResumes() []*models.Resume
	GetResumeById(id string) (*models.Resume, error)
	SaveResume(resume *models.Resume) error
}

type DBResumeRepository struct {
	resumes []*models.Resume
}

func New() *DBResumeRepository {
	return &DBResumeRepository{
		resumes: []*models.Resume{
			&models.Resume{"1", "Full Stack Go Developer", "Summary"},
			&models.Resume{"2", "Revel Developer", "Summary"},
			&models.Resume{"3", "Full Stack Javascript Developer", "Summary"},
			&models.Resume{"4", "Full Stack Python Developer", "Summary"},
		},
	}
}

func (r *DBResumeRepository) GetResumes() []*models.Resume {
	return r.resumes
}

func (r *DBResumeRepository) GetResumeById(id string) (*models.Resume, error) {
	for _, resume := range r.resumes {
		if resume.Id == id {
			return resume, nil
		}
	}
	return nil, errors.New("resume not found")
}

func (r *DBResumeRepository) SaveResume(resume *models.Resume) error {
	r.resumes = append(r.resumes, resume)
	return nil
}

var resumeRepository *DBResumeRepository

func GetResumeRepository() (r ResumeRepository) {
	if resumeRepository == nil {
		resumeRepository = New()
	}
	return resumeRepository
}