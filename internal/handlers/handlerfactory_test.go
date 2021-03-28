package handlers

import (
	"testing"
)

func Test_Create_YamlConfigLoader_Success(t *testing.T) {
	_, err := NewHandler("yaml", YamlConfigLoader{Path: "hi"})
	if err != nil {
		t.Errorf("Test_Create_YamlConfigLoader_Success failed, expected err to be nil, got %v", err)
	}
}

func Test_Create_ConfigLoader_Failure(t *testing.T) {
	_, err := NewHandler("bla", YamlConfigLoader{Path: "hi"})
	if err == nil {
		t.Errorf("Test_Create_YamlConfigLoader_Success failed, expected err to be not nil, got %v", err)
	}
}
