package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// タスクに関する操作を定義するインターフェイス
type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository はtaskRepository構造体の新しいインスタンスを作成し、ITaskRepositoryインターフェースを返します。
// 引数dbはGORMによるデータベース接続を指定します。
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

// GetAllTasks は指定されたユーザーIDに紐づく全てのタスクを取得します。
// 取得したタスクは引数tasksに格納され、エラーが発生した場合はそのエラーを返します。
func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

// GetTaskById は指定されたタスクIDとユーザーIDに基づいて特定のタスクを取得します。
// 成功した場合、取得したタスクは引数taskに格納され、エラーが発生した場合はそのエラーを返します。
func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

// CreateTask は新しいタスクをデータベースに追加します。
// タスクの追加に成功した場合はnilを返し、エラーが発生した場合はそのエラーを返します。
func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTask は指定されたタスクIDのタスクの内容を更新します。
// 引数として更新内容を含むtaskオブジェクトと、ユーザーID、タスクIDが必要です。
// 更新が成功した場合はnilを返し、何らかのエラーが発生した場合はそのエラーを返します。
func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	updates := map[string]interface{}{
		"title":    task.Title,
		"deadline": task.Deadline,
	}
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

// DeleteTask は指定されたユーザーIDとタスクIDに基づいてタスクを削除します。
// 削除に成功した場合はnilを返し、何らかのエラーが発生した場合や対象のタスクが存在しない場合はエラーを返します。
func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
