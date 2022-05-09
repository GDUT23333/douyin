package main

import (
	"github.com/gin-gonic/gin"
	"sync"
	"user_module/dao"
)

func main(){
	//initial
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	go func() {
		dao.InitialDB()
		waitGroup.Done()
	}()
	r := gin.Default()
	go func() {
		InitRouter(r)
		waitGroup.Done()
	}()
	waitGroup.Wait()
	r.Run()
}


