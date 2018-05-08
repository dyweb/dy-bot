// Copyright 2018 The Dongyue members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/dyweb/dy-bot/cli/dy-bot/server/config"
	"github.com/dyweb/dy-bot/pkg/event"
	"github.com/dyweb/gommon/util/logutil"
)

// DefaultAddress is the default address daemon will listen to.
const DefaultAddress = ":6789"

var log = logutil.NewPackageLogger()

type Server struct {
	config config.Config
	// manager processes webhook event from GitHub.
	manager *event.Manager
}

func NewServer(cfg config.Config) *Server {
	return &Server{
		config:  cfg,
		manager: event.NewManager(),
	}
}

func (s *Server) Run() error {
	// start webserver
	listenAddress := s.config.HTTPListen
	if listenAddress == "" {
		listenAddress = DefaultAddress
	}

	r := mux.NewRouter()

	// register ping api
	r.HandleFunc("/ping", pingHandler).Methods("GET")

	// github webhook API
	r.HandleFunc("/events", s.gitHubEventHandler).Methods("POST")

	log.Infof("Listening on %s for the repository %s/%s", listenAddress, s.config.Owner, s.config.Repo)
	return http.ListenAndServe(listenAddress, r)
}

// gitHubEventHandler handles webhook events from github.
func (s *Server) gitHubEventHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("/events request received")
	eventType := r.Header.Get("X-Github-Event")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	r.Body.Close()

	if err := s.manager.HandleEvent(eventType, data); err != nil {
		log.Errorf("Errored when handle webhook events: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

// pingHandler handles ping request to return health of server.
func pingHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("/ping request received")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{'O', 'K'})
	return
}
