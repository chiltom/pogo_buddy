package handlers

import (
  "html/template"
  "github.com/chiltom/pogo_buddy/internal/db"
  "github.com/chiltom/pogo_buddy/internal/services"
)

// Can add more handler types as needed
type Handlers struct {
  User *UserHandlers
  tmpl *template.Template
  db *db.DB
}

func New(dbConn *db.DB) *Handlers {
  tmpl := template.Must(template.ParseGlob("static/html/*.html"))
  
  userStore := db.NewUserStore(dbConn.DB)
  userSvc := services.NewUserService(userStore)

  userHandlers := &UserHandlers {
    userSvc: userSvc,
    tmpl: tmpl,
  }

  return &Handlers {
    User: userHandlers,
    tmpl: tmpl,
    db: dbConn,
  }
}
