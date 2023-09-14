package corest

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

type Response[T any] struct {
	Status  bool   `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

func JsonResponse[T any](writer http.ResponseWriter, data Response[T]) {
	writer.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(data)
	if err != nil {
		return
	}
	_, err = writer.Write(res)
	if err != nil {
		return
	}
}

type HttpServer struct {
	port     string
	endpoint string
	router   *chi.Mux
}

func New(port string, endpoint string) *HttpServer {
	router := chi.NewRouter()
	router.Use(useCors().Handler)
	return &HttpServer{
		port:     port,
		endpoint: endpoint,
		router:   router,
	}
}

func (h *HttpServer) AddController(route func(r chi.Router)) {
	h.router.Route(h.endpoint, route)
}

func (h *HttpServer) Start() {
	fmt.Printf("API listen on :%s\n", h.port)
	err := http.ListenAndServe(":"+h.port, h.router)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func useCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
