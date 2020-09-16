package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ProjectTriggerService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewProjectTriggerService(sling *sling.Sling, uriTemplate string) *ProjectTriggerService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &ProjectTriggerService{
		sling: sling,
		path:  path,
	}
}

func (s *ProjectTriggerService) Get(id string) (*model.ProjectTrigger, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("ProjectTriggerService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.ProjectTrigger), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

func (s *ProjectTriggerService) GetByProjectID(id string) (*[]model.ProjectTrigger, error) {
	var triggersByProject []model.ProjectTrigger

	triggers, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	triggersByProject = append(triggersByProject, *triggers...)

	return &triggersByProject, nil
}

// GetAll returns all instances of a ProjectTrigger.
func (s *ProjectTriggerService) GetAll() (*[]model.ProjectTrigger, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	var p []model.ProjectTrigger
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.ProjectTriggers), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.ProjectTriggers)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// Add creates a new ProjectTrigger.
func (s *ProjectTriggerService) Add(projectTrigger *model.ProjectTrigger) (*model.ProjectTrigger, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if projectTrigger == nil {
		return nil, errors.New("ProjectTriggerService: invalid parameter, projectTrigger")
	}

	err = projectTrigger.Validate()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, projectTrigger, new(model.ProjectTrigger), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

func (s *ProjectTriggerService) Delete(id string) error {
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("ProjectTriggerService: invalid parameter, id")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *ProjectTriggerService) Update(resource *model.ProjectTrigger) (*model.ProjectTrigger, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if resource == nil {
		return nil, errors.New("ProjectTriggerService: invalid parameter, resource")
	}

	err = resource.Validate()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.ProjectTrigger), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

func (s *ProjectTriggerService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("ProjectTriggerService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("ProjectTriggerService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &ProjectTriggerService{}
