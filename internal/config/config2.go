package config

//
// import (
// 	"net/url"
// 	"os"
// 	"time"
//
// 	"atomicgo.dev/cursor"
// 	"github.com/manifoldco/promptui"
// 	"github.com/pterm/pterm"
// )
//
// func setupConfig() Config {
// 	cfg := Config{}
// 	area, _ := pterm.DefaultArea.Start()
// 	area.Update(
// 		pterm.DefaultSection.Sprint("Hello there 👋"),
// 		pterm.LightBlue("Let's perform some basic setup."),
// 	)
// 	time.Sleep(2 * time.Second)
// 	area.Clear()
// 	area.Stop()
//
// 	for {
// 		requiredValidator := func(input string) error {
// 			if input == "" {
// 				return errors.New("value is required")
// 			}
// 			return nil
// 		}
//
// 		prompt := promptui.Prompt{
// 			Label:       pterm.LightBlue("Enter you JIRA username"),
// 			HideEntered: true,
// 			Validate:    requiredValidator,
// 		}
// 		result, err := prompt.Run()
// 		if err != nil {
// 			os.Exit(0)
// 		}
// 		cfg.JiraLogin = result
//
// 		prompt = promptui.Prompt{
// 			Label:       pterm.LightBlue("Now enter your password 🤫"),
// 			HideEntered: true,
// 			Mask:        '*',
// 			Validate:    requiredValidator,
// 		}
// 		result, err = prompt.Run()
// 		if err != nil {
// 			os.Exit(0)
// 		}
// 		cfg.JiraPassword = result
//
// 		urlValidator := func(input string) error {
// 			u, err := url.ParseRequestURI(input)
// 			if err != nil {
// 				return err
// 			}
// 			if u.Host == "" {
// 				return errors.New("host is missing")
// 			}
// 			return nil
// 		}
//
// 		prompt = promptui.Prompt{
// 			Label:       pterm.LightBlue("Almost done! Now enter JIRA url"),
// 			HideEntered: true,
// 			Validate:    urlValidator,
// 		}
// 		result, err = prompt.Run()
// 		if err != nil {
// 			os.Exit(0)
// 		}
// 		cfg.JiraURL = result
//
// 		confirmed, _ := pterm.DefaultInteractiveConfirm.Show(pterm.Sprint(
// 			pterm.LightBlue("Got it👌"),
// 			pterm.LightBlue("\nYour login is: "), pterm.Yellow(cfg.JiraLogin),
// 			pterm.LightBlue("\nPassword is: "), pterm.Yellow(cfg.JiraPassword),
// 			pterm.LightBlue("\nJIRA url is: "), pterm.Yellow(cfg.JiraURL),
// 			pterm.LightBlue("\nCorrect?"),
// 		))
// 		if confirmed {
// 			cursor.ClearLinesUp(5)
// 			break
// 		}
// 		cursor.ClearLinesUp(5)
// 	}
//
// 	return cfg
// }
