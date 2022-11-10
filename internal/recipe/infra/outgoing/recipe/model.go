package recipe

import (
	"github.com/baransonmez/coffein/internal/recipe/business/domain"
	"github.com/google/uuid"
	"time"
)

type Recipe struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	BeanID      string    `db:"bean_id"`
	Description string    `db:"description"`
	Steps       []Step    `db:"step"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

type Step struct {
	RecipeID          string    `db:"recipe_id"`
	StepOrder         uint8     `db:"step_order"`
	Description       string    `db:"description"`
	DurationInSeconds int32     `db:"duration_in_seconds"`
	DateCreated       time.Time `db:"date_created"`
	DateUpdated       time.Time `db:"date_updated"`
}

func (dbRecipe *Recipe) ToRecipe() *domain.Recipe {
	uuidFromString, _ := uuid.Parse(dbRecipe.ID)
	userUuidFromString, _ := uuid.Parse(dbRecipe.UserID)
	beanUuidFromString, _ := uuid.Parse(dbRecipe.BeanID)
	dbToDomainModel := domain.Recipe{
		ID:          uuidFromString,
		Description: dbRecipe.Description,
		UserID:      userUuidFromString,
		BeanID:      beanUuidFromString,
		Steps:       stepsToDomainModel(dbRecipe.Steps),
		DateCreated: dbRecipe.DateCreated,
		DateUpdated: dbRecipe.DateUpdated,
	}
	return &dbToDomainModel
}

func stepsToDomainModel(steps []Step) []domain.Step {
	var stepsVO []domain.Step
	for _, s := range steps {
		stepsVO = append(stepsVO, s.stepToDomainModel())
	}
	return stepsVO
}

func (s Step) stepToDomainModel() domain.Step {
	return domain.Step{
		Description:       s.Description,
		DurationInSeconds: s.DurationInSeconds,
		Order:             s.StepOrder,
	}
}

func StepsFromDomainModel(stepsDO []domain.Step) []Step {
	var stepsDB []Step
	for _, s := range stepsDO {
		stepsDB = append(stepsDB, stepFromDomainModel(s))
	}
	return stepsDB
}

func stepFromDomainModel(s domain.Step) Step {
	return Step{
		Description:       s.Description,
		DurationInSeconds: s.DurationInSeconds,
	}
}
