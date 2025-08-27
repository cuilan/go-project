#!/bin/bash
set -e

# ==============================================================================
# 脚本说明:
# 此脚本在 Docker 构建环境的 builder 阶段内部执行。
# 它会自动扫描 /app/cmd/ 目录下的所有子目录（即所有应用），
# 并将它们编译成 Go 二进制文件，存放到 /app/bin/ 目录下。
# ==============================================================================

echo "--- Starting build process inside Docker ---"

# 定义源码和输出目录
APP_ROOT="/app"
OUTPUT_DIR="${APP_ROOT}/bin"
CMD_DIR="${APP_ROOT}/cmd"

# 创建输出目录
echo "Creating output directory: ${OUTPUT_DIR}"
mkdir -p "${OUTPUT_DIR}"

# 检查 cmd 目录是否存在
if [ ! -d "${CMD_DIR}" ]; then
    echo "Error: Directory ${CMD_DIR} not found."
    exit 1
fi

# 遍历 /app/cmd/ 目录下的所有子目录进行编译
for cmd_path in "${CMD_DIR}"/*; do
    # 检查是否为目录
    if [ -d "${cmd_path}" ]; then
        app_name=$(basename "${cmd_path}")
        package_path="./cmd/${app_name}"
        output_path="${OUTPUT_DIR}/${app_name}"

        echo ""
        echo "--> Building application: ${app_name}"
        echo "    Package path: ${package_path}"
        echo "    Output path:  ${output_path}"

        # 执行编译
        # 使用 -trimpath 减小二进制文件大小
        # 使用 -ldflags "-s -w" 剥离调试信息，进一步减小体积
        go build -trimpath -ldflags="-s -w" -o "${output_path}" "${package_path}"

        echo "--> Successfully built ${app_name}"
    fi
done

echo ""
echo "--- All builds completed successfully ---"
