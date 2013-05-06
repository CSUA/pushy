package main 

type Configuration struct {
	Port         int                `json:"port"`
	User         string             `json:"user"`
	Group        string             `json:"group"`
	Repositories []RepositoryConfig `json:"repositories"`
}

type RepositoryConfig struct {
	Url          string   `json:"url"`
	Path         string   `json:"path"`
	Branch       string   `json:"branch"`
	PreCommands  []string `json:"precommands"`
	PostCommands []string `json:"postcommands"`
}

func (config *Configuration) FindRepositoryConfig(repository Repository) (repoConfig *RepositoryConfig) {
	for i := 0; i < len(config.Repositories); i++ {
		if config.Repositories[i].Url == repository.Url {
			repoConfig = &config.Repositories[i]
			return
		}
	}
	return
}
