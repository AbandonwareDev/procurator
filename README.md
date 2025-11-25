# Procurator

My GoLang project manager/helper TUI 

![](./.media/vhs.gif)


## Description

The idea is to use simple text editor in separate terminal pane (e.g. tmux) to have kinda IDE experience.

It's TUI which automatically run `go mod tidy; go vet` if any file in directory changes and shows result.
Apart from that, here is regular quick commands like `git add .`, `git push`, `go fmt`...


### Installation

`go install github.com/AbandonwareDev/procurator@latest`

### Usage

`procurator` in git repo directory 

## TODO

 - add other quick commands 
 - check all TODO in code (like `go vet` only on `.go` files change) and write here
 - clean code
 - add programming language autodetect
 - add other programming language support
 - make quick commands configurable/modular (like in project file with commands) (or presets/user-defined)








