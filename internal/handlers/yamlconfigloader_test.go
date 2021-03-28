package handlers

import "testing"

func Test_LoadYaml_Success(t *testing.T) {

	yamlConfigLoader := YamlConfigLoader{
		Path: "data/default_test_data.yaml",
	}

	config, err := yamlConfigLoader.Load()

	if err != nil {
		t.Errorf("Test_LoadYaml_Success failed, expected err to be nil, got %v", err)
	}

	if config.Tests[0].Name != "Example Success 1" {
		t.Errorf("Test_LoadYaml_Success failed, expected config.Tests[0].Name to be 'Example Success', got %v", config.Tests[0].Name)
	}

	if config.Export.Http != "http://localhost:3000/injest" {
		t.Errorf("Test_LoadYaml_Success failed, expected config.Export.Http to be 'http://localhost:3000/injest', got %v", config.Export.Http)
	}
}
