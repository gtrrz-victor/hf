package api

import (
	"fmt"
	"hello-fresh/database/model"
	"hello-fresh/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func addWeeklyMenuRoutes(r *chi.Mux) {
	r.Route("/weeklyMenus", func(r chi.Router) {
		r.Get("/", listWeeklyMenus)
		r.Post("/", createWeeklyMenu)
		r.Route("/{weeklyMenuID}", func(router chi.Router) {
			router.Put("/", updateWeeklyMenu)
			router.Delete("/", deleteWeeklyMenu)
			router.Route("/recipes", func(r chi.Router) {
				r.Get("/", listMenuRecipes)
				r.Patch("/", patchWeekRecipes)
				r.Put("/", updateWeekRecipes)
			})
		})
	})
}

func validateParameterWeekID(r *http.Request) (int, *ErrorResponse) {
	weeklyMenuID := chi.URLParam(r, "weeklyMenuID")
	if weeklyMenuID == "" {
		return 0, ErrorRenderer(fmt.Errorf("weeklyMenu ID is required"))
	}
	id, err := strconv.Atoi(weeklyMenuID)
	if err != nil {
		return 0, ErrorRenderer(fmt.Errorf("weeklyMenu ID must be numeric"))
	}
	return id, nil
}

func updateWeekRecipes(w http.ResponseWriter, r *http.Request) {
	var id int
	var errResponse *ErrorResponse
	if id, errResponse = validateParameterWeekID(r); errResponse != nil {
		render.Render(w, r, errResponse)
		return
	}

	if exist := dbInstance.ExistWeeklyMenu(id); !exist {
		render.Render(w, r, ErrNotFound)
		return
	}
	data := &models.WeekRecipes{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := dbInstance.ReplaceRecipesToWeek(*data, id); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func patchWeekRecipes(w http.ResponseWriter, r *http.Request) {
	var id int
	var errResponse *ErrorResponse
	if id, errResponse = validateParameterWeekID(r); errResponse != nil {
		render.Render(w, r, errResponse)
		return
	}
	if exist := dbInstance.ExistWeeklyMenu(id); !exist {
		render.Render(w, r, ErrNotFound)
		return
	}
	data := &models.WeekRecipes{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := dbInstance.AddRecipesToWeek(*data, id); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func listMenuRecipes(w http.ResponseWriter, r *http.Request) {
	var recipes []render.Renderer
	var id int
	var errResponse *ErrorResponse
	if id, errResponse = validateParameterWeekID(r); errResponse != nil {
		render.Render(w, r, errResponse)
		return
	}
	if exist := dbInstance.ExistWeeklyMenu(id); !exist {
		render.Render(w, r, ErrNotFound)
		return
	}
	recipesDB := dbInstance.ListWeekRecipes(id)
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

func updateWeeklyMenu(w http.ResponseWriter, r *http.Request) {
	var id int
	var errResponse *ErrorResponse
	if id, errResponse = validateParameterWeekID(r); errResponse != nil {
		render.Render(w, r, errResponse)
		return
	}

	if exist := dbInstance.ExistWeeklyMenu(id); !exist {
		render.Render(w, r, ErrNotFound)
		return
	}

	data := &models.WeeklyMenuRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	weeklyMenu, err := data.ToDBModel()
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	weeklyMenu.Id = id

	if err := weeklyMenu.Update(dbInstance.DB); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteWeeklyMenu(w http.ResponseWriter, r *http.Request) {
	var id int
	var errResponse *ErrorResponse
	if id, errResponse = validateParameterWeekID(r); errResponse != nil {
		render.Render(w, r, errResponse)
		return
	}
	done, err := dbInstance.RemoveWeeklyMenu(id)
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

func createWeeklyMenu(w http.ResponseWriter, r *http.Request) {
	var weeklyMenu *model.WeeklyMenu
	var err error
	data := &models.WeeklyMenuRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	weeklyMenu, err = data.ToDBModel()
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := weeklyMenu.Insert(dbInstance.DB); err != nil {
		log.Println(err)
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s/%d", r.RequestURI, weeklyMenu.Id))
	w.WriteHeader(http.StatusCreated)
}

func listWeeklyMenus(w http.ResponseWriter, r *http.Request) {
	var weeklyMenus []render.Renderer
	weeklyMenusDB := dbInstance.ListWeeklyMenu()
	if len(weeklyMenusDB) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	for i := 0; i < len(weeklyMenusDB); i++ {
		weeklyMenu := &models.WeeklyMenu{}
		weeklyMenu.From(weeklyMenusDB[i])
		weeklyMenus = append(weeklyMenus, weeklyMenu)
	}
	render.RenderList(w, r, weeklyMenus)
}
