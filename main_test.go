package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/charmbracelet/x/exp/golden"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestMarshalYAML(t *testing.T) {
	cfg := Config{
		Field1: []string{}, // empty
		Field2: []string{"single item"},
		Field3: []string{"multiple", "items"},
		Field4: nil,
	}
	bts, err := yaml.Marshal(cfg)
	require.NoError(t, err)
	golden.RequireEqual(t, bts)
}

func TestUnmarshalYAML(t *testing.T) {
	bts, err := os.ReadFile("./testdata/src.yaml")
	require.NoError(t, err)
	var cfg Config
	require.NoError(t, yaml.Unmarshal(bts, &cfg))
	require.Equal(t, Config{
		Field1: []string{""},
		Field2: []string{"single item"},
		Field3: []string{"multiple", "items"},
		Field4: nil,
	}, cfg)
}

func TestMarshalJSON(t *testing.T) {
	cfg := Config{
		Field1: []string{}, // empty
		Field2: []string{"single item"},
		Field3: []string{"multiple", "items"},
		Field4: nil,
	}
	bts, err := json.Marshal(cfg)
	require.NoError(t, err)
	golden.RequireEqual(t, bts)
}

func TestUnmarshalJSON(t *testing.T) {
	bts, err := os.ReadFile("./testdata/src.json")
	require.NoError(t, err)
	var cfg Config
	require.NoError(t, json.Unmarshal(bts, &cfg))
	require.Equal(t, Config{
		Field1: []string{""},
		Field2: []string{"single item"},
		Field3: []string{"multiple", "items"},
		Field4: nil,
	}, cfg)
}
