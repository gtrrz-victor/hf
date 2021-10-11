package models

import (
	"hello-fresh/database/model"
	"net/http"
)

type RecipeRequest struct {
	Difficulty  string   `json:"difficulty"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Title       string   `json:"title"`
	Substitle   string   `json:"subtitle"`
	CookTime    int      `json:"cook_time"`
	Tags        []string `json:"tags"`
	Utensils    []string `json:"utensils"`
	Allergens   []string `json:"allergens"`
}

type Recipe struct {
	*RecipeRequest
	Id int `json:"id"`
}

func (*Recipe) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (r *Recipe) From(recipeDB model.Recipe) {
	r.Id = recipeDB.Id
	r.RecipeRequest = &RecipeRequest{}
	r.RecipeRequest.Difficulty = recipeDB.Difficulty
	r.RecipeRequest.Description = recipeDB.Description
	r.RecipeRequest.Image = recipeDB.Image
	r.RecipeRequest.Title = recipeDB.Title
	r.RecipeRequest.Substitle = recipeDB.Substitle
	r.RecipeRequest.CookTime = recipeDB.CookTime
	r.RecipeRequest.Tags = recipeDB.Tags
	r.RecipeRequest.Utensils = recipeDB.Utensils
	r.RecipeRequest.Allergens = recipeDB.Allergens
}

func (u *RecipeRequest) Bind(r *http.Request) error {
	return nil
}

func (r *RecipeRequest) ToDBModel() *model.Recipe {
	recipeDB := &model.Recipe{}
	recipeDB.Difficulty = r.Difficulty
	recipeDB.Description = r.Description
	recipeDB.Image = r.Image
	recipeDB.Title = r.Title
	recipeDB.Substitle = r.Substitle
	recipeDB.CookTime = r.CookTime
	recipeDB.Tags = r.Tags
	recipeDB.Utensils = r.Utensils
	recipeDB.Allergens = r.Allergens
	return recipeDB
}
