package database

import (
	"fmt"
	"hello-fresh/database/model"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	HOST = "database"
	PORT = 5432
)

var connection Conn

type Conn struct {
	DB *sqlx.DB
}

func Initialize() Conn {
	connection = Conn{}
	var err error
	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, dbUser, dbPassword, dbName)

	connection.DB, err = sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = connection.DB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Database connection established")
	model.CreateRecipeTable(connection.DB)
	model.CreateWeeklyMenuTable(connection.DB)
	model.CreateRecipesMenuTable(connection.DB)
	return connection
}

func (c Conn) ListRecipes() []model.Recipe {
	recipes := []model.Recipe{}
	if err := c.DB.Select(&recipes, fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", model.RecipeTableName)); err != nil {
		log.Println(err)
	}
	return recipes
}

func (c Conn) ListWeeklyMenu() []model.WeeklyMenu {
	weeklyMenus := []model.WeeklyMenu{}
	if err := c.DB.Select(&weeklyMenus, fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", model.WeeklyTableName)); err != nil {
		log.Println(err)
	}
	return weeklyMenus
}

func (c Conn) ListWeekRecipes(id int) []model.Recipe {
	weeklyMenu := &model.WeeklyMenu{Id: id}
	recipes, err := weeklyMenu.FindRecipes(c.DB)
	if err != nil {
		log.Println(err)
		return []model.Recipe{}
	}
	return *recipes
}

func (c Conn) RemoveWeeklyMenu(id int) (bool, error) {
	weeklyMenu := &model.WeeklyMenu{Id: id}
	return weeklyMenu.Delete(c.DB)
}
func (c Conn) DeleteWeekRecipes(id int) error {
	weeklyMenu := &model.WeeklyMenu{Id: id}
	return weeklyMenu.DeleteRecipes(c.DB)
}

func (c Conn) ReplaceRecipesToWeek(recipes []int, weekID int) error {
	if err := c.ExistRecipes(recipes); err != nil {
		return err
	}
	if err := c.DeleteWeekRecipes(weekID); err != nil {
		return err
	}
	return c.AddRecipesToWeek(recipes, weekID)
}

func (c Conn) AddRecipesToWeek(recipes []int, weekID int) error {
	if err := c.ExistRecipes(recipes); err != nil {
		return err
	}
	weeklyMenu := &model.WeeklyMenu{Id: weekID}
	alreadyLinkedRecipes, err := weeklyMenu.FindRecipesId(c.DB)
	if err != nil {
		return err
	}

	tx := c.DB.MustBegin()
	for _, recipeID := range recipes {
		found := false
		for _, alreadyLinkedRecipeId := range alreadyLinkedRecipes {
			found = found || alreadyLinkedRecipeId == recipeID
		}
		if !found {
			if _, err := tx.Exec(fmt.Sprintf("INSERT INTO %s (week_id, recipe_id) VALUES ($1, $2)", model.RecipesMenuTable), weekID, recipeID); err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit()
}

func (c Conn) ExistRecipes(ids []int) error {
	var recipesNotFound []int
	for _, recipeId := range ids {
		if !c.ExistRecipe(recipeId) {
			recipesNotFound = append(recipesNotFound, recipeId)
		}
	}
	if len(recipesNotFound) > 0 {
		return fmt.Errorf("recipes with id %d not found", recipesNotFound)
	}
	return nil
}

func (c Conn) ExistWeeklyMenu(id int) bool {
	weeklyMenu := &model.WeeklyMenu{}
	if err := c.DB.Get(weeklyMenu, fmt.Sprintf("SELECT * FROM %s WHERE id=$1", model.WeeklyTableName), id); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (c Conn) RemoveRecipe(id int) (bool, error) {
	recipe := &model.Recipe{Id: id}
	return recipe.Delete(c.DB)
}

func (c Conn) ExistRecipe(id int) bool {
	recipe := &model.Recipe{}
	if err := c.DB.Get(recipe, fmt.Sprintf("SELECT * FROM %s WHERE id=$1", model.RecipeTableName), id); err != nil {
		log.Println(err)
		return false
	}
	return true
}
