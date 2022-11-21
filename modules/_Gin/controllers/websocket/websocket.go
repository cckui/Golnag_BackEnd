package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gopkg.in/olahol/melody.v1"
	//"log"
)

var GinMelody *melody.Melody

func WS_Setup(c *gin.Context) {

	GinMelody.HandleRequest(c.Writer, c.Request)

	GinMelody.HandleMessage(func(s *melody.Session, msg []byte) {
		bridge_msg := gjson.GetBytes(msg, "cmd")

		fmt.Sprintf("%s", bridge_msg.Str)

		if bridge_msg.Str == "park" {

			GinMelody.Broadcast([]byte("{\"cmd\":\"led_on\"}"))

		} else if bridge_msg.Str == "led_off" {

			GinMelody.Broadcast([]byte("{\"cmd\":\"unpark\"}"))
		} else {
			GinMelody.Broadcast([]byte("{\"cmd\":\"" + bridge_msg.Str + "\"}"))
		}
	})

}
