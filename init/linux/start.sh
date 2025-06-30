#!/bin/bash

# 此脚本由 systemd 服务执行，用于启动应用程序。
set -e

# 脚本期望从项目的根目录运行。
# 二进制文件与此脚本位于同一目录中。
SCRIPT_DIR=$(dirname "$0")
APP_NAME="your-go-project"
EXECUTABLE_PATH="${SCRIPT_DIR}/${APP_NAME}"
CONFIG_DIR="${SCRIPT_DIR}/../configs"

echo "正在启动 ${APP_NAME}..."
echo "可执行文件: ${EXECUTABLE_PATH}"
echo "配置文件目录: ${CONFIG_DIR}"

# 运行应用程序
# 环境变量由 systemd 服务文件加载。
${EXECUTABLE_PATH} --config-dir=${CONFIG_DIR}