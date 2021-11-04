package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

const userIniTemplate = `[email]
from = <from_gmail_email_address>
password = <gmail_password>
to = <comma_delimited_recepient_emails>

[mangadex]
username = <mangadex_username>
password = <mangadex_password>`

const userIniName = "user_config.ini"

func LoadUserIni() *ini.File {
	cfg, err := ini.Load(userIniName)
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		os.Exit(1)
	}
	return cfg
}
