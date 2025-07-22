package utils

import (
	"errors"
	"log/slog"
	"strings"
)

// MaskDsn 将密码替换为 **********
func MaskDsn(driver, dsn string) (string, error) {
	if driver == "mysql" {
		// username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
		parts := strings.Split(dsn, "@")
		if len(parts) < 2 {
			slog.Error("invalid dsn", "dsn", dsn)
			return "", errors.New("invalid dsn")
		}
		userPass := strings.Split(parts[0], ":")
		if len(userPass) < 2 {
			slog.Error("invalid dsn", "dsn", dsn)
			return "", errors.New("invalid dsn")
		}
		userPass[1] = "********"
		parts[0] = strings.Join(userPass, ":")
		return strings.Join(parts, "@"), nil
	}
	if driver == "postgres" {
		// postgres://user:password@host:port/dbname?sslmode=disable
		parts := strings.Split(dsn, "@")
		if len(parts) < 2 {
			slog.Error("invalid dsn", "dsn", dsn)
			return "", errors.New("invalid dsn")
		}
		userPass := strings.Split(parts[0], ":")
		if len(userPass) < 3 {
			slog.Error("invalid dsn", "dsn", dsn)
			return "", errors.New("invalid dsn")
		}
		userPass[2] = "********"
		parts[0] = strings.Join(userPass, ":")
		return strings.Join(parts, "@"), nil
	}
	if driver == "oracle" {
		// user/password@host:port/dbname
		parts := strings.Split(dsn, "@")
		if len(parts) < 2 {
			slog.Error("invalid dsn", "dsn", dsn)
			return "", errors.New("invalid dsn")
		}
		userPass := strings.Split(parts[0], "/")
		if len(userPass) < 2 {
			slog.Error("invalid dsn", "dsn", dsn)
			return "", errors.New("invalid dsn")
		}
		userPass[1] = "********"
		parts[0] = strings.Join(userPass, "/")
		return strings.Join(parts, "@"), nil
	}
	if driver == "sqlserver" {
		// server=host;user id=user;password=password;database=dbname
		parts := strings.Split(dsn, ";")
		if len(parts) < 4 {
			slog.Error("invalid dsn", "dsn", dsn)
			return "", errors.New("invalid dsn")
		}
		userPass := strings.Split(parts[2], "=")
		if len(userPass) < 2 {
			slog.Error("invalid dsn", "dsn", dsn)
			return "", errors.New("invalid dsn")
		}
		userPass[1] = "********"
		parts[2] = strings.Join(userPass, "=")
		return strings.Join(parts, ";"), nil
	}
	// sqlite3 no password
	return dsn, nil
}
