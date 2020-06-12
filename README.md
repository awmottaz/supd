`supd` (**s**crum **upd**ate) is a CLI app I made for myself to track my daily updates for work. Every day I like to track my plan for the day and the things I actually did that day.

This app is under development.

## TODO

### Core features

#### main command

- [x] basic app can configure the file containing the scrum updates (`-f, --file` flag)
- [ ] prompt for today's plan
- [ ] prompt for completed tasks
- [ ] prompt for note with `-t, --note` flag
- [ ] `-e, --edit` flag opens the update file in `$EDITOR`

#### `print` sub-command

- [ ] prints today's update by default (pretty)
- [ ] `--json` flag prints as JSON
- [ ] `-d, --date DATE` flag selects an update on `DATE`
- [ ] `-n, --num N` flag selects the last `N` updates
- [ ] `-s, --since DATE` flag selects all updates between `DATE` and today, inclusive

### Nice-to-have features

#### main command

- [ ] pretty help text (modify Usage)

#### `print` sub-command

- [ ] `-c, --clipboard` flag copies output to clipboard

### Other

- [ ] better comments
- [ ] code review from an experienced Go dev
- [ ] automated versioning
- [ ] automated deployments
