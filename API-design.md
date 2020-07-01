# supd API Design

- [supd API Design](#supd-api-design)
  - [Primary functions](#primary-functions)
  - [Updates data](#updates-data)
  - [Viewing an update](#viewing-an-update)
  - [CLI API](#cli-api)
    - [Config](#config)
    - [Global options](#global-options)
    - [`set` command](#set-command)
      - [flags](#flags)
    - [`add` command](#add-command)
      - [flags](#flags-1)
    - [`rm` command](#rm-command)
      - [flags](#flags-2)
    - [`view` command](#view-command)
      - [flags](#flags-3)
    - [`edit` command](#edit-command)
      - [flags](#flags-4)

This document describes the desired API and functionality of the `supd` tool. For documentation about what `supd` can actually do, run `supd help`.

## Primary functions

The `supd` program is designed to make it easy to track the information you need for daily scrum updates. This includes three things:

1. What did you do on the previous work day?
2. What is your plan for today?
3. Additional notes for your team (e.g. notice of vacation)

Thus, the following actions are required from this program:

- set your plan for today
- track tasks as you complete them
- create notes as you need to
- view update information

## Updates data

All of the data for these updates will be stored in a JSON file.

The data is formatted as a map from ISO date strings (`yyyy-mm-dd`) to objects with the following properties:

- `plan` (`string`): the plan for that day
- `done` (`Array<string>`): a list of tasks completed that day
- `notes` (`Array<string>`): a list of notes associated with that day

By convention, entries are sorted in reverse-chronological order so that the most recent entries are at the top. `supd` will do this automatically.

Here is a sample updates file:

```json
{
    "2020-06-25": {
        "plan": "I will complete X",
        "done": [
            "Thing to prepare for X",
            "Most of X"
        ]
    },
    "2020-06-23": {
        "plan": "Finishing up Y",
        "done": [
            "Consulted Alice re: Y thing",
            "Completed Y"
        ],
        "notes": [
            "Reminder: I will be gone tomorrow"
        ]
    }
}
```

## Viewing an update

When viewing an update, you usually want the following information:

- Your plan from today's entry
- Your done list from the previous entry
- Your notes from the previous entry and today's entry

## CLI API

### Config

Configure global options for `supd` with a TOML-formatted config file. By default, the config is read from `$HOME/.supd.toml`.

```toml
# Set the filepath for the updates file
updatesFile = "$HOME/updates.json"

# Set the editor to use when running the edit command
editor = "$EDITOR"
```

### Global options

- `--config PATH`: Set the path to a config file.
- `--updates-file PATH`: Set the path to the updates file

### `set` command

Use this command to imperatively set a value.

**Set a plan for today**

```
$ supd set plan "This is my plan"
plan saved for 2020-06-30
```

**overwrite the plan for today**

```
$ supd set plan "This is a new plan" --force
plan updated for 2020-06-30
```

#### flags

- `--force` overwrite plan
- `--date DATE` to set on a different date

### `add` command

For adding completed tasks and notes.

**add a completed task**

```
$ supd add done "Did a thing"
completed task saved for 2020-06-30
```

**add a note**

```
$ supd add note "Reminder: I will be gone on Friday"
note saved for 2020-06-30
```

#### flags

- `--date DATE` to set on a different date

### `rm` command

Delete things.

**Delete an entire update**

```
$ supd rm update --date 2020-06-23
This will remove all update information for June 23, 2020.
Are you sure? [y/N] y
Update for 2020-06-23 removed
```

**Delete your plan for today**

```
$ supd rm plan
Remove plan for today, June 30? [y/N] y
Plan for 2020-06-30 removed
```

**Delete a completed task**

```
$ supd rm done
DONE on 2020-06-30:
  1: Did a thing
  2: Did another thing
  3: Got stuck on a thing
Remove which one? [1-3] 2
Removing task 2: "Did another thing"
Are you sure? [y/N] y
Task removed for 2020-06-30
```

#### flags

- `--date DATE` to target a specific date
- `-y` to skip confirmation prompts
- `-n NUMBER` if you already know the index of a task/note you wish to delete
- `--all` to remove all tasks/notes

### `view` command

Use this to view different pieces of information.

**view your update for today**

```
$ supd view update
DID on Monday, June 29:
  1: Task one
  2: Task two

PLAN for today, Tuesday, June 30:
  Complete task foo.

NOTES
  1: Reminder: I will be gone on Friday
```

**view just your plan for today**

```
$ supd view plan
Complete task foo.
```

**view what you've done today**

```
$ supd view done
1: Did a thing
2: Did a second thing
3: Got stuck on a thing
```

#### flags

- `--reverse` to print in reverse-chronological instead of chronological order
- `--json` to print as a JSON string, useful for scripts or parsing with `jq`

Additionally, there are several date selectors for refining the collection of updates to view. Each of these selectors can be passed multiple times, and multiple selectors can be used, for selecting any dates you want.

- `--date DATE`: selects an update from a specific date. `DATE` must be provided in `yyyy-mm-dd` format, e.g. `-d 2020-06-13`.
- `--last N`: selects the last `N` non-empty updates, including today. For example, `-n 3` will select the updates from today, yesterday, and the day before, assuming all three are present in the updates file.
- `--since DATE`: selects all non-empty updates since the given date, including today and `DATE` itself.
- - `--xdate DATE`: removes the update from a specific date from the selection.
- `--xlast N`: removes the last `N` non-empty updates from the selection, including today.
- `--xsince DATE`: removes all updates since the given date from the selection.


### `edit` command

Open the updates file in an editor for manual editing.

```
$ supd edit
```

#### flags

- `--editor COMMAND_OR_PATH` to set the editor to use
