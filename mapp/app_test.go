package mapp

import (
	"context"
	"net"
	"testing"

	"github.com/charliego93/logger"
	"github.com/stretchr/testify/require"
)

func TestNewApp(t *testing.T) {
	NewApp().Run(context.Background())
}

func TestDefaultFunc(t *testing.T) {
	var arr []string
	t.Log(append([]string{"first"}, arr...))
}

func TestCheckAddress(t *testing.T) {
	host, port, err := net.SplitHostPort(":8080")
	require.NoError(t, err)
	logger.Infof("Host: %s, Port: %s", host, port)
}
