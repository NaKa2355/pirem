package config

import (
	"encoding/json"
	"os"
)

type DeviceConfig struct {
	Name       string          `json:"name"`
	ID         string          `json:"id"`
	ModuleName string          `json:"module_name"`
	Config     json.RawMessage `json:"config"`
}

type Config struct {
	Devices                   []DeviceConfig `json:"devices"`
	DeviceMutexLockDeadlineMs int            `json:"device_mutex_lock_deadline_ms"`
	DeviceSendIRIntervalMs    int            `json:"device_send_ir_interval_ms"`
	EnableReflection          bool           `json:"enable_reflection"`
	Debug                     bool           `json:"debug"`
}

func ReadConfig(filePath string) (*Config, error) {
	config := &Config{}
	config_data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(config_data, config)
	return config, err
}
