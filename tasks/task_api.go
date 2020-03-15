package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type Task struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	User      string    `json:"user"`
	Completed string    `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateTaskTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError := db.CreateTable(&Task{}, opts)

	if createError != nil {
		log.Printf("Error while creating todo table, Reasion: %v\n", createError)
		return createError
	}

	log.Printf("Task Table Created")

	return nil
}

var dbConnect *pg.DB

func InitializeDB(db *pg.DB) {
	dbConnect = db
}

func GetTasks(c *gin.Context) {
	var tasks []Task
	userId := c.Param("userId")
	log.Printf("user: %v", userId)
	err := dbConnect.Model(&tasks).Select()
	if err != nil {
		log.Printf("Error while getting tasks for a user, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Tasks not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Tasks for user",
		"data":    tasks,
	})
	return
}

func PostTask(c *gin.Context) {
	var task Task
	c.BindJSON(&task)
	id := guuid.New().String()
	message := task.Message
	user := task.User
	completed := "false"
	created_at := time.Now()
	updated_at := time.Now()

	insertError := dbConnect.Insert(&Task{
		ID:        id,
		Message:   message,
		User:      user,
		Completed: completed,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	})

	if insertError != nil {
		log.Printf("Error while inserting new Task, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Task created successfully",
	})
	return
}

func UpdateTask(c *gin.Context) {
	taskId := c.Param("taskId")
	var task Task
	c.BindJSON(&task)

	_, err := dbConnect.Model(&Task{}).Set("completed = ?", "true").Where("id = ?", taskId).Update()

	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Task Edited successfully",
	})
	return
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")
	task := &Task{ID: taskId}
	err := dbConnect.Delete(task)

	if err != nil {
		log.Printf("Error while deleting a single task, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something Went Wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Task deleted",
	})
	return
}
