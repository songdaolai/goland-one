package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

func main() {
	r := gin.Default()

	// 开关控制接口
	r.GET("/switch/:state", func(c *gin.Context) {
		state := c.Param("state") // 获取路径参数

		// 校验是否为合法值
		if state != "0" && state != "1" {
			c.JSON(400, gin.H{"error": "只能是0或1"})
			return
		}

		err := sendUDP(state, "192.168.2.34:7788") // 替换为你的设备IP:端口
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "发送成功 ✅",
			"state":   state,
		})
	})

	// 启动服务
	r.Run(":8080")
}

// sendUDP 向目标地址发送UDP消息
func sendUDP(message string, addr string) error {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return fmt.Errorf("连接失败: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("发送失败: %v", err)
	}
	return nil
}
