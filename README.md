`supd` (**s**crum **upd**ate) is a CLI app I made for myself to track my daily updates for work. Every day I like to track my plan for the day and the things I actually did that day.

This app is under development. Progress to a v1 release is being tracked in [this issue](https://github.com/awmottaz/supd/issues/4).

- [Installation](#installation)
  - [Download](#download)
  - [Build](#build)
- [Usage](#usage)
  - [`help`: App documentation](#help-app-documentation)
  - [`version`: Version information](#version-version-information)
  - [`edit`: Edit the updates file](#edit-edit-the-updates-file)
  - [`plan`: Set and view your plan for today](#plan-set-and-view-your-plan-for-today)
  - [`did`: Track completed tasks](#did-track-completed-tasks)
  - [`note`: Take other notes](#note-take-other-notes)
  - [`show`: View updates from selected dates](#show-view-updates-from-selected-dates)
    - [Update Selection](#update-selection)
    - [Print Order](#print-order)
    - [Examples](#examples)

## Installation

### Download

You can download a pre-built binary from the [releases](https://github.com/awmottaz/supd/releases).

### Build

Alternatively, clone this repo and install with

```sh
make install
```

You must have Go 1.11+ installed with modules enabled for this to work. [See installation instructions for Go.](http://golang.org/doc/install.html)

## Usage

Run `supd help` for current usage instructions and supported functions.

**The following usage instructions are not finalized.** These instructions currently serve as a sketch of the app design for my own reference. After the v1 release, these will become the documentation instead of a design sketch.

### `help`: App documentation

Use the `help` command for usage instructions.

```
$ supd help
Usage:

    supd <command> [arguments]

Summary:
    supd is a simple command line app to help you track your daily
    updates, the plan you have for the day and the tasks you actually
    completed. These updates are all saved in a single JSON file (see
    "Environment" below).

Environment:

    If the SUPD_FILE environment variable is set, then this file will
    be used for the updates file. The "-f" global flag overrides this
    setting. If neither are set, "$HOME/supd.json" is used.

Options:

    -f    set the path for the updates file

Commands:

    did        document a task completed today
    edit       open the updates file for editing
    help       show this usage information
    note       save a note from today
    plan       set/view your plan for today
    show       print selected updates
    version    display detailed version information

Use "supd help <command>" for more information about a command.
```

### `version`: Version information

Call the `version` command for more detailed information.

```
$ supd version
supd build information

VERSION: 0.1
HASH:    624801389b0fc71334db06dfeacc63ddead6609e
DATE:    2020-06-13T14:22:13Z
URL:     https://github.com/awmottaz/supd/releases/0.1
```

### `edit`: Edit the updates file

The `edit` command opens the updates file for you to edit manually. If set, it will use the `EDITOR` environment variable. Otherwise, it tries to use `vim`.

```
$ supd edit
Opening /home/user/supd.json with vim...
```

### `plan`: Set and view your plan for today

Set your plan

```
$ supd plan "I will complete task foo"
plan written to /home/user/supd.json
```

View your plan

```
$ supd plan
I will complete task foo
```

Message if there is no plan set for today

```
$ supd plan
**no plan set for today**
```

### `did`: Track completed tasks

Add a completed task for today

```
$ supd did "completed part 1 of foo"
task written to /home/user/supd.json
```

View completed tasks for today

```
$ supd did -l
1: completed part 1 of foo
2: did another thing
```

Remove a task from the completed list. This re-numerates the list.

```
$ supd did -d 1
deleted task: "completed part 1 of foo"

$ supd did -l
1: did another thing
```

Provide the `-d` flag multiple times to delete multiple tasks.

```
$supd did -d 2 -d 3
deleted task: "task 2"
deleted task: "task 3"
```

### `note`: Take other notes

Just like tracking tasks, you can track arbitrary notes. This has the same options as `did`.

```
$ supd note "this is a note"
note written to /home/user/supd.json

$ supd note -l
1: this is a note
2: this is another note

$ supd note -d 1
deleted note: "this is a note"
```

### `show`: View updates from selected dates

Without any arguments, this pretty-prints today's update.

```
$ supd show
* Update 2020-06-13 *
PLAN
    create the `show` command for supd app
DID
    1: created the sub-command namespace
    2: added ability to print today's update
NOTES
    1: the text/template package is really nice to use
```

Pass the `-json` flag to print in JSON format. Useful for scripts or to find specific data with `jq`, for example.

```
$supd show -json
[
  {
    "date": "2020-06-13",
    "plan": "create the `show` command for supd app",
    "did": [
      "created the sub-command namespace",
      "added ability to print today's update"
    ],
    "notes": ["the text/template package is really nice to use"]
  }
]
```

#### Update Selection

The following options can be used to select other updates. These options can be combined and used multiple times to make any selection of updates you want. This selection is resolved to a list of updates present in the updates JSON file, and each of them is pretty-printed in chronological order.

- `-d DATE`: selects an update from a specific date. `DATE` must be provided in `yyyy-mm-dd` format, e.g. `-d 2020-06-13`.
- `-n N`: selects the last `N` non-empty updates, including today. For example, `-n 3` will select the updates from today, yesterday, and the day before, assuming all three are present in the updates file.
- `-s DATE`: selects all non-empty updates since the given date, including today and `DATE` itself.

You can also prefix each of these selectors with `x` to exclude updates from selection:

- `-xd DATE`: removes the update from a specific date from the selection.
- `-xn N`: removes the last `N` non-empty updates from the selection, including today.
- `-xs DATE`: removes all updates since the given date from the selection.

#### Print Order

By default, all selected updates are shown in chronological order. Use the `-r` flag to show in reverse chronological order.

#### Examples

Select a number of dates. In this example, only the updates from June 5th and June 8th are shown because there was no update for June 7th. This example also demonstrates that updates will be shown in chronological order, regardless of how

```
$ supd show -d 2020-06-08 -d 2020-06-07 -d 2020-06-05
* Update 2020-06-05 *
...
* Update 2020-06-07 *
...
```

Select the last 2 updates and the update from May 19th.

```
$ supd show -n 2 -d 2020-05-19
```

Select all updates in May.

```
$ supd show -s 2020-05-01 -xs 2020-06-01
```

Select the update five updates ago.

```
$ supd show -n 5 -xn 4
```
