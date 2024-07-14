package app

import (
	"awesomeProject/internal/controller"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/service"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx"
	"log"
	"net/http"
)

type App struct {
	c controller.Controller
	//config
}

func (a *App) initDeps() {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "default",
	})
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(conn)
	service := service.NewService(repo)
	a.c = controller.NewController(service)

}
func (a *App) Run() error {
	a.initDeps()
	r := chi.NewRouter()
	r.Group(
		func(r chi.Router) {
			r.Get("/user", a.c.GetUser)
			r.Post("/user", a.c.CreateUser)
			r.Delete("/user", a.c.DeleteUser)
			r.Put("/user", a.c.UpdateUser)
		})
	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}
	return nil
}
