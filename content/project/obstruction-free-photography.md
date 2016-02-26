---
title: "Obstruction Free Photography"
date: 2015-12-12
brief: "Inspired by all the obstacles when trying to take photographs in Bangalore, I try implementing this computational photography paper and run it on images taken around the CMU campus."
type: project
thumbnail: "/images/logo-ofp.png"
---

SIGGRAPH 2015 had a paper on [obstruction free photography](https://sites.google.com/site/obstructionfreephotography/). The problem statement was that there are several real-life situations where your subject is obstructed (by a fence, a window pane, etc).

The key ideas were take a short video (about 5 seconds) with slight changes in the camera position. Thus, your camera captures pixels behind the obstruction. Then, calculate the relative positions of the frames. Using these relative positions, the image is split into a foreground and background layer (the foreground moves a lot more compared to the background). Finally, run it through an optimization function. This function rewards smooth changes and penalizes harsh ones. Essentially, it "picks the right colour for each pixel". The implementation details were a bit tricky.

I used C++ with OpenCV and DLib for this project. The runtimes for the completely unoptimized source code were about 10 minutes per iteration of optimization (for five 640x480px images). I fixed any unnecessary memory copies to improve performance. I think this code can easily be parallelized on a GPU. If a GPU is not available on a device, multiple threads working on different parts of the image would also improve performance.

## Results
### The Cathedral of Learning through a staircase railing
![The Cathedral of Learning through a railing](/images/ofp-cathedral-learning.png)
![The Cathedral of Learning through a railing (fixed)](/images/ofp-cathedral-learning-fixed.png)

### CMU's Smith Hall
![CMU's Smith Hall](/images/ofp-smith-hall.png)
![CMU's Smith Hall (fixed)](/images/ofp-smith-hall-fixed.png)

## Issues

### When do I stop the iterations?
Each iteration is supposed to produce a better result. If you run it for too few iterations (say, just 1), you end up with something that is mostly cleaned up but has some artifacts. If you run it for too long (say, 20 iterations), over optimization creeps in and produces artifacts that look like JPEG compression. A good heuristic might be to always run it for 5 iterations but that's wasted computation. Some kind of a quality check metric would be very useful here.

![Over optimization produces bad results](/images/ofp-over-optimize.png)

### Image alignment is not always perfect
I've noticed that the input image sequence often needs hand-holding when figuring out the relative orientations between different frames. I suspect my code might have bugs - so I'm still looking. But if the relative orientations are not estimated correctly, the results look terrible. One way to fix this would be to use a phone's IMU to have an initial estimate of phone's position. This may require some calibration, though (which is not a good idea).

![Image misalignment](/images/ofp-image-misalign.png)

## Conclusion
I feel my code would use a lot more fine tuning to get it working just right. I have not done justice to the amount of effort that went into this paper's implementation by the authors. I've found several things that can go wrong when working with a real-life dataset (we captured the videos ourselves!). I've found some potential techniques that can be used to fix those issues. Overall, this was a good project to work on!
