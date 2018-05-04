package directmsg

import (
	s "strings"

	"github.com/JayDShah/TwitterAPI"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var log = logger.GetLogger("activity-twitterdirectmessage")

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	consumerKey := s.TrimSpace(context.GetInput("consumerKey").(string))
	consumerSecret := s.TrimSpace(context.GetInput("consumerSecret").(string))
	accessToken := s.TrimSpace(context.GetInput("accessToken").(string))
	accessTokenSecret := s.TrimSpace(context.GetInput("accessTokenSecret").(string))
	directmsg := s.TrimSpace(context.GetInput("directmsg").(string))
	user := s.TrimSpace(context.GetInput("user").(string))

	if len(consumerKey) == 0 {
		context.SetOutput("statusCode", "101")
		context.SetOutput("message", "Consumer Key field is blank")
	} else if len(consumerSecret) == 0 {

		context.SetOutput("statusCode", "102")

		context.SetOutput("message", "Consumer Secret field is blank")

	} else if len(accessToken) == 0 {

		context.SetOutput("statusCode", "103")

		context.SetOutput("message", "Access Token field is blank")

	} else if len(accessTokenSecret) == 0 {

		context.SetOutput("statusCode", "104")

		context.SetOutput("message", "Access Token Secret field is blank")

	} else if len(directmsg) == 0 {

		context.SetOutput("statusCode", "105")

		context.SetOutput("message", "Direct Message cannot be blank")

	} else if len(user) == 0 {

		context.SetOutput("statusCode", "106")

		context.SetOutput("message", "User field is blank")

	} else {
		code, msg := Twitter.DirectMessage(consumerKey, consumerSecret, accessToken, accessTokenSecret, directmsg, user)
		context.SetOutput("statusCode", code)

		context.SetOutput("message", msg)
	}

	return true, err
}
