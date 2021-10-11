package models

import (
	"fmt"
	"hello-fresh/database/model"
	"net/http"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

// First create a type alias

type WeekRecipes []int

type WeeklyMenuRequest struct {
	NumPeople    int    `json:"numPeople"`
	DaysPerWeek  int    `json:"daysPerWeek"`
	FirstDayWeek string `json:"firstDayWeek"`
}

type WeeklyMenu struct {
	*WeeklyMenuRequest
	Id int `json:"id"`
}

func (*WeekRecipes) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*WeekRecipes) Bind(r *http.Request) error {
	return nil
}

func (*WeeklyMenu) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*WeeklyMenuRequest) Bind(r *http.Request) error {
	return nil
}

func (r *WeeklyMenu) From(weeklyMenuDB model.WeeklyMenu) {
	r.Id = weeklyMenuDB.Id
	r.WeeklyMenuRequest = &WeeklyMenuRequest{}
	r.WeeklyMenuRequest.NumPeople = weeklyMenuDB.NumPeople
	r.WeeklyMenuRequest.DaysPerWeek = weeklyMenuDB.DaysPerWeek
	r.WeeklyMenuRequest.FirstDayWeek = weeklyMenuDB.FirstDayWeek.Format(layoutISO)
}

func (r *WeeklyMenuRequest) ToDBModel() (*model.WeeklyMenu, error) {
	var err error
	if err = r.Validate(); err != nil {
		return nil, err
	}
	weeklyMenuDB := &model.WeeklyMenu{}
	weeklyMenuDB.NumPeople = r.NumPeople
	weeklyMenuDB.DaysPerWeek = r.DaysPerWeek
	weeklyMenuDB.FirstDayWeek, _ = time.Parse(layoutISO, r.FirstDayWeek)
	return weeklyMenuDB, err
}

func (r *WeeklyMenuRequest) Validate() error {
	firstDay, err := time.Parse(layoutISO, r.FirstDayWeek)
	if err != nil {
		return err
	}
	if firstDay.Weekday() != time.Monday {
		return fmt.Errorf("invalid time %s, first day week must be Monday", r.FirstDayWeek)
	}
	if r.NumPeople != 2 && r.NumPeople != 4 {
		return fmt.Errorf("numPeople field must be 2 or 4")
	}
	if r.DaysPerWeek != 3 && r.DaysPerWeek != 4 && r.DaysPerWeek != 5 {
		return fmt.Errorf("daysPerWeek field must be 3, 4 or 5")
	}
	return nil
}
