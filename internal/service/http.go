package service

import (
    "encoding/json"
    "errors"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/Gorynychdo/bstreee/internal/bstree"
    log "github.com/sirupsen/logrus"
)

var errIncorrectValue = errors.New("incorrect value")

// responseWriter WriteHeader implementer for logging status code in middleware
type responseWriter struct {
    http.ResponseWriter
    code int
}

func (w *responseWriter) WriteHeader(statusCode int) {
    w.code = statusCode
    w.ResponseWriter.WriteHeader(statusCode)
}

type server struct {
    tree   *bstree.Tree
    router *mux.Router
}

// NewServer http server constructor
func NewServer(tree *bstree.Tree) *server {
    s := &server{
        tree:   tree,
        router: mux.NewRouter(),
    }
    s.configureRouter()
    return s
}

// Serve http.ListenAndServe implementation
func (s *server) Serve(addr string) error {
    log.WithFields(log.Fields{
        "package": "http",
    }).Info("start http server")
    return http.ListenAndServe(addr, s.router)
}

func (s *server) configureRouter() {
    s.router.Use(s.logMiddleware)
    s.router.HandleFunc("/search", s.searchValue).Methods("GET")
    s.router.HandleFunc("/insert", s.insertValue).Methods("POST")
    s.router.HandleFunc("/delete", s.deleteValue).Methods("DELETE")
}

func (s *server) logMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        rw := &responseWriter{w, http.StatusOK}
        next.ServeHTTP(rw, r)
        log.WithFields(log.Fields{
            "package": "http",
            "method":  r.Method,
            "path":    r.URL.Path,
            "status":  rw.code,
        }).Info()
    })
}

func (s *server) searchValue(w http.ResponseWriter, r *http.Request) {
    data := r.URL.Query().Get("val")
    value, err := strconv.Atoi(data)
    if err != nil {
        s.error(w, http.StatusBadRequest, errIncorrectValue)
        return
    }

    found := s.tree.Search(value)
    s.respond(w, http.StatusOK, map[string]bool{"found": found})
}

func (s *server) insertValue(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Value int `json:"val"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        s.error(w, http.StatusBadRequest, errIncorrectValue)
        return
    }
    _ = r.Body.Close()
    s.tree.Insert(req.Value)
    s.respond(w, http.StatusNoContent, nil)
}

func (s *server) deleteValue(w http.ResponseWriter, r *http.Request) {
    data := r.URL.Query().Get("val")
    value, err := strconv.Atoi(data)
    if err != nil {
        s.error(w, http.StatusBadRequest, errIncorrectValue)
        return
    }

    s.tree.Delete(value)
    s.respond(w, http.StatusNoContent, nil)
}

func (s *server) error(w http.ResponseWriter, code int, err error) {
    s.respond(w, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, code int, data interface{}) {
    w.WriteHeader(code)
    if data == nil {
        return 
    }
    if err := json.NewEncoder(w).Encode(data); err != nil {
        log.WithFields(log.Fields{
            "package": "http",
        }).Error(err)
    }
}
