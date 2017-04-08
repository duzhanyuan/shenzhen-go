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

// Package api has types for communicating with the UI.
package api

// Interface is "the API" - the interface defining available requests.
type Interface interface {
	SetPosition(*SetPositionRequest) error
}

// Empty is just an empty message.
type Empty struct{}

// Request is the embedded base of all requests.
type Request struct {
	Graph string `json:"graph"`
}

// ChannelRequest is the embedded base of all requests to do with channels.
type ChannelRequest struct {
	Request
	Channel string `json:"channel"`
}

// NodeRequest is the embedded base of all requests to do with nodes.
type NodeRequest struct {
	Request
	Node string `json:"node"`
}

// SetPositionRequest is a request to change the position of a node.
type SetPositionRequest struct {
	NodeRequest
	X int `json:"x"`
	Y int `json:"y"`
}