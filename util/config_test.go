package util

import (
    "os"
    "testing"
    "github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
    // Test with environment variables and temp config file
    t.Run("EnvironmentVariables", func(t *testing.T) {
        // Create temp dir and config file
        tempDir, err := os.MkdirTemp("", "config-test")
        require.NoError(t, err)
        defer os.RemoveAll(tempDir)

        configContent := `
DB_DRIVER=mysql
DB_SOURCE=mysql://root:password@localhost:3306/simple_bank
SERVER_ADDRESS=127.0.0.1:9090
`
        err = os.WriteFile(tempDir+"/app.env", []byte(configContent), 0644)
        require.NoError(t, err)

        // Set environment variables that should override file config
        os.Setenv("DB_DRIVER", "postgres")
        os.Setenv("DB_SOURCE", "postgres://test:test@localhost:5432/test_db")
        os.Setenv("SERVER_ADDRESS", "0.0.0.0:8080")
        
        defer func() {
            os.Unsetenv("DB_DRIVER")
            os.Unsetenv("DB_SOURCE")
            os.Unsetenv("SERVER_ADDRESS")
        }()
        
        config, err := LoadConfig(tempDir)
        require.NoError(t, err)
        
        // Environment variables should take precedence
        require.Equal(t, "postgres", config.DBDriver)
        require.Equal(t, "postgres://test:test@localhost:5432/test_db", config.DBSource)
        require.Equal(t, "0.0.0.0:8080", config.ServerAddress)
    })
    
    t.Run("InvalidPath", func(t *testing.T) {
        _, err := LoadConfig("/invalid/path")
        require.Error(t, err)
    })
}
