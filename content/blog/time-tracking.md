---
title: "Time Tracking with Toggl and Tasker"
brief: "Time is money they say. Know the exact worth of your time down to the dollar."
date: 2018-01-05
type: blog
---

Time is money they say. Knowing the worth of your time as a dollar amount per hour can be eye opening. Often, time is billed for working hours only - however, there can be sickness, vacations, family time and side projects. Factoring these into your equations could lead a different view. Since this is just the start of the year, it's a good excuse to start tracking these things. In fact, this aligns with the idea of the [quantified self](https://en.wikipedia.org/wiki/Quantified_self).

## Toggl
[Toggl](https://toggl.com/) to track time is almost exactly what's needed for this. Toggl lets you define projects ("Meetings", "Office", "Sleep") and lets you track time for each of them. The tool seems to be primarily geared towards freelancers who work for clients or have an hourly rate. I don't use those settings just yet - but I can imagine a few scenarios where this might be useful.

The webapp and mobile app allow you to start timers and edit your records. However, this is somewhat cumbersome - we'll see a better way later.

## Tasker
I consider [Tasker](http://tasker.dinglisch.net/) a visual programming language app. It lets you work with dates, time, location, other apps, launch REST API calls - all from within a visual environment. It even lets you write javascript code if these tools are not sufficient.

One of the nice things about Tasker is it lets you execute tasks and execute them from other tasks. This lets you build hierarchies of tasks with shared variables - almost like subroutines.

## Tasker + Toggl
Using [Toggl's REST API](https://github.com/toggl/toggl_api_docs/blob/master/toggl_api.md), it is possible to build tasks that [start and stop timers](https://github.com/toggl/toggl_api_docs/blob/master/chapters/time_entries.md). It is possible to request the currently running timer, stop it and start a new timer.

Using such a scheme, you can devise tasks that run timers 24x7. Here are some examples I use:

### Sleep
This task accomplishes the following:

- Stop any currently running timer
- Start the sleep timer
- Activate the sleep tracking app
- Set an alarm to an appropriate time
- Turn of all wifi lights in the house

### Commute

- Stop any currently running timer
- Start the commute timer
- Turn off Wifi
- Turn on Location
- If it is expected to rain, show a popup

Some more examples of timers:

- Side project
- Nap
- Office
- Meetings
- Social
- Fitness

Once you have a bunch of such tasks in place, you can pretty much designate each second of your day to certain activities. One approach to improve your habit of always having a timer running is to export these tasks as individual apps. Tasker lets you export tasks with their own icon, package name and version number. Since these apps just make REST calls, they don't have a UI. They just execute and do their thing - making it convenient.

## Issues
The primary issue is with Toggl's Android app is that it does not respond to REST calls in the background. In fact, if you open the app and try to refresh the timer list, it crashes. There's probably some state in the Android app that expects only a specific timer to be running at any time. Restarting the app usually fixes it though.

Another issue is with REST calls. This could be a Tasker issue or Toggl's API issue. The API call just takes too long once in a while. It leaves you hanging - did the API token change? Did you miss a corner case in the task (when parsing dates, for example).

## Conclusion
I've been using this system for a couple of weeks now. However, most of my time has just been personal time from the end-of-year vacations. As the weeks progress, hopefully I'll have more unbiased data about my use of time. Hopefully, I'll be able to gain some value out of it.
