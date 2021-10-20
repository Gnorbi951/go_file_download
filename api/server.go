package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Server struct {
	*mux.Router

	items []Item
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		items:  []Item{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/download", s.download()).Methods("GET")
}

func (s *Server) download() http.HandlerFunc {

	testItem := Item{ID: uuid.New(), Name: "TestItem"}
	fmt.Println(testItem)

	return func(w http.ResponseWriter, r *http.Request) {
		// Open csv
		fileName := "kek.csv"
		file, err := os.Open(fileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Serve csv
		w.Header().Set("Content-Disposition", "attachment; filename=test.csv")
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		// http.ServeContent(w, r, fileName, time.Now(), file)
	}
}
