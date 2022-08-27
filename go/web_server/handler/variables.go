package handler

import "text/template"

// Model stores the models that will be rendered
var Model = template.Must(template.ParseFiles("templates/index.html"))
