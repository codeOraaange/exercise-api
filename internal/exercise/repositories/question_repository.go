package repositories

import (
	exerciseEntities "exercise-api/internal/exercise/entities"
	"exercise-api/internal/exercise/shared"

	"gorm.io/gorm"
)

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(
	db *gorm.DB,
) *questionRepository {
	return &questionRepository{
		db: db,
	}
}

func (q *questionRepository) Create(question *exerciseEntities.QuestionModel) (err error) {
	err = q.db.Create(&question).Error
	return
}

func (q *questionRepository) GetById(questionId int) (question *shared.GetQuestionByIdDTO, err error) {
	err = q.db.Table("questions").
		Select(
			"id",
			"exercise_id",
			"body",
		).
		Where("id = ?", questionId).First(&question).Error
	return
}