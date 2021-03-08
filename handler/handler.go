package handler

import (
	"GinStudy/log"
	"GinStudy/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"net/http"
)

var(
	Ginserver *gin.Engine

)

func InitService(){
	Ginserver.MaxMultipartMemory=500<<20
	gin.DefaultWriter=log.Logger
	Ginserver.Use(gin.Logger())

	//gMock=viper.GetBool("server.mock")
	baseUrl:=viper.GetString("server.baseUrl")

	//r:=gin.Default()
	group:=Ginserver.Group(baseUrl)
	{
		group.POST("/test1",Test1)
		//group.POST("/test2",Test2)
	}
	//Ginserver.POST(baseUrl,Test1)


	//baseUrl:=viper.GetString("server.baseUrl")
	//
	//Ginserver.POST(baseUrl, func(c *gin.Context) {
	//	var json test1Req
	//	//json := make(map[string]interface{})
	//	if err := c.ShouldBindJSON(&json); err != nil {
	//		c.String(http.StatusBadRequest, err.Error())
	//	}
	//	c.String(http.StatusOK, json.Action)
	//})


}

type test1Req struct {
	Action string `json:"action"`
	Data interface{} `json:data`
}
func Test1(c *gin.Context) {
	req:=test1Req{}
	var rsp interface{}
	var err interface{}
	if err=c.ShouldBindJSON(&req);err!=nil{
		rsp=ErrorMsg{ErrCodeApiGatewayFormat,ErrMsgApiGatewayFormat}
		c.JSON(http.StatusOK,&rsp)
		return
	}

	//log.Logger.Info().Msgf("%v",c.Request)
	log.Logger.Info().Msgf("%v",req.Data)
	//rsp=Test1fdata2{
	//	Error: ErrorMsg{ErrCodeOK,ErrMsgOK},
	//}
	rsp=Test1f(req.Data.(map[string]interface{}))
	c.JSON(http.StatusOK,&rsp)
}
type Test1fdata struct {
	Id string `json:"id"`
	Name string `json:"name"`
}
type Test1fdata2 struct {
	Error ErrorMsg                                 `json:"Error"`
	Data  Test1fdata `json:"Data"`
	}
func Test1f(m map[string]interface{}) interface{}{
	req:=Test1fdata{}
	rsp:=Test1fdata2{
		Error: ErrorMsg{ErrCodeOK,ErrMsgOK},
	}
	err:=mapstructure.Decode(m,&req)
	if err!=nil{
		return rsp.Error.WithErrRsp(ErrCodeInvalidRequest, ErrMsgInvalidRequest)
	}
	name,_:=model.GetName()
	fmt.Printf("name     %v",name)
	rsp.Data=Test1fdata{
		Id: req.Id,
		Name: name,
	}
	return rsp
}