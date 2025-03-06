package handlers

import (
  "html/template"
  "net/http"
  "log"

  "github.com/chiltom/pogo_buddy/internal/models"
  "github.com/chiltom/pogo_buddy/internal/services"
)

type UserHandlers struct {
  userSvc *services.UserService
  tmpl    *template.Template
}

func (h *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

  if err := r.ParseForm(); err != nil {
    http.Error(w, "Bad Request", http.StatusBadRequest)
    return
  }

  user := &models.User{
    Email:     r.FormValue("email"),
    Password:  r.FormValue("password"),
    FirstName: r.FormValue("first_name"),
    LastName:  r.FormValue("last_name"),
  }

  id, err := h.userSvc.Create(user)
  if err != nil {
    log.Printf("failed to create user: %v", err)
    http.Error(w, "Failed to create user", http.StatusInternalServerError)
    return
  }

  data := struct{
    UserID int
  }{id}
  if err := h.tmpl.ExecuteTemplate(w, "uc_success.html", data); err != nil {
    log.Printf("failed to render template: %v", err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
  }
}
