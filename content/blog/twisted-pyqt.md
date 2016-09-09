---
title: "Twisted + PyQt"
date: "2014-05-22"
tools:
  - twisted
  - pyqt
summary: "Lessons on creating a GUI with two event loops. Twisted for interacting with a webservice and PyQt for a GUI"
brief: "I favored the reactor over signals/slots"
type: blog
---

I've been using PyQt and Twisted to make a GUI for a webservice. The webservice was entirely written with Twisted. The UI is written in PyQt.

The only reason I chose this combination was to be able to reuse the codebase from the service in the GUI (you know, standard calls like fetch this, run a remote procedure, do X exactly like the webservice does, etc).

Turns out it was a great idea. The Twisted reactor makes the code much more readable and helps skip the signal/slot mechanism of PyQt.

### Multiple threads
As with any UI, there are multiple threads carrying the heavy weight - they execute the blocking calls. Things like fetch XYZ from the webservice. If you don't have threads, the UI would be blocked and that's not a good thing.

With threads, everything is usable - however, you need to display progress for whatever the thread is doing. This means, multiple threads updating the UI (progress bars and what not) - this is not good.

The Qt way of solving this problem is to use signals and slots. They work - but can we do better?

### The Twisted Reactor
This is where the reactor steps in. If you initialize the reactor in the main UI thread, you can queue messages in the reactor - from different threads (using the callFromThread method). This means, you don't need a signal/slot mechanism - just the reactor is sufficient. This was a big ah-ha moment!

**tl;dr**: I used the Twisted reactor to communicate across multiple threads
