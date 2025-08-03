package web

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"startGo/calculator"
)

type Server struct {
	Config calculator.PriceConfig
	tmpl   *template.Template
}

func NewServer(cfg calculator.PriceConfig) (*Server, error) {
	tmpl, err := template.ParseGlob(filepath.Join("web", "templates", "*.html"))
	if err != nil {
		return nil, err
	}
	return &Server{Config: cfg, tmpl: tmpl}, nil
}

func (s *Server) Start(addr string) error {
	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/calculator", s.handleCalculator)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	s.tmpl.ExecuteTemplate(w, "index.html", nil)
}

func (s *Server) handleCalculator(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err == nil {
			style, _ := strconv.ParseFloat(r.FormValue("style"), 64)
			location, _ := strconv.ParseFloat(r.FormValue("location"), 64)
			color, _ := strconv.ParseFloat(r.FormValue("color"), 64)
			size, _ := strconv.Atoi(r.FormValue("size"))
			price := calculator.EstimateByDetails(s.Config, style, location, color, float64(size))
			s.tmpl.ExecuteTemplate(w, "calculator.html", struct{ Price float64 }{price})
			return
		}
	}
	s.tmpl.ExecuteTemplate(w, "calculator.html", nil)
}
