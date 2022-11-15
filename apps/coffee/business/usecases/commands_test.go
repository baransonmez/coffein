package usecases

import (
	"errors"
	"github.com/baransonmez/coffein/internal/coffee/business/domain"
	"testing"
	"time"
)

func TestNewCoffeeBean_Validate(t *testing.T) {
	type fields struct {
		Name      string
		Roaster   string
		Origin    string
		Price     int
		RoastDate time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{name: "success", fields: fields{
			Name:      "test",
			Roaster:   "test",
			Origin:    "test",
			Price:     1,
			RoastDate: time.Time{},
		}, wantErr: nil},
		{name: "name is empty fail", fields: fields{
			Name:      "",
			Roaster:   "test",
			Origin:    "test",
			Price:     1,
			RoastDate: time.Time{},
		}, wantErr: NameCannotBeEmptyError},
		{name: "roaster is empty fail", fields: fields{
			Name:      "test",
			Roaster:   "",
			Origin:    "test",
			Price:     1,
			RoastDate: time.Time{},
		}, wantErr: RoasterCannotBeEmptyError},
		{name: "origin is empty fail", fields: fields{
			Name:      "test",
			Roaster:   "test",
			Origin:    "",
			Price:     1,
			RoastDate: time.Time{},
		}, wantErr: OriginCannotBeEmptyError},
		{name: "price lower than 1 fail", fields: fields{
			Name:      "test",
			Roaster:   "test",
			Origin:    "test",
			Price:     -1,
			RoastDate: time.Time{},
		}, wantErr: PriceCannotBeLowerThan1Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NewCoffeeBean{
				Name:      tt.fields.Name,
				Roaster:   tt.fields.Roaster,
				Origin:    tt.fields.Origin,
				Price:     tt.fields.Price,
				RoastDate: tt.fields.RoastDate,
			}
			err := c.Validate()
			if tt.wantErr == nil && err != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			} else if !errors.Is(err, tt.wantErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCoffeeBean_toDomainModel(t *testing.T) {
	type fields struct {
		Name      string
		Roaster   string
		Origin    string
		Price     int
		RoastDate time.Time
	}
	roastdate := time.Date(3, 3, 3, 3, 3, 3, 3, time.UTC)
	tests := []struct {
		name   string
		fields fields
		want   domain.Bean
	}{
		{name: "success", fields: fields{
			Name:      "bean name",
			Roaster:   "bean roaster",
			Origin:    "bean origin",
			Price:     42,
			RoastDate: roastdate,
		}, want: domain.Bean{Name: "bean name", Roaster: "bean roaster", Origin: "bean origin", Price: 42, RoastDate: roastdate}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NewCoffeeBean{
				Name:      tt.fields.Name,
				Roaster:   tt.fields.Roaster,
				Origin:    tt.fields.Origin,
				Price:     tt.fields.Price,
				RoastDate: tt.fields.RoastDate,
			}
			createdModel := c.toDomainModel()

			if createdModel.Name != tt.fields.Name {
				t.Errorf("toDomainModel() = %v, want %v", createdModel, tt.want)
			}

			if createdModel.Roaster != tt.fields.Roaster {
				t.Errorf("toDomainModel() = %v, want %v", createdModel, tt.want)
			}

			if createdModel.Origin != tt.fields.Origin {
				t.Errorf("toDomainModel() = %v, want %v", createdModel, tt.want)
			}

			if createdModel.RoastDate != tt.fields.RoastDate {
				t.Errorf("toDomainModel() = %v, want %v", createdModel, tt.want)
			}
		})
	}
}
