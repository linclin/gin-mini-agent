package file

import (
	"gin-mini-agent/models"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// FileCp 文件复制功能
func FileCp(c *gin.Context) {
	// 从请求体获取源文件路径和目标文件路径
	srcPath := c.PostForm("src")
	destPath := c.PostForm("dest")

	if srcPath == "" || destPath == "" {
		models.FailWithMessage("源路径和目标路径不能为空", c)
		return
	}

	// 执行文件复制
	err := copyFile(srcPath, destPath)
	if err != nil {
		models.FailWithMessage("文件复制失败: "+err.Error(), c)
		return
	}

	models.OkWithData("文件复制成功", c)
}

// copyFile 复制文件
func copyFile(src, dest string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 创建目标文件
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 复制文件内容
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	// 确保内容写入磁盘
	return destFile.Sync()
}
