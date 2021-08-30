
# sexonthebash 🍹

<div align=center>
<hr size=15px color="ff5733" width=75%>
	🍾		🧊
	<h3><i>TL;DR</i></h3>
No talks needed,  drinks are on me. I want to use it [<a href="#-usage">🚀</a>]

<hr size=15px color="ff5733" width=75%>
</div>

Stealth shell input and output listeners. Differerent approach for keylogging with Shell. 

It provides 2 utilities imitating a shell to capture input and output:
 - `sexonthebash`: capture output and input of **bash** commands
 - `shellonthebeach`: capture input and output of **/bin/sh** command (the same thing as `sexonthebash` with less vodka)

You are free to do what you want with these data afterwards (DNS exfiltration, etc).

Could be used to:
- spy on other users on the same machine, for CTF for example
- get an interactive `/bin/sh` interactive (ie. With arrow keys, backspace etc)


| ***Just for educational purposes, do not use it if against someone in real-life if you do not have permission*** |
|:------------------------------------------------------------------------------------------------------------------:|
|*Any idea, criticism, contribution is welcome*|



## 🔦 Idea

Provide a different approach for keylogging (log only command on `bash`/`sh`). But some juicy information could be grabbed from it.

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

### Launch sexonthebash
In a bash session:
```
./stealth_bash.sh
```
Launch the command `exit` to exit the listener and find the log in `.input.log` and `.output.log`

### Launch shellonthebeach
In a sh session:
```
./stealth_shell.sh
```
Launch the command `exit` to exit the listener. Nothing is done with captured commands from now

## 💭Limits/improvements

📬 **Please tell me if you see some bugs,improvements etc!** (with issue, PR etc)

### shellonthebeach

- The command outputs are not printing in real-time
- Nothing is done with captured commands

### sexonthebash

Be able to store (or send to a remote place) the captured output and input in a near real time (ie. A way to not waiting anymore that the user has finished his bash session to get those data)
