#!/bin/bash

# 此脚本将 Go 项目安装为 systemd 服务。
# 需要以 sudo 权限运行。

set -e

# --- 配置 ---
# !!! 用户在运行此脚本前应修改这些变量。 !!!
APP_NAME="your-go-project"
USER="youruser"   # 服务将以此用户身份运行。
GROUP="yourgroup" # 服务将以此用户组身份运行。

# --- 路径 ---
# 脚本假定它位于解压后的 'release/init/linux' 目录中。
SOURCE_DIR=$(dirname "$0")
# 可执行文件、配置等位于 'release' 根目录
RELEASE_ROOT_DIR="${SOURCE_DIR}/../.."
# 智能查找二进制文件，它应该是 release 根目录下唯一的可执行文件
BINARY_SOURCE_PATH=$(find "${RELEASE_ROOT_DIR}" -maxdepth 1 -type f -executable -name "${APP_NAME}*")

# 目标路径
INSTALL_DIR="/opt/${APP_NAME}"
SERVICE_FILE_DEST="/etc/systemd/system/${APP_NAME}.service"

echo "--- 开始为 ${APP_NAME} 进行安装 ---"

# 检查是否以 root 身份运行
if [ "$EUID" -ne 0 ]; then
  echo "请以 root 或 sudo 身份运行此脚本。"
  exit 1
fi

# 检查二进制文件是否存在
if [ ! -f "${BINARY_SOURCE_PATH}" ]; then
    echo "在路径 ${BINARY_SOURCE_PATH} 未找到二进制文件"
    echo "请确保压缩包结构正确。"
    exit 1
fi

echo "1. 正在创建用户和用户组 '${USER}:${GROUP}' (如果不存在)..."
if ! getent group "${GROUP}" >/dev/null; then
    groupadd -r "${GROUP}"
    echo "用户组 '${GROUP}' 已创建。"
fi
if ! id "${USER}" >/dev/null 2>&1; then
    useradd -r -g "${GROUP}" -s /sbin/nologin -m -d "${INSTALL_DIR}" "${USER}"
    echo "用户 '${USER}' 已创建。"
fi

echo "2. 正在创建安装目录: ${INSTALL_DIR}"
mkdir -p "${INSTALL_DIR}/configs"

echo "3. 正在复制应用文件..."
cp -v "${BINARY_SOURCE_PATH}" "${INSTALL_DIR}/${APP_NAME}"
cp -v "${RELEASE_ROOT_DIR}/init/linux/start.sh" "${INSTALL_DIR}/start.sh"
cp -vR "${RELEASE_ROOT_DIR}/configs/"* "${INSTALL_DIR}/configs/"
cp -v "${RELEASE_ROOT_DIR}/init/linux/${APP_NAME}.env" "${INSTALL_DIR}/${APP_NAME}.env"

echo "4. 正在设置文件权限..."
chown -R "${USER}:${GROUP}" "${INSTALL_DIR}"
chmod +x "${INSTALL_DIR}/start.sh"
chmod +x "${INSTALL_DIR}/${APP_NAME}"

echo "5. 正在创建并安装 systemd 服务文件..."
# 使用变量替换创建一个临时的服务文件
SERVICE_FILE_TEMPLATE="${RELEASE_ROOT_DIR}/init/linux/${APP_NAME}.service"
TEMP_SERVICE_FILE="/tmp/${APP_NAME}.service.$$"

sed -e "s|{{APP_NAME}}|${APP_NAME}|g" \
    -e "s|{{USER}}|${USER}|g" \
    -e "s|{{GROUP}}|${GROUP}|g" \
    -e "s|{{INSTALL_DIR}}|${INSTALL_DIR}|g" \
    "${SERVICE_FILE_TEMPLATE}" > "${TEMP_SERVICE_FILE}"

# 复制处理后的服务文件
cp -v "${TEMP_SERVICE_FILE}" "${SERVICE_FILE_DEST}"
rm "${TEMP_SERVICE_FILE}"


echo "6. 正在重载 systemd，启用并启动服务..."
systemctl daemon-reload
systemctl enable "${APP_NAME}.service"
systemctl start "${APP_NAME}.service"

echo "--- 安装完成! ---"
echo "您可以使用以下命令检查服务状态: systemctl status ${APP_NAME}.service"
echo "您可以使用以下命令查看服务日志: journalctl -u ${APP_NAME}.service -f"
