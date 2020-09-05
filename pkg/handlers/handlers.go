package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/docker/distribution/notifications"
	"github.com/labstack/echo/v4"

	"github.com/ziwon/dokkery/pkg/app"
	"github.com/ziwon/dokkery/pkg/config"
	"github.com/ziwon/dokkery/pkg/exec"
	"github.com/ziwon/dokkery/pkg/models"
)

func HandleEventPullOrPush(c echo.Context) error {
	app := c.Get("app").(*app.App)

	payload, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		app.R.Logger.Errorf("Error on reading registry event payload: %s", err)
		return err
	}
	defer c.Request().Body.Close() //nolint:errcheck

	envelope := &notifications.Envelope{}
	err = json.Unmarshal(payload, envelope)
	if err != nil {
		app.R.Logger.Errorf("Error on unmarshalling registry event payload: %s", err)
		return err
	}

	ret := ""
	for _, event := range envelope.Events {
		switch event.Action {
		case "delete":
			ret, _ = handleDeleteEvent(app.R, app.Config, event)
		case "pull":
			ret, _ = handlePullEvent(app.R, app.Config, event)
		case "push":
			ret, _ = handlePushEvent(app.R, app.Config, event)
		default:
			ret = "Unknown event action"
		}
	}

	return c.JSON(http.StatusOK, &models.Response{
		Data: ret,
	})
}

func handleDeleteEvent(r *echo.Echo, config config.Config, event notifications.Event) (string, error) {
	r.Logger.Debug("handling delete event...")
	str := fmt.Sprintf("Delete Event: %s, %s, %s", event.Target.URL, event.Target.Repository, event.Target.Tag)
	return str, nil
}

func handlePullEvent(r *echo.Echo, config config.Config, event notifications.Event) (string, error) {
	r.Logger.Debug("handling pull event...")
	str := fmt.Sprintf("Pull Event: %s, %s, %s", event.Target.URL, event.Target.Repository, event.Target.Tag)
	return str, nil
}

func handlePushEvent(r *echo.Echo, config config.Config, event notifications.Event) (string, error) {
	r.Logger.Debug("handling push event...")

	domain := config.Registry.Domain
	services := config.Registry.OnPush.Services
	repo := event.Target.Repository
	sha256 := getSHA256Code(event.Target.URL)
	tag := event.Target.Tag

	//imageTag := fmt.Sprintf("%s/%s:%s", domain, repo, tag)
	imageSha256 := fmt.Sprintf("%s/%s:%s@%s", domain, repo, tag, sha256)

	for _, service := range services {
		if strings.Contains(service.Image, repo) {
			for _, cmd := range service.Pre {
				c := strings.Replace(cmd, "{}", imageSha256, 1)
				go func(c string) {
					notify(exec.Execute(c),
						fmt.Sprintf("to pull: `%s:%s`", repo, sha256),
						config)
				}(c)
			}

			for _, cmd := range service.Post {
				c := strings.Replace(cmd, "{}", service.Name, 1)
				go func(c string) {
					notify(exec.Execute(c),
						fmt.Sprintf("to update: `%s`", repo),
						config)
				}(c)
			}
		}
	}

	return "", nil
}

func getSHA256Code(url string) string {
	idx := strings.LastIndex(url, "/")
	return url[idx+1:]
}

func notify(ret bool, text string, config config.Config) {
	if ret {
		msg := fmt.Sprintf(config.Notify.Slack.Message.Success.Head, text)
		sendSlack(
			config.Notify.Slack.WebHook,
			config.Notify.Slack.Channel,
			msg)
	} else {
		msg := fmt.Sprintf(config.Notify.Slack.Message.Fail.Head, text)
		sendSlack(
			config.Notify.Slack.WebHook,
			config.Notify.Slack.Channel,
			msg)
	}
}

func sendSlack(webhook string, channel string, message string) {
	payload := slack.Payload{
		Text:     message,
		Username: "dokkery",
		Channel:  channel,
	}
	err := slack.Send(webhook, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
