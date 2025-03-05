package handlers

import (
  "html/template"
  "net/http"

  "github.com/chiltom/pogo_buddy/internal/db"
  "github.com/chiltom/pogo_buddy/internal/services"
)

type Handlers struct {
  userSvc *services.UserService
  tmpl    *template.Template
}

