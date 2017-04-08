// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"log"
	"net/http"

	"github.com/google/shenzhen-go/api"
	"github.com/google/shenzhen-go/model"
)

type apiHandler struct{}

// API handles all API requests.
var API apiHandler

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s api: %s", r.Method, r.URL)
	api.Dispatch(h, w, r)
}

func lookupGraph(graph string) (*model.Graph, error) {
	g := loadedGraphs[graph]
	if g == nil {
		return nil, api.Statusf(http.StatusNotFound, "graph not loaded: %q", graph)
	}
	return g, nil
}

func lookupChannel(graph, channel string) (*model.Graph, *model.Channel, error) {
	g, err := lookupGraph(graph)
	if err != nil {
		return nil, nil, err
	}
	c := g.Channels[channel]
	if c == nil {
		return nil, nil, api.Statusf(http.StatusNotFound, "no such channel: %q", channel)
	}
	return g, c, nil
}

func lookupNode(graph, node string) (*model.Graph, *model.Node, error) {
	g, err := lookupGraph(graph)
	if err != nil {
		return nil, nil, err
	}
	n := g.Nodes[node]
	if n == nil {
		return nil, nil, api.Statusf(http.StatusNotFound, "no such node: %q", node)
	}
	return g, n, nil
}

func (h apiHandler) SetPosition(req *api.SetPositionRequest) error {
	_, n, err := lookupNode(req.Graph, req.Node)
	if err != nil {
		return err
	}
	n.X, n.Y = req.X, req.Y
	return nil
}