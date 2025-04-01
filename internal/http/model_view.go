package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: 10000,
		Msg:  "",
		Data: data,
	})
}

func Fail(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: 10001,
		Msg:  "失败",
		Data: data,
	})
}

func FailWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: 10001,
		Msg:  msg,
		Data: data,
	})
}

func FailWithCodeMsg(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ---------------- request ----------------

const (
	Nfs PushTypeStruct = 0
	Ftp PushTypeStruct = 1
)

// PushTypeStruct 推送类型，0：NFS，1：FTP
type PushTypeStruct int

type PushConfigModel struct {
	Id           int64          `json:"id"`
	Satellite    string         `json:"satellite"`
	SaveStatus   int            `json:"saveStatus"`   // 保存状态，0：初始保存L012，1：保存推送类型，2：测试，3：测试完成
	FileLevel    string         `json:"fileLevel"`    // 文件级别
	L2Algorithms []L2Algorithm  `json:"l2Algorithms"` // 只有保存L2才校验此参数
	PushType     PushTypeStruct `json:"pushType"`     // 推送类型，0：NFS，1：FTP
	NfsPath      string         `json:"nfsPath"`      // NFS映射路径
	FtpHost      string         `json:"ftpHost"`      // FTP地址
	FtpPort      int            `json:"ftpPort"`      // FTP端口
	FtpName      string         `json:"ftpName"`      // FTP用户名
	FtpPasswd    string         `json:"ftpPasswd"`    // FTP密码
	TestKey      string         `json:"testKey"`      // 测试key
}

// ---------------- response ----------------

type HealthResp struct {
	Satellite string `json:"satellite"`
}

// SatelliteFileLevel 卫星文件类型配置
type SatelliteFileLevel struct {
	Satellite    string                   `json:"satellite"`
	L0           SatelliteFileLevelConfig `json:"L0"`
	L1           SatelliteFileLevelConfig `json:"L1"`
	L2           SatelliteFileLevelConfig `json:"L2"`
	L2Algorithms []L2Algorithm            `json:"l2Algorithms"`
}

type SatelliteFileLevelConfig struct {
	Enable bool   `json:"enable"`
	Path   string `json:"path"`
}

type L2Algorithm struct {
	GroupName string `json:"groupName"`
	KindGroup string `json:"kindGroup"`
	AlgoName  string `json:"algoName"`
	Name      string `json:"name"`
}
