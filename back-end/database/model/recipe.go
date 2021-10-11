package model

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const RecipeTableName = "recipe"

var recipeSchema = fmt.Sprintf(`
CREATE TABLE %s (
	id serial PRIMARY KEY,
	difficulty text NOT NULL,
	description text NOT NULL,
	image text NOT NULL,
	title text NOT NULL,
	subtitle text NOT NULL,
	cook_time integer NOT NULL,
	tags text[],
	utensils text[],
	allergens text[]
);`, RecipeTableName)

type Recipe struct {
	Id          int            `db:"id"`
	Difficulty  string         `db:"difficulty"`
	Description string         `db:"description"`
	Image       string         `db:"image"`
	Title       string         `db:"title"`
	Substitle   string         `db:"subtitle"`
	CookTime    int            `db:"cook_time"`
	Tags        pq.StringArray `db:"tags"`
	Utensils    pq.StringArray `db:"utensils"`
	Allergens   pq.StringArray `db:"allergens"`
}

func CreateRecipeTable(db *sqlx.DB) {
	if !isTableCreated(db, RecipeTableName) {
		db.MustExec(recipeSchema)
	}
}

func (r *Recipe) Insert(db *sqlx.DB) (err error) {
	tx := db.MustBegin()
	if err = tx.QueryRowx(fmt.Sprintf(`
		INSERT INTO %s (difficulty, tags, allergens, cook_time, utensils, description, image, title, subtitle) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		RETURNING id`, RecipeTableName),
		r.Difficulty, r.Tags, r.Allergens, r.CookTime, r.Utensils, r.Description, r.Image, r.Title, r.Substitle).Scan(&r.Id); err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit()
	return
}

func (r *Recipe) Update(db *sqlx.DB) error {
	sqlStatement := fmt.Sprintf(`
	UPDATE %s
	SET difficulty = $2, tags = $3, allergens= $4, cook_time = $5, utensils = $6, description = $7, image= $8,title= $9,subtitle = $10
	WHERE id = $1;`, RecipeTableName)
	_, err := db.Exec(sqlStatement, r.Id, r.Difficulty, r.Tags, r.Allergens, r.CookTime, r.Utensils, r.Description, r.Image, r.Title, r.Substitle)
	return err
}

func (r *Recipe) Delete(db *sqlx.DB) (bool, error) {
	res, err := db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id=$1", RecipeTableName), r.Id)
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
