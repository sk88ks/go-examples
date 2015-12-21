package main

import (
	"encoding/json"

	"github.com/bluele/slack"
)

const (
	token       = "xoxp-2577181228-2577181230-14153880915-7382e64cb0"
	channelName = "batch"
)

var (
	slackAPI *slack.Slack
)

type AttachMent struct {
	PreText string `json:"pretext"`
	Text    string `json:"text"`
}

type invalidReferenceContent struct {
	ID       string
	AlbumID  string
	ArtistID string
}

func main() {
	slackAPI = slack.New(token)

	contents := []invalidReferenceContent{
		{
			ID:       "track000",
			AlbumID:  "album000",
			ArtistID: "artist000",
		},
		{
			ID:       "track001",
			AlbumID:  "album001",
			ArtistID: "artist001",
		},
	}

	err := sendInvalidMetaNotiToSlack("batch", "Invalid refference tracks\n", contents)
	if err != nil {
		panic(err)
	}

}

// SlackAttachMent is attachment value to post message to slack
type SlackAttachMent struct {
	PreText string `json:"pretext"`
	Text    string `json:"text"`
}

func sendInvalidMetaNotiToSlack(channelName, baseText string, contents []invalidReferenceContent) error {
	if len(contents) == 0 {
		return nil
	}

	channel, err := slackAPI.FindChannelByName(channelName)
	if err != nil {
		return err
	}

	attachments := []SlackAttachMent{}

	for i := range contents {
		attachment := SlackAttachMent{
			PreText: contents[i].ID,
		}

		if contents[i].AlbumID != "" {
			attachment.Text += "AlbumID: " + contents[i].AlbumID + " "
		}
		if contents[i].ArtistID != "" {
			attachment.Text += "ArtistID: " + contents[i].ArtistID
		}
		attachments = append(attachments, attachment)
	}

	attachmentsString, err := json.Marshal(attachments)
	if err != nil {
		return err
	}

	err = slackAPI.ChatPostMessage(channel.Id, baseText, &slack.ChatPostMessageOpt{
		AttachMents: string(attachmentsString),
	})
	if err != nil {
		return err
	}

	return nil
}
