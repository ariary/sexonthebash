
# sexonthebash 🍹

Stealth shell input and output listeners. Differerent approach for keylogging with Shell. 

It provides 2 utilities imitating a shell to capture input and output:
 - `sexonthebash`: capture output of **bash** commands
 - `shellonthebeach`: capture input and output of **/bin/sh** command (the same thing as `sexonthebash` with more vodka)

You are free to do what you want with these data afterwards (DNS exfiltration, etc).

Could be used to:
- spy on other users on the same machine, for CTF for example
- get an interactive `/bin/sh` interactive (ie. With arrow keys, backspace etc)
	

| 🚧🚧*Under construction*.🚧🚧  |
|:------------------------------------------------------------------------------------------------------------------:|
***Just for educational purposes, do not use it if against someone in real-life if you do not have permission***|
|*Any idea, criticism, contribution is welcome*|

	

##  Table of contents

 - [🔦 Idea](#-idea)
 - [💺 Installation](#-installation)
 - [🚀 Usage](#-usage)
 - [💭Limits/improvements](#limitsimprovements)
	
## 🔦 Idea

Provide a different approach for keylogging (log only command on `bash.`/`sh`). But some juicy information could be grabbed from it.

The aim of `sexonthebash` is to be launched under a bash session by any means and to imitate bash behaviour to lure the victim in a way that the victim doesn't feel like they're somewhere other than a bash session.

Idem for `shellonthebeach`, to be launched in place of a `sh` tty

## 💺 Installation

### Prerequisite

 - golang installed
 - make installed
 - be ill-intentioned
 
 Clone the repo and download the dependencies locally:

```
git clone https://github.com/ariary/AravisFS.git
make before.build
```

### sexonthebash
	make build.sexonthebash

### shellonthebeach
	make build.shellonthebeach

## 🚀 Usage 

Launch it, and exec your usual shell commands. If you see nothing suspect it is perfect.



## 💭Limits/improvements

📬 **Please tell me if you see some bugs,improvments etc!** (with issue, PR etc)

### sexontheBash

 Two important aspect of a BASH are **the prompt string** and its **interactiveness**. By piping `os.Stdin` we couldn't have these two features (=> we could never lure someone onthinking he is on a bash session so). This is why we only capture `os.Stdout`

### shellonthebeach

The command outputs are not printing in real-time
