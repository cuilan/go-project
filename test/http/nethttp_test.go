package http

import (
	"go-project/internal/api/nethttp"
	"go-project/internal/orm/gosql"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

// TestWithCleanup 会在包中所有其他测试运行之前被调用。
func TestWithCleanup(t *testing.T) {
	// --- 单个测试的 Setup ---
	t.Log("  【TestWithCleanup】=> Setup: 测试准备工作...")

	viper.SetConfigFile("../test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
	// 手动初始化 gosql 模块
	gosqlModule := gosql.NewModule()
	gosqlModule.Init(viper.GetViper())
	// 手动初始化 nethttp 模块
	nethttpModule := nethttp.NewModule()
	nethttpModule.Init(viper.GetViper())

	// --- 使用 t.Cleanup 注册 Teardown 函数 ---
	// t.Cleanup 接收一个函数，这个函数会在当前测试函数执行完毕后被调用。
	// 即使测试失败或 panic，它也保证会被执行。
	// 可以注册多个 Cleanup 函数，它们会以后进先出（LIFO）的顺序执行。
	t.Cleanup(func() {
		t.Log("  【TestWithCleanup】=> Teardown: 清理测试资源...")
		nethttpModule.Close()
		gosqlModule.Close()
	})

	t.Log("  【TestWithCleanup】=> 开始执行测试逻辑...")

	t.Run("测试用户注册", userRegister)
	t.Run("测试用户登录", userLogin)

	t.Log("  【TestWithCleanup】=> 测试逻辑执行完毕。")
}

// 测试用户注册
// 注意：由于 http api接口的集成测试需要执行前置后置准备工作
// 需确保执行入库在 TestWithCleanup 中，否则会报错
// 业务逻辑的测试方法名不能以 Test 开头，且对外不可见
// 确保测试逻辑方法不能被外部调用，也不可单独执行
func userRegister(t *testing.T) {
	t.Log("  【userRegister】=> 测试用户注册...")
	reqJson := `{"username": "zhangsan", "password": "123456"}`
	resp, err := http.Post("http://localhost:8081/user/register", "application/json", strings.NewReader(reqJson))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(body))
}

func userLogin(t *testing.T) {
	t.Log("  【userLogin】=> 测试用户登录...")
	reqJson := `{"username": "zhangsan", "password": "123456"}`
	resp, err := http.Post("http://localhost:8081/user/login", "application/json", strings.NewReader(reqJson))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(body))
}
