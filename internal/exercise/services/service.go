package services

import (
	"exercise-api/internal/exercise/entities"
	"exercise-api/internal/exercise/model"
)

// exercise
type ExerciseValidator interface {
	ValidateCreateExercisePayload(*entities.ExerciseCreateRequest) error
}

type ExerciseRepository interface {
	Create(*entities.ExerciseModel) error
	GetById(int) (*model.GetExerciseById, error)
}

// question
type QuestionValidator interface {
	ValidateCreateQuestionPayload(*entities.QuestionCreateRequest) error
}

type QuestionRepository interface {
	Create(*entities.QuestionModel) error
	GetById(int) (*model.GetQuestionById, error)
}

// answer
type AnswerValidator interface {
	ValidateCreateAnswerPayload(*entities.AnswerCreateRequest) error
}

type AnswerRepository interface {
	Create(*entities.AnswerModel) error
}
