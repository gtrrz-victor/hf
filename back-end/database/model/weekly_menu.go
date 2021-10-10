package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

const weeklyTableName = "weeklymenu"
const recipesMenuTable = "recipesmenu"

var recipesMenuSchema = fmt.Sprintf(`
CREATE TABLE %s (
	week_id int REFERENCES %s (id) ON UPDATE CASCADE ON DELETE CASCADE,
	recipe_id int REFERENCES %s (id) ON UPDATE CASCADE ON DELETE CASCADE,
	 CONSTRAINT recipe_week_pkey PRIMARY KEY (week_id, recipe_id)
);`, recipesMenuTable, weeklyTableName, recipeTableName)

var weeklySchema = fmt.Sprintf(`
CREATE TABLE %s (
	id serial PRIMARY KEY,
	first_day_week DATE NOT NULL,
	num_people integer NOT NULL,
	days_per_week integer NOT NULL
);`, weeklyTableName)

type WeeklyMenu struct {
	Id           int       `db:"id"`
	NumPeople    int       `db:"num_people"`
	DaysPerWeek  int       `db:"days_per_week"`
	FirstDayWeek time.Time `db:"first_day_week"`
}

func CreateRecipesMenuTable(db *sqlx.DB) {
	if !isTableCreated(db, recipesMenuTable) {
		db.MustExec(recipesMenuSchema)
	}
}

func CreateWeeklyMenuTable(db *sqlx.DB) {
	if !isTableCreated(db, weeklyTableName) {
		db.MustExec(weeklySchema)
	}
}

func (w *WeeklyMenu) Insert(db *sqlx.DB) (err error) {
	tx := db.MustBegin()
	if err = tx.QueryRowx(fmt.Sprintf(`
		INSERT INTO %s (num_people, days_per_week, first_day_week) 
		VALUES ($1, $2, $3) 
		RETURNING id`, weeklyTableName), w.NumPeople, w.DaysPerWeek, w.FirstDayWeek).Scan(&w.Id); err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit()
	return
}

func (w *WeeklyMenu) AddRecipe(recipe *Recipe, db *sqlx.DB) (err error) {
	tx := db.MustBegin()
	if _, err = tx.Exec(fmt.Sprintf("INSERT INTO %s (week_id, recipe_id) VALUES ($1, $2)", recipesMenuTable), w.Id, recipe.Id); err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit()
	return
}

func (w *WeeklyMenu) Delete(db *sqlx.DB) error {
	res, err := db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id=$1", weeklyTableName), w.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	return err
}

func (w *WeeklyMenu) Update(db *sqlx.DB) error {
	sqlStatement := fmt.Sprintf(`
	UPDATE %s
	SET num_people = $2, days_per_week = $3, first_day_week = $4
	WHERE id = $1;`, weeklyTableName)
	_, err := db.Exec(sqlStatement, w.Id, w.NumPeople, w.DaysPerWeek, w.FirstDayWeek)
	return err
}

func (w *WeeklyMenu) FindRecipes(db *sqlx.DB) (*[]Recipe, error) {
	recipes := []Recipe{}
	sql := strings.ReplaceAll(`
	SELECT
		$recipeTable.*
    FROM
		$recipeTable JOIN $recipesMenu ON $recipesMenu.recipe_id = $recipeTable.id 
		JOIN $weeklyMenu ON $recipesMenu.week_id = $weeklyMenu.id
	WHERE
		$weeklyMenu.id = $1`, "$recipeTable", recipeTableName)
	sql = strings.ReplaceAll(sql, "$recipesMenu", recipesMenuTable)
	sql = strings.ReplaceAll(sql, "$weeklyMenu", weeklyTableName)
	err := db.Select(&recipes, sql, w.Id)
	return &recipes, err
}
