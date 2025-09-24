/*
 * Copyright 2025 coze-dev Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package postgresql

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config 保存 PostgreSQL 连接配置
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

// NewConfigFromEnv 从环境变量读取配置
func NewConfigFromEnv() *Config {
	cfg := &Config{
		Host:     getEnv("POSTGRES_HOST", "127.0.0.1"),
		Port:     getEnv("POSTGRES_PORT", "5432"),
		User:     getEnv("POSTGRES_USER", "postgres"),
		Password: getEnv("POSTGRES_PASSWORD", ""),
		Database: getEnv("POSTGRES_DATABASE", "postgres"),
		SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"), // 默认禁用 SSL
	}
	return cfg
}

// DSN 返回 PostgreSQL 连接字符串
func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode,
	)
}

func New() (*gorm.DB, error) {
	cfg := NewConfigFromEnv()
	dsn := cfg.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres, dsn=%s, err=%w", dsn, err)
	}

	return db, nil
}

// getEnv 工具函数：优先取环境变量，没设置就用默认值
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
