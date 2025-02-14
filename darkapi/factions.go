package darkapi

import (
	"github.com/darklab8/fl-darkstat/darkcore/web"
	"github.com/darklab8/fl-darkstat/darkcore/web/registry"
	"github.com/darklab8/fl-darkstat/darkstat/configs_export"
)

// ShowAccount godoc
// @Summary      Getting list of Factions
// @Tags         factions
// @Accept       json
// @Produce      json
// @Success      200  {array}  	configs_export.Faction
// @Router       /api/factions [get]
// @Param        filter_to_useful    query     string  false  "filter items only to useful, usually they are sold, or have goods, or craftable or findable in loot, or bases that are flight reachable from manhattan"
func GetFactions(webapp *web.Web, api *Api) *registry.Endpoint {
	return &registry.Endpoint{
		Url:     "GET " + ApiRoute + "/factions",
		Handler: GetItemsT(webapp, api.app_data.Configs.Factions, configs_export.FilterToUsefulFactions),
	}
}
