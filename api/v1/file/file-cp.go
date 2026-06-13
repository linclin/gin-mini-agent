package file

import (
	"gin-mini-agent/models"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// FileCp 文件复制功能
func FileCp(c *gin.Context) {
	srcPath := c.PostForm("src")
	destPath := c.PostForm("dest")

	if srcPath == "" || destPath == "" {
		models.Fail(c, "源路径和目标路径不能为空")
		return
	}

	err := copyFile(srcPath, destPath)
	if err != nil {
		models.Fail(c, "文件复制失败: "+err.Error())
		return
	}

	models.Ok(c, "文件复制成功")
}

// copyFile 复制文件
func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}
