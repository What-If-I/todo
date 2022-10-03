package notion

import (
	"context"

	"github.com/dstotijn/go-notion"
)

type Page struct {
	ID      string
	Title   string
	URL     string
	Tags    map[string]string
	Content string
}

type Client struct {
	client *notion.Client
}

func NewClient(token string) *Client {
	return &Client{
		client: notion.NewClient(token),
	}
}

func (c *Client) CreateTodo(ctx context.Context, parentDB, title, content string) (Page, error) {
	var pageContent []notion.Block
	if content != "" {
		pageContent = append(pageContent,
			notion.Block{Paragraph: &notion.RichTextBlock{Text: []notion.RichText{{Text: &notion.Text{Content: content}}}}},
		)
	}

	p, err := c.client.CreatePage(ctx, notion.CreatePageParams{
		ParentType: notion.ParentTypeDatabase,
		ParentID:   parentDB,
		DatabasePageProperties: &notion.DatabasePageProperties{
			"Name":     {Title: []notion.RichText{{Text: &notion.Text{Content: title}}}},
			"Status":   {Select: &notion.SelectOptions{Name: "To Do 🤖"}},
			"Category": {Select: &notion.SelectOptions{Name: "💼  Work"}},
			"Priority": {Select: &notion.SelectOptions{Name: "Не срочно"}},
		},
		Children: pageContent,
	})
	if err != nil {
		return Page{}, err
	}

	tags := map[string]string{
		"Status":   "To Do 🤖",
		"Category": "💼  Work",
		"Priority": "Не срочно",
	}
	return Page{ID: p.ID, Title: title, Tags: tags, Content: content, URL: p.URL}, nil
}
