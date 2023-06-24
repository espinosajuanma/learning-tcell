# Learning `tcell` Go Lang module

I am a terminal enthusiastic, but I've never made a TUI application. All the tools I make are CLI tools where I provide a set of commands, subcommands or inputs and then I got an output.

It is time that I learn a framework so I can make at least a simple TUI application.

What I want so far is be able to make something for me, I am into the Vim keybindings world, so that is going to be my goal.

## Directory and progress

I don't want to have those kind of repositories where there is a bunch of folders with `step 1`, `step 2` and being redundant to have the code of 1 in 2 as well. We have version control for something. And anyway knowledge is cumulative, the last commit of this repository would be until I stop learning with this project, but it will continue with other projects, I hope.

So I am going to have all at root level and probably make a sub-directory to take some notes about stuff I want to keep closer.

## Challenges

These are some of the challenges I'll be trying to achieve.

- Say `Hello World` on Full Screen
  - Center it
- Use `q` to back to the terminal
  - Disable `Ctrl + C`
- Print a "pop-up" in the middle of screen
- Print two panes vertical
  - One white, the other magenta
- Switch panes using `h` and `l`
  - Magenta one should be the selected one
- Print a list of items in one pane
- Press `Enter` to select the handle the magenta pane
  - Switch pane status to active -> Change color to green
- Terminal resize event - keep it at full screen
