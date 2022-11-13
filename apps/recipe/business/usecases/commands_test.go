package usecases

import (
	"github.com/baransonmez/coffein/internal/recipe/business/domain"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestNewRecipe_Validate(t *testing.T) {
	type fields struct {
		UserID      string
		BeanID      string
		Description string
		Steps       []domain.Step
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "complete command success", fields: fields{
			UserID:      uuid.New().String(),
			BeanID:      uuid.New().String(),
			Description: "description",
			Steps:       []domain.Step{{Order: 1, Description: "step 1", DurationInSeconds: 3}},
		}, wantErr: false},
		{name: "userId is Empty fail", fields: fields{
			UserID:      "",
			BeanID:      uuid.New().String(),
			Description: "description",
			Steps:       []domain.Step{{Order: 1, Description: "step 1", DurationInSeconds: 3}},
		}, wantErr: true},
		{name: "BeanID is Empty fail", fields: fields{
			UserID:      uuid.New().String(),
			BeanID:      "",
			Description: "description",
			Steps:       []domain.Step{{Order: 1, Description: "step 1", DurationInSeconds: 3}},
		}, wantErr: true},
		{name: "Description is Empty fail", fields: fields{
			UserID:      uuid.New().String(),
			BeanID:      uuid.New().String(),
			Description: "",
			Steps:       []domain.Step{{Order: 1, Description: "step 1", DurationInSeconds: 3}},
		}, wantErr: true},
		{name: "Steps length is 0 fail", fields: fields{
			UserID:      uuid.New().String(),
			BeanID:      uuid.New().String(),
			Description: "description",
			Steps:       []domain.Step{},
		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NewRecipe{
				UserID:      tt.fields.UserID,
				BeanID:      tt.fields.BeanID,
				Description: tt.fields.Description,
				Steps:       tt.fields.Steps,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewRecipe_toDomainModel(t *testing.T) {
	type fields struct {
		UserID      string
		BeanID      string
		Description string
		Steps       []domain.Step
	}
	userId := uuid.New()
	beanId := uuid.New()
	tests := []struct {
		name   string
		fields fields
		want   domain.Recipe
	}{

		{name: "create domain object", fields: fields{
			UserID:      userId.String(),
			BeanID:      beanId.String(),
			Description: "description",
			Steps:       []domain.Step{{Order: 1, Description: "step desc", DurationInSeconds: 3}},
		}, want: domain.Recipe{
			UserID:      userId,
			BeanID:      beanId,
			Description: "description",
			Steps:       []domain.Step{{Order: 1, Description: "step desc", DurationInSeconds: 3}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NewRecipe{
				UserID:      tt.fields.UserID,
				BeanID:      tt.fields.BeanID,
				Description: tt.fields.Description,
				Steps:       tt.fields.Steps,
			}
			got := r.toDomainModel()

			if !reflect.DeepEqual(got.BeanID, tt.want.BeanID) {
				t.Errorf("toDomainModel() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got.UserID, tt.want.UserID) {
				t.Errorf("toDomainModel() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got.Description, tt.want.Description) {
				t.Errorf("toDomainModel() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got.Steps, tt.want.Steps) {
				t.Errorf("toDomainModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
