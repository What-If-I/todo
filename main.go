package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pterm/pterm"

	"todo/internal/config"
	"todo/internal/notion"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		exit(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <title> [content]")
		return
	}

	title := os.Args[1]
	var content string
	if len(os.Args) > 2 {
		content = os.Args[2]
	}

	client := notion.NewClient(cfg.ApiToken)
	page, err := client.CreateTodo(context.Background(), cfg.DefaultPage().ID, title, content)
	if err != nil {
		exit(err)
	}
	PrintPage(page)
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func PrintPage(page notion.Page) {
	outerBox := pterm.DefaultBox.
		WithTitle(page.URL).
		WithTitleBottomCenter(true)

	pageBox := renderPageBox(page.Title, page.Content)
	tableBox := pterm.DefaultBox.WithTitle("Tags")
	tagsTable := renderTagsTable(page.Tags)
	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: pageBox}, {Data: tableBox.Sprint(tagsTable)}},
	}).Srender()

	outerBox.Print(panels)
}

func renderTagsTable(tags map[string]string) string {
	columns := make([]string, 0, len(tags))
	data := make([]string, 0, len(tags))
	for k, v := range tags {
		columns = append(columns, k)
		data = append(data, v)
	}

	table, _ := pterm.DefaultTable.WithData(pterm.TableData{columns, data}).Srender()
	return table
}

func renderPageBox(title, content string) string {
	if content == "" {
		title, content = "Todo", title
	}

	return pterm.DefaultBox.WithTitle(title).Sprint(content)
}
