---
title: "Kinomotion"
brief: "I built an app to make cinemagraphs on the go (with Android). This was my hack at Angelhack Spring 2013 which won me a place into their 'hackcelarator'"
type: project
date: 2013-06-22
thumbnail: "/images/logo-kinomotion.png"
---

Kinomotion is an android app that lets you create cinemagraphs. Take a video, select regions to block motion and export as gif!

I took part in a local hackathon in Bangalore to build this. We had 24 hours to build this from scratch.

Being an ardent OpenCV fan, I thought I'd install the library on my phone and get done with the demo and a couple of hours. I was hoping for some sleep at night as well!

But things don't always go as expected - things kept crashing because of this OpenCV dependency. After spending hours setting up OpenCV, I decided to switch. The "core" technology here was just some image blending applied over multiple frames. This could be achieved with much simpler toolkits. No need for OpenCV (which is a beast of a toolkit).

I ended up using Android's native Canvas API to achieve this. This turned out to be a really good decision! The Canvas is hardware accelerated by default - I got insane performance because of this.

I had a prototype to demo a few hours before the demo. I scouted around for some good cinemagraph locations. A few came up:

 * The foosball table
 * Plants rustling in the wind
 * People typing away at their keyboard
 * An auto speeding away while a pedestrian froze in time

I got a few of these and generated gifs of them. And that was my demo - I showed the original vs the cinemagraph. Turns out the judges liked it very much! And I got into the hackcelerator!
