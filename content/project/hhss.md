---
title: "The Hand Held Step Scanner"
date: 2009-07-10
brief: "My internship project at CSIO, Chandigarh - this was a device based on image processing to help the blind read documents on the go. I helped port certain modules over to OpenCV."
type: project
thumbnail: "/img/logo-hhss.png"
---

The HHSS is a prototype reading aid for the blind. Most such aids at the time were similar to flat-bed scanners. They weren’t portable and were limited to a certain document size. For example - you couldn’t read a newspaper with a flatbed scan device. This project at the Central Scientific Instruments Organization (in Chandigarh) aimed to get rid of this limitation.

### A portable device
The device consisted of two cameras housed like eyes. The device had a button to start scanning and was portable enough to be held with one hand.

### My role
I worked on the image-processing side of things. The code took images from the two cameras and stitched them together. The original code was completely in C/C++. I converted all existing image processing function calls to OpenCV.

Then I implemented the stiching module for combining the two images. At that time, OpenCV did not come with a lot of the image processing good-ness it does now. So everything had to be written from scratch.

After that, I helped modularize the calls to the camera API functions. They were testing with two cameras. So this modularization helped them swap between cameras much more easily.
