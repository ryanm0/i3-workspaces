# i3-workspaces

![screenshot](http://pix.toile-libre.org/upload/original/1390061002.png)

A little utility to display workspaces in i3. To be used with
[bar](https://github.com/raedwulf/bar/tree/xft-full-utf8), the
lightweight bar.

i3-workspaces listens on events for workspaces changes; as soon as there
is a change, it will ask i3 for all its workspaces and send it properly
formatted to the stdout which is read by bar afterwards.
The format is hardcoded to display workspaces in the center of the bar,
with the focused workspace being underlined.

## Dependencies
- github.com/proxypoke/i3ipc

## Usage
```shell
> go build
> ./i3-workspaces | bar
```

## Licence
CC0
