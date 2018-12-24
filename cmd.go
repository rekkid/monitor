package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"unsafe"
)

type Command struct {
	Cmd string `json:"Cmd" binding:"required"`
}

func RunCmd() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.POST("/cmd", func(c *gin.Context) {
		var command Command
		if err := c.ShouldBindJSON(&command); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		log.Info("Running command and waiting for it to finish...")
		cmd := exec.Command(command.Cmd)
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			log.Error(err)
			c.JSON(
				http.StatusOK,
				gin.H{
					"code":  http.StatusOK,
					"error": err.Error(),
				},
			)
		}

		str := *(*string)(unsafe.Pointer(&stdoutStderr))
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":   http.StatusOK,
				"result": str,
			},
		)
	})

	return e
}
