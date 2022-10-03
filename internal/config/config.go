package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/pterm/pterm"
)

type Config struct {
	ApiToken        string          `toml:"ApiToken"`
	DefaultPageName string          `toml:"DefaultPage"`
	Page            map[string]Page `toml:"Page"`
}

type Page struct {
	ID   string            `toml:"ID"`
	Tags map[string]string `toml:"Tags"`
}

func (c *Config) DefaultPage() Page {
	return c.Page[c.DefaultPageName]
}

func LoadConfig() (Config, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("cannot obtain home dir: %s\n", err)
	}
	homeConfig := filepath.Join(dirname, ".todo_conf.toml")

	if _, err := os.Stat(homeConfig); err != nil {
		cfg := SetupConfig()
		err := writeConfig(cfg, homeConfig)
		if err != nil {
			return Config{}, fmt.Errorf("create config: %w", err)
		}
		pterm.Println(pterm.Green(pterm.Sprintf("Config saved at: %s", homeConfig)))
	}

	var cfg Config
	if _, err := toml.DecodeFile(homeConfig, &cfg); err != nil {
		return Config{}, fmt.Errorf("cannot decode config file: %s", err)
	}

	return cfg, nil
}

func SetupConfig() Config {
	cfg := Config{Page: map[string]Page{}}
	area, _ := pterm.DefaultArea.WithFullscreen(true).Start()
	area.Update(
		pterm.DefaultSection.Sprint("Hello there 👋 Let's perform some basic setup."),
		pterm.LightBlue("Enter you notion API token below.\n"),
		pterm.Yellow("You can obtain your API token by following the instructions here: https://developers.notion.com/docs/getting-started#step-1-create-an-integration\n"),
	)

	token := readLine()
	cfg.ApiToken = token

	area.WithCenter(true).Update(pterm.LightBlue("Got it.\nYour token is: "), pterm.Yellow(token))
	time.Sleep(2 * time.Second)

	setupDBBullets, _ := pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Text: "Clone the following database: https://fetinin.notion.site/5fd9edf71e4442dfab7c1075e543588e?v=a796cedf3179433da4162782a64a5599", TextStyle: pterm.NewStyle(pterm.FgBlue), Bullet: "-"},
		{Text: "Share database page with your integration.", TextStyle: pterm.NewStyle(pterm.FgBlue), Bullet: "-"},
		{Text: "Copy database ID. The ID is 32 characters long, containing numbers and letters.", TextStyle: pterm.NewStyle(pterm.FgBlue), Bullet: "-"},
		{Text: "Enter your DB ID below.", TextStyle: pterm.NewStyle(pterm.FgBlue), Bullet: "-"},
	}).Srender()

	area.Update(
		pterm.LightBlue("Now, let's set up your database page:\n"),
		setupDBBullets,
		pterm.Yellow("How to setup integration and obtain DB ID: https://developers.notion.com/docs/getting-started#step-2-share-a-database-with-your-integration"),
	)

	dbID := readLine()
	cfg.DefaultPageName = "work"
	cfg.Page["work"] = Page{
		ID: dbID,
		Tags: map[string]string{
			"Status": "To Do 🤖", "Category": "💼 Work", "Priority": "Не срочно",
		},
	}

	area.WithCenter(true).Update(pterm.LightBlue("Done!\nYour db page is: "), pterm.Yellow(dbID))
	time.Sleep(2 * time.Second)

	area.Stop()
	return cfg
}

func writeConfig(cfg Config, path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(cfg)
	if err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}

func readLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
