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

package database

import (
	"fmt"
	"os"
	"strings"

	"github.com/coze-dev/coze-studio/backend/infra/impl/mysql"
	"github.com/coze-dev/coze-studio/backend/infra/impl/postgresql"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dbType := os.Getenv("DB_TYPE")

	// 处理环境变量未设置的情况
	if dbType == "" {
		dbType = "mysql"
	}
	// 转换为小写以支持大小写不敏感的匹配
	dbType = strings.ToLower(dbType)

	switch dbType {
	case "mysql":
		return mysql.New()
	case "postgres", "postgresql", "pg":
		return postgresql.New()
	default:
		// 抛出不支持的数据库类型异常
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
