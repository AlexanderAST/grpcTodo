package store

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"todoGRPC/pkg/model"
)

type TaskRepository struct {
	store *Store
}

func (r *TaskRepository) CreateFolder(f *model.Folder) (string, error) {
	jsonData, err := json.Marshal(f)
	if err != nil {
		return "", err
	}

	err = r.store.db.Set(ctx, "folder:"+f.ID, jsonData, 0).Err()
	if err != nil {
		return "", err
	}
	return "successfully created", nil
}

func (r *TaskRepository) DeleteFolder(id string) string {
	err := r.store.db.Del(ctx, "folder:"+id).Err()
	if err != nil {
		return "cannot delete this folder"
	}

	return "successfully deleted"

}

func (r *TaskRepository) GetAllInFolder(id string) ([]int, error) {
	var cursor uint64
	var keys []int

	ctx := context.Background()
	for {
		results, nextCursor, err := r.store.db.Scan(ctx, cursor, "taskID:*", 10).Result()
		if err != nil {
			log.Printf("Error scanning database: %v", err)
			return nil, err
		}

		for _, key := range results {
			value, err := r.store.db.Get(ctx, key).Result()
			if err != nil {
				log.Printf("Error getting value for key %s: %v", key, err)
				continue
			}

			var task struct {
				ID       string `json:"id"`
				FolderID string `json:"folder_id"`
			}
			if err := json.Unmarshal([]byte(value), &task); err != nil {
				log.Printf("Error unmarshaling JSON for key %s: %v", key, err)
				continue
			}

			if task.FolderID == id {
				taskID, err := strconv.Atoi(task.ID)
				if err != nil {
					log.Printf("Error converting task ID to int for key %s: %v", key, err)
					continue
				}
				keys = append(keys, taskID)
			}
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	return keys, nil
}

func (r *TaskRepository) CreateTask(t *model.Task) (string, error) {
	jsonData, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	err = r.store.db.Set(ctx, "taskID:"+t.ID, jsonData, 0).Err()
	if err != nil {
		return "", err
	}

	return "task successfully created", nil
}

func (r *TaskRepository) DeleteTask(id string) string {
	err := r.store.db.Del(ctx, "taskID:"+id).Err()
	if err != nil {
		return "cannot delete this task"
	}

	return "successfully deleted"

}

func (r *TaskRepository) GetOpenTask(id string) (string, error) {

	result, err := r.store.db.Get(ctx, "taskID:"+id).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r *TaskRepository) GetAllTasks() ([]int, error) {
	var cursor uint64
	var keys []int

	ctx := context.Background()
	for {
		results, nextCursor, err := r.store.db.Scan(ctx, cursor, "taskID:*", 10).Result()
		if err != nil {
			log.Printf("Error scanning database: %v", err)
			return nil, err
		}

		for _, key := range results {
			value, err := r.store.db.Get(ctx, key).Result()
			if err != nil {
				log.Printf("Error getting value for key %s: %v", key, err)
				continue
			}

			var task struct {
				ID       string `json:"id"`
				FolderID string `json:"folder_id"`
			}
			if err := json.Unmarshal([]byte(value), &task); err != nil {
				log.Printf("Error unmarshaling JSON for key %s: %v", key, err)
				continue
			}

			taskID, err := strconv.Atoi(task.ID)
			if err != nil {
				log.Printf("Error converting task ID to int for key %s: %v", key, err)
				continue
			}
			keys = append(keys, taskID)

		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	return keys, nil
}

func (r *TaskRepository) UpdateTask(id string, t *model.Task) (string, error) {
	jsonData, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	err = r.store.db.Set(ctx, "taskID:"+id, jsonData, 0).Err()
	if err != nil {
		return "", err
	}

	return "task successfully changed", nil

}
