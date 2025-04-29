package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/capybartender/task-cli/internal/models"
)

const defaultFileName = "tasks.json"

type Storage struct {
	filePath string
}

func Init(filePath string) *Storage {
	if filePath == "" {
		filePath = defaultFileName
	}
	return &Storage{filePath: filePath}
}

func validateFilePath(filePath string) (shouldCreate bool, err error) {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) || fileInfo.IsDir() {
		return true, nil		
	}

	return false, err
}

func (s *Storage) Load() (*models.TaskList, error) {
	shouldCreate, err := validateFilePath(s.filePath)

	if shouldCreate {
		return &models.TaskList{}, nil		
	}
	
	if err != nil {
		return nil, err
	}

	jsonFile, err := os.Open(s.filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	var tasks models.TaskList
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		return nil, err
	}

	if tasks == nil {
		return &models.TaskList{}, nil
	}

	if len(tasks) == 0 {
		return &models.TaskList{}, nil
	}

	return &tasks, nil
}

func (s *Storage) Save(tasks *models.TaskList) error {
	shouldCreate, err := validateFilePath(s.filePath)
	if shouldCreate {
		return &models.TaskList{}, nil		
	}
	
	return nil
}
