package api

import (
	"fmt"
	"hello-fresh/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func addRecipesRoutes(r *chi.Mux) {
	r.Route("/recipes", func(r chi.Router) {
		r.Get("/", listRecipes)
		r.Post("/", createRecipe)
		r.Route("/{recipeID}", func(router chi.Router) {
			router.Put("/", updateRecipe)
			router.Delete("/", deleteRecipe)
		})
	})
}

func validateParameterID(r *http.Request) (int, *ErrorResponse) {
	recipeID := chi.URLParam(r, "recipeID")
	if recipeID == "" {
		return 0, ErrorRenderer(fmt.Errorf("recipe ID is required"))
	}
	id, err := strconv.Atoi(recipeID)
	if err != nil {
		return 0, ErrorRenderer(fmt.Errorf("recipe ID must be numeric"))
	}
	return id, nil
}

func updateRecipe(w http.ResponseWriter, r *http.Request) {
	var id int
	var errResponse *ErrorResponse
	if id, errResponse = validateParameterID(r); errResponse != nil {
		render.Render(w, r, errResponse)
		return
	}

	if exist := dbInstance.ExistRecipe(id); !exist {
		render.Render(w, r, ErrNotFound)
		return
	}

	data := &models.RecipeRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	recipe := data.ToDBModel()
	recipe.Id = id

	if err := recipe.Update(dbInstance.DB); err != nil {
		log.Println(err)
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	var id int
	var errResponse *ErrorResponse
	if id, errResponse = validateParameterID(r); errResponse != nil {
		render.Render(w, r, errResponse)
		return
	}
	done, err := dbInstance.RemoveRecipe(id)
	if err != nil {
		log.Println(err)
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if !done {
		render.Render(w, r, ErrNotFound)
	}
	w.WriteHeader(http.StatusOK)
}

func listRecipes(w http.ResponseWriter, r *http.Request) {
	var recipes []render.Renderer
	recipesDB := dbInstance.ListRecipes()
	if len(recipesDB) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	for i := 0; i < len(recipesDB); i++ {
		recipe := &models.Recipe{}
		recipe.From(recipesDB[i])
		recipes = append(recipes, recipe)
	}
	render.RenderList(w, r, recipes)
}

func createRecipe(w http.ResponseWriter, r *http.Request) {
	data := &models.RecipeRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	recipe := data.ToDBModel()
	if err := recipe.Insert(dbInstance.DB); err != nil {
		log.Println(err)
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s/%d", r.RequestURI, recipe.Id))
	w.WriteHeader(http.StatusCreated)
}
