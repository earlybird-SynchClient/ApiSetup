package funcmapmaker

import (
	"html/template"
	"net/http"
	"github.com/qor/render"
	"github.com/earlybird-SynchClient/ApiSetup/utils"
	"strings"
	"github.com/earlybird-SynchClient/ApiSetup/models/market"
	"strconv"
)

// AddFuncMapMaker add FuncMapMaker to view
func AddFuncMapMaker(view *render.Render) *render.Render {
	oldFuncMapMaker := view.FuncMapMaker

	view.FuncMapMaker = func(render *render.Render, r *http.Request, w http.ResponseWriter) template.FuncMap {
		funcMap := template.FuncMap{}
		if oldFuncMapMaker != nil {
			funcMap = oldFuncMapMaker(render, r, w)
		}

		funcMap["raw"] = func(str string) template.HTML {
			return template.HTML(utils.HTMLSanitizer.Sanitize(str))
		}

		funcMap["get_frontend"] = func() bool {
			result := strings.Split(r.RequestURI, "/")
			if len(result) >= 3 {
				if strings.Contains(result[1], "auth") || strings.Contains(result[1], "verify"){
					return false;
				}
			}
			return true
		}

		funcMap["get_markets"] = func() []market.Market {
			markets := make([]market.Market, 0)
			utils.GetDB(r).Find(&markets)
			return markets
		}

		funcMap["get_market_index"] = func() uint {
			result := strings.Split(r.RequestURI, "/")
			if len(result) >= 3 {
				if strings.Contains(result[1], "auth") || strings.Contains(result[1], "verify") || strings.Contains(result[1], "admin"){
					return 0;
				} else {
					i, err := strconv.Atoi(result[1])
					if err != nil {
						return 0;
					}
					return uint(i)
				}
			}

			markets := make([]market.Market, 0)
			utils.GetDB(r).Find(&markets)

			if (len(markets) == 0) {
				return 0
			} else {
				return markets[0].ID
			}
		}

		funcMap["get_market_action"] = func() string {
			result := strings.Split(r.RequestURI, "/")
			if len(result) >= 3 {
				if strings.Contains(result[2], "invoice"){
					return "invoice"
				} else if strings.Contains(result[2], "supplier"){
					return "supplier"
				}
			}

			return "invoice"
		}

		return funcMap
	}

	return view
}
