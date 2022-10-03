# ✍ Todo - utility to save your TODOs
Simple CLI that helps save TODOs to Notion.

## Usage
```bash
Usage: todo <title> [content]
```

## Install
### From source
```bash
go install github.com/What-If-I/tlog@latest 
```
_go >= 1.18 required_

## Configuration
Upon first run, utility will create config file called `.todo_conf.toml` at your home directory. You can edit config to setup different pages. Only page set as DefaultPage can be used at this moment.

Config example:
```bash
cat ~/.time_logger_conf.toml
```
```toml
ApiToken = "secret_xxx"
DefaultPage = "work"

[Page]
    [Page.work]
    ID = "page_uuid"
        [Page.work.Tags]
        Category = "💼 Work"
        Priority = "Не срочно"
        Status = "To Do 🤖"
```

### Things to do
- [ ] Add `config show`, `config set-page`, `config edit` commands
- [ ] Automate releases with https://goreleaser.com/quick-start/
