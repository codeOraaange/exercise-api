package exercise

import (
	exerciseEntities "exercise-api/internal/exercise/entities"
	"exercise-api/internal/shared/entities"
)

type ExerciseService interface {
	Create(*exerciseEntities.ExerciseCreateRequest) entities.BaseResponse[exerciseEntities.ExerciseCreateResponse]
}