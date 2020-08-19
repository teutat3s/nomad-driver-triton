package plugin

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/teutat3s/nomad-driver-triton/types"
)

// firstValidAuth tries a list of auth backends, returning first error or AuthConfiguration
func firstValidAuth(repo string, backends []authBackend) (*docker.AuthConfiguration, error) {
	for _, backend := range backends {
		auth, err := backend(repo)
		if auth != nil || err != nil {
			return auth, err
		}
	}
	return nil, nil
}

// authFromTaskConfig generates an authBackend for any auth given in the task-configuration
func authFromTaskConfig(tc types.TaskConfig) authBackend {
	return func(string) (*docker.AuthConfiguration, error) {
		// If all auth fields are empty, return
		if len(tc.Auth.Username) == 0 && len(tc.Auth.Password) == 0 && len(tc.Auth.Email) == 0 && len(tc.Auth.ServerAddr) == 0 {
			return nil, nil
		}
		return &docker.AuthConfiguration{
			Username:      tc.Auth.Username,
			Password:      tc.Auth.Password,
			Email:         tc.Auth.Email,
			ServerAddress: tc.Auth.ServerAddr,
		}, nil
	}
}

// authIsEmpty returns if auth is nil or an empty structure
func authIsEmpty(auth *docker.AuthConfiguration) bool {
	if auth == nil {
		return false
	}
	return auth.Username == "" &&
		auth.Password == "" &&
		auth.Email == "" &&
		auth.ServerAddress == ""
}
