# Pomodoro tracker

I use a very simple pomodoro clock to stay focused and to keep track of
how long I've been working. It is a simple bash script, which outputs
the pomo time to stdout. 
<https://github.com/kmjayadeep/scripts/blob/master/pomo>

It just have 2 commands, `pomo start [time]` and `pomo clear`. Running
just `pomo` outputs the timer to stdout. The script is used in the
polybar config to show at the top bar.

This project is a simple api server to track the pomodoro clock timings and
do some analysis on top of the data - like hours worked per day, hours
worked per month etc.

In the pomo script, a simple curl can be added to call the endpoints in
this prject to indicate start and stop of pomo.

Another *motive* of this project is for me to gain some more experience
building APIs in GO. I have no plan to make this project feature rich or
customize to work with different scenarios. Feel free to play with it,
and use at your own risk!


## TODO

more details about the project and instructions to use will be added soon
