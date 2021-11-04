package main

import (
	"fmt"
	"go-meater-meter/config"
	"go-meater-meter/meater/authentication"
	"time"
)

func main() {
	userCfg := config.LoadUserIni()
	pgCfg := config.LoadPgIni()
	to_email := gomail.ParseToEmail(user_cfg)
	sessionToken := authentication.GetAuth(user_cfg)
	manga := feed.GetFollowedMangaFeedList(sessionToken)
	newLastRunTime := time.Now()

	for _, element := range manga {
		chapterCreationDate := title.ParseCreationDate(element["createdDate"])
		newLastRunTime = chapterCreationDate
		logTime := time.Now().Format(time.RFC3339)

		if lastRunTime.After(chapterCreationDate) || lastRunTime.Equal(chapterCreationDate) {
			fmt.Println(logTime, "Skipping alert for", element["title"], "Chapter", element["chapter"]+".")
			continue
		}

		emailBody := gomail.PrepMessageBody(element)
		alert, err := gomail.SendEmailSMTP(to_email, emailBody, element["title"], user_cfg)

		if err == nil && alert {
			fmt.Println(logTime, "Alert sent for", element["title"], "Chapter", element["chapter"]+".")
		} else {
			fmt.Println(logTime, "Failed to send alert for", element["title"], "Chapter", element["chapter"]+".")
		}
	}
	//Update the last run time ini
	config.UpdateAppIni(newLastRunTime)
}
