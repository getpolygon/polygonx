// BSD 3-Clause License

// Copyright (c) 2021, Michael Grigoryan
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.

// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.

// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package wellknown

import (
	"net/http"

	"github.com/getpolygon/corexp/internal/deps"
	"github.com/getpolygon/corexp/internal/gen/postgres_codegen"
	"github.com/getpolygon/corexp/internal/settings"
	"github.com/getpolygon/corexp/internal/types"
	"github.com/go-chi/render"
)

func genMetadata(s *settings.Settings) map[string]interface{} {
	return map[string]interface{}{
		"usage": map[string]interface{}{
			"private": !s.Security.ReportInstanceStats,
		},
	}
}

// Constant version specifier for the NodeInfo specification.
// http://nodeinfo.diaspora.software/schema.html
const NodeInfoVersion string = "2.1"
const NodeInfoSoftwareName string = "Polygon"
const NodeInfoSoftwareVersion string = "0.2.0"
const NodeInfoSoftwareHomepage string = "https://polygon.am/"
const NodeInfoSoftwareRepository string = "https://github.com/getpolygon"

// This route will provide publicly accessible information about current Polygon instance,
// including instance configuration, total users, comments and posts count etc. By default
// fetching instance usage, including users, posts and comment count without personal info
// is disabled and has to be specified manually in the configuration by the user.
func NodeInformation(deps *deps.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var stats postgres_codegen.GetFullUsageStatsRow

		// Only fetching instance's statistics with user's consent. Usage
		// statistics will mostly be blank, and will contain `0`s because
		// of this.
		if deps.Settings.Security.ReportInstanceStats {
			var err error

			stats, err = deps.Postgres.GetFullUsageStats(r.Context())
			if err != nil {
				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r, err)
				return
			}
		}

		nodeinfo := types.NodeInfo{
			Version:           NodeInfoVersion,
			OpenRegistrations: deps.Settings.Security.OpenRegistrations,

			Usage: types.NodeInfoUsage{
				LocalPosts:    stats.PostsCount,
				LocalComments: stats.CommentsCount,
				Users: types.NodeInfoUsers{
					Total:          stats.UsersCount,
					ActiveOneMonth: stats.ActiveUsersMonth,
					ActiveHalfYear: stats.ActiveUsersHalfYear,
				},
			},

			Software: types.NodeInfoSoftware{
				Name:       NodeInfoSoftwareName,
				Version:    NodeInfoSoftwareVersion,
				Homepage:   NodeInfoSoftwareHomepage,
				Repository: NodeInfoSoftwareRepository,
			},

			Metadata:  genMetadata(deps.Settings),
			Services:  make([]types.NodeInfoService, 0),
			Protocols: make([]types.NodeInfoProtocol, 0),
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, nodeinfo)
		return
	}
}
