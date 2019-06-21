// Author: James Mallon <jamesmallondev@gmail.com>
// layout package - package offers functions to deal with templates and layouts
package layout

import (
	log "axeman/libs/logger"
	"html/template"
	"net/http"
)

// Renderer function -
func Renderer(res http.ResponseWriter, layout string, pageData interface{}, views ...string) {
	tpl := template.Must(template.ParseFiles(views...))
	err := tpl.ExecuteTemplate(res, layout, pageData)
	if err != nil {
		log.SLog.WriteLog("error", err.Error(), log.SLog.GetTraceMsg())
	}
}
