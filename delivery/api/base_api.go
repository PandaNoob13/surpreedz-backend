package api

import (
	"surpreedz-backend/delivery/api/response"
	"surpreedz-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type BaseApi struct{}

func (b *BaseApi) ParseRequestBody(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(body); err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) ParseRequestFormData(c *gin.Context, requestModel interface{}, postFormKey ...string) error {
	mapRes := make(map[string]interface{})
	for _, v := range postFormKey {
		mapRes[v] = c.PostForm(v)
	}
	err := mapstructure.Decode(mapRes, &requestModel)
	utils.IsError(err)
	return nil
}

func (b *BaseApi) Success(c *gin.Context, data interface{}) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		c.Header("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}
	response.NewSuccessJsonResponse(c, data).Send()
}

func (b *BaseApi) Failed(c *gin.Context, err error) {
	response.NewErrorJsonResponse(c, err).Send()
}
