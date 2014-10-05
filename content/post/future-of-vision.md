---
title: "The future of computer vision"
brief: "The nascent branch of computer vision is filled with a minefield of problems that need to be solved."
summary: "The nascent branch of computer vision is filled with a minefield of problems that need to be solved. What are they?"
date: 2014-10-15
type: post
---

A friend and I had a discussion about computer vision recently. It was a good chat - we thought about some limitations that currently keep us from doing awesome things!

### Stand alone image processing is hard
Running computer vision algorithms on consumer devices is plain simple, hard. Highly parallelized algorithms only run on specialized chips (available only on high-end devices). They also make devices bulky and pricy. Not everyone is going to adopt such hardware.

Sure, doing a few tasks is possible - filters, stabilization, editing. Some tasks require intensive processing - identifying objects, correlating images, etc (at least as of now).

Desktop technology will probably always be a step ahead of mobile. As processors get better, we'll find newer things possible only on desktops. This makes me think there will always be the need to off-load _some_ tasks to the cloud.

### The internet of things
We're almost at the brink of the age of internet of things. Companies like Cisco are already working to setup infrastructure for the billions of devices that will connect to the internet over the next decade.

Cameras will just become one "thing" on the internet. They're to be supplemented with other "things". We're already starting to see this on a small scale:

* Helmets that augment navigation direction on your view
* Cameras stabilizing themselves with gyroscopes
* Cloud based security cameras

In each of these cases, there is computer vision involved - and there's another "thing" that helps supplemnts data. This additional data helps resolve some of the "uncertainities" inherent to computer vision.

Individually, each "thing" is pretty neat - but together they solve an everyday life problem!

In some of these cases, there's a central server involved - with high processing power. This server responds back near real-time.

This leads me to:

### Newer data transfer mechanisms
3G just doesn't cut it anymore. 3g is a power hog. If we're to make a lot of "things", they need to be able to run for a long time before requiring new batteries.

Wifi needs to be made more efficient. The band probably also needs to be made wider (to support a large number of devices in a relatively small area).

Bandwidth needs to be come cheaper. A 4.7GB dvd can hold ~2hours of HQ video. A 140GB data cap would allow only for ~30 hours of video. If you stream this for a day or two, your done. Can there be a media codec designed for use by computer vision applications? 

This suggests a big change in how data is managed - both by the ISP and by the consumer.

### Large scale machine learning
One of the central problems of image processing is identification of objects. As with most algorithms, there is a top-down and a bottom-up approach. They're exactly what you might imagine

* *Bottom-up* is combining individual pixels into identifiable objects (say a couch or a laptop).
* *Top-down* is starting with big chunks of the image and breaking it down till the desired granularity (like a beach scene could be broken into beach > water + sand + palm trees + people > waves + sand castle + palm tree leaves + men + women).

Neither class of algorithms "knows" what they're looking for. They usually just run some mathematical correlation between pixels to sort out what's what. Even humans need to be trained with "X is a couch" statements for years.

Machine learning of this sort on a large scale would produce good "knowledge" about the world. This might be worth a LOT in the near future, I imagine.
