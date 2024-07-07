# Deepwork

## What the flip is this?

Cal Newport recommends keeping a scoreboard measuring time spent working deeply in his book **Deep Work**.
This is an implementation of that scoreboard in the form of a CLI app.

## Install

If you have Go installed, simply run `go install github.com/Salam4nder/deepwork`.


## Usage

Open a new shell and run `deepwork` to start the timer. This will create a `deepwork.txt` file in your home directory. 
When you're done with you work you can interrupt the application (for e.g with CTRL-C) and you work duration will be written to a file.

Run `deepwork -p` to print your recorded durations.
