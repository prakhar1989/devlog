---
title: "My first Google WebToolkit App"
date: "2010-01-08"
summary: "I made a quiz/strategy game for our technical festival, Quark 2010"
---

For the past couple of weeks, I’ve been experimenting with the Google Web Toolkit v2.0.0. And I just completed my first GWT application today 

An application for Contrive (an event for Quark ’10) had to be made. There were many options, Sliverlight, Flash, Java. I decided to go the JavaScript + HTML route. Quite a bizzare choice for an application so complex. But the GWT made life easy. :D

![My first GWT app](/images/contrive.jpg)

I’m sure I would have torn apart my hair if GWT hadn’t been developed. It needs to check if the connection to the server is established or not. Then it needs to continuously update according to the moves by other players. And it needs to make sure questions are answered in the given time limit. And all these things need to happen on as many browsers as possible.

I’ve got this application working on almost all major browsers without any glitch (except IE :|). It also works on some older browsers with slight graceful degradations in the UI.

The event begins on 15th January… and I’m taking the application through intensive tests, trying to figure out exactly at what load the server fails. We need to get around 600 teams through the eliminations as quickly as possible. That is a HUGE amount of load. And we need to make sure that the thing actually works. But I’m guessing the software can easily handle hundreds of people simultaneously.

Oh, and go register for Contrive right now :D  It’s a completely online event, and you just might win some money  :)
