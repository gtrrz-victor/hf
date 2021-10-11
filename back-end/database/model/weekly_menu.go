package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

const WeeklyTableName = "weeklymenu"
const RecipesMenuTable = "recipesmenu"

var recipesMenuSchema = fmt.Sprintf(`
CREATE TABLE %s (
	week_id int REFERENCES %s (id) ON UPDATE CASCADE ON DELETE CASCADE,
	recipe_id int REFERENCES %s (id) ON UPDATE CASCADE ON DELETE CASCADE,
	 CONSTRAINT recipe_week_pkey PRIMARY KEY (week_id, recipe_id)
);`, RecipesMenuTable, WeeklyTableName, RecipeTableName)

var weeklySchema = fmt.Sprintf(`
CREATE TABLE %s (
	id serial PRIMARY KEY,
	first_day_week DATE NOT NULL,
	num_people integer NOT NULL,
	days_per_week integer NOT NULL
);`, WeeklyTableName)

type WeeklyMenu struct {
	Id           int       `db:"id"`
	NumPeople    int       `db:"num_people"`
	DaysPerWeek  int       `db:"days_per_week"`
	FirstDayWeek time.Time `db:"first_day_week"`
}

func CreateRecipesMenuTable(db *sqlx.DB) {
	if !isTableCreated(db, RecipesMenuTable) {
		db.MustExec(recipesMenuSchema)
	}
}

func CreateWeeklyMenuTable(db *sqlx.DB) {
	if !isTableCreated(db, WeeklyTableName) {
		db.MustExec(weeklySchema)
	}
}

func (w *WeeklyMenu) Insert(db *sqlx.DB) (err error) {
	tx := db.MustBegin()
	if err = tx.QueryRowx(fmt.Sprintf(`
		INSERT INTO %s (num_people, days_per_week, first_day_week) 
		VALUES ($1, $2, $3) 
		RETURNING id`, WeeklyTableName), w.NumPeople, w.DaysPerWeek, w.FirstDayWeek).Scan(&w.Id); err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit()
	return
}

func (w *WeeklyMenu) Delete(db *sqlx.DB) (bool, error) {
	res, err := db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id=$1", WeeklyTableName), w.Id)
	if err != nil {
		return false, err
	}
	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsDeleted == 0 {
		return false, err
	}
	return true, err
}
func (w *WeeklyMenu) DeleteRecipes(db *sqlx.DB) error {
	if _, err := db.Exec(fmt.Sprintf("DELETE FROM %s WHERE week_id=$1", RecipesMenuTable), w.Id); err != nil {
		return err
	}
	return nil
}

func (w *WeeklyMenu) Update(db *sqlx.DB) error {
	sqlStatement := fmt.Sprintf(`
	UPDATE %s
	SET num_people = $2, days_per_week = $3, first_day_week = $4
	WHERE id = $1;`, WeeklyTableName)
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
		$weeklyMenu.id = $1`, "$recipeTable", RecipeTableName)
	sql = strings.ReplaceAll(sql, "$recipesMenu", RecipesMenuTable)
	sql = strings.ReplaceAll(sql, "$weeklyMenu", WeeklyTableName)
	err := db.Select(&recipes, sql, w.Id)
	return &recipes, err
}

func (w *WeeklyMenu) FindRecipesId(db *sqlx.DB) ([]int, error) {
	recipesId := []int{}
	sql := strings.ReplaceAll(`
	SELECT
		$recipesMenu.recipe_id
    FROM
		$weeklyMenu JOIN $recipesMenu ON $recipesMenu.week_id = $weeklyMenu.id 
	WHERE
		$weeklyMenu.id = $1`, "$recipeTable", RecipeTableName)
	sql = strings.ReplaceAll(sql, "$recipesMenu", RecipesMenuTable)
	sql = strings.ReplaceAll(sql, "$weeklyMenu", WeeklyTableName)
	err := db.Select(&recipesId, sql, w.Id)
	return recipesId, err
}
