package jsondb

import "github.com/jhead/phantom/internal/services/model"

// Functions for manipulating the model.Server model in the JSON database

func (database Database) ListServers() (map[string]model.Server, error) {
	contents, err := database.readJSON()

	if err != nil {
		return nil, nil
	}

	updatedServers := make(map[string]model.Server, len(contents.Servers))
	for id, server := range contents.Servers {
		// ID isn't stored in the object since it's already the map key
		server.ID = id
		updatedServers[id] = server
	}

	return updatedServers, nil
}

func (database Database) GetServer(id string) (model.Server, error) {
	servers, err := database.ListServers()

	if err != nil {
		return model.Server{}, err
	}

	if server, exists := servers[id]; !exists {
		return model.Server{}, model.ServerNotFoundError
	} else {
		return server, nil
	}
}

func (database Database) CreateServer(server model.Server) error {
	contents, err := database.readJSON()

	if err != nil {
		panic(err)
		// return err
	}

	if _, exists := contents.Servers[server.ID]; exists {
		return model.ServerExistsError
	}

	contents.Servers[server.ID] = server

	return database.writeJSON(contents)
}

func (database Database) UpdateServer(server model.Server) error {
	return nil
}

func (database Database) DeleteServer(id string) error {
	return nil
}
