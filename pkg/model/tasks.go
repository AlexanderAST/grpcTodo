package model

type Folder struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Time        string `json:"time"`
	UserID      string `json:"user_id"`
	FolderID    string `json:"folder_id"`
}
