package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
	"github.com/liucxer/resource-manage/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"net/http/pprof"
)

var RootRouter = (ginS.Group("/resource-manage").
	Use(middleware.CORSMiddleware(), middleware.LoggerAccessToFile()).(*gin.RouterGroup))
var V1Router = RootRouter.Group("/v1")

var ResourcePath string

func InitRouter(resourcePath string, hosts []string) {
	ResourcePath = resourcePath

	ginS.Use(middleware.LoggerAccessToFile())
	// pprof
	ginS.GET("/debug/pprof/", gin.WrapF(pprof.Index))
	ginS.GET("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	ginS.GET("/debug/pprof/profile", gin.WrapF(pprof.Profile))
	ginS.POST("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	ginS.GET("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	ginS.GET("/debug/pprof/trace", gin.WrapF(pprof.Trace))
	ginS.GET("/debug/pprof/allocs", gin.WrapH(pprof.Handler("allocs")))
	ginS.GET("/debug/pprof/block", gin.WrapH(pprof.Handler("block")))
	ginS.GET("/debug/pprof/goroutine", gin.WrapH(pprof.Handler("goroutine")))
	ginS.GET("/debug/pprof/heap", gin.WrapH(pprof.Handler("heap")))
	ginS.GET("/debug/pprof/mutex", gin.WrapH(pprof.Handler("mutex")))
	ginS.GET("/debug/pprof/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))

	// swagger
	ginS.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ginS.Use(middleware.CORSMiddleware()).Static("/web", "./templates")
	// 资源
	ginS.Use(middleware.CORSMiddleware(), middleware.AccessLimit(hosts)).StaticFS("/video", http.Dir(ResourcePath+"/video/"))     // 浏览视频
	ginS.Use(middleware.CORSMiddleware(), middleware.AccessLimit(hosts)).StaticFS("/picture", http.Dir(ResourcePath+"/picture/")) // 浏览图片
	ginS.POST("/resource-manage/v1/videos", VideoCreate)

	//ginS.POST("/api/upload/startChunk", StartChunk) // 上传视频(管理员)
	ginS.POST("/resource-manage/v1/pictures", PictureCreate)
	ginS.POST("/resource-manage/v1/videos/multipart", MultipartVideoCreate) // 上传图片(管理员)
}
