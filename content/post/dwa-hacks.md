---
title: "Organizing the DreamWorks Animation Hackathon"
brief: "We asked them to implement quite a bit - simluation, path finding, webgl and sockets!"
summary: "We conducted a hackathon at IIIT Bangalore. Simulate a flock of fish and render the results on screen (or in a browser!)"
date: 2015-07-30
type: post
---

On the 28th and 29th of this month, DreamWorks Animation, Bangalore organized a hackathon at IIIT Bangalore. Being part of the organization team, I learned some good things about the other side of the table. This blog post is more about the technical aspects of the event.

The problem statement was to implement a simulator for a flock of fish. We provide a map of obstacles through which the fish must find their way and reach the destination. This simulation must be streamed to a web frontend in realtime. Teams of up to 4 members were formed. As you can imagine, this task stretched the participants a bit - in a good way!

![Implement a 2D fish flocking algorithm](/images/fish-flock-2d.jpg)

The technologies involved are very diverse - [OpenVDB](http://www.openvdb.org/) (DreamWorks' open source volume data format), WebGL, WebSockets, physics engines, interprocess communication, servers and a lot of other things. The participants were able to pick up these technologies and split tasks amongst themselves to accomplish the larger goals - this was super impressive! There was no limitation on the language to be used - though we recommended C/C++ for OpenVDB.

While devising the hackathon, we implemented the problem statement ourselves just to get a sense of the pain points. Here's what our architecture for the problem statement looked like.

### Our approach

The simulation solver was written in C++. This was to ensure we do not run into performance issues because of lot of fish. The solver needs to find a path and also keep the fish constrained under the rules of flocking.

We created a viewer for C++ to see how the flocking was doing. The viewer was quite straightforward - a bunch of `gl*` commands to read a data structure and render it onto the screen. We had a start position, end position and each fish that had to be rendered.

Our end goal was to make the fish show up in a web browser. One method could be to render an image in the backend (maybe with OpenGL) and transmit that to the browser. However, for obvious reasons, this would be quite wasteful of bandwidth.

A better approach would be to just transmit the position of the fish and render the fish in Javascript. This is the route we chose. To transmit, we could have written a web server in C/C++ but due to familiarity, we decided to go with Python.

We used Tornado to create a simple web server that can serve some static files and be the mediator between the C++ simluation process and the browser. We used a Named Pipe to communicate between the C++ and Python process. For serialization, we used Protobuf.

For communication between the browser and the webserver we used a websocket. The Tornado server wouldl simply forward data packets from C++ to the websocket. There was a slight difference between the C++ packet received and the websocket packet transmitted (we add some stuff to help with WebGL).

Interestingly, we found that the C++ process was able to generate ~5000 packets of data per second. Each packet was between 10-20 bytes of data. Python was able to read and forward ~3000 packets of data per second. However, Javascript was only able to process ~500 packets of data per second. This was sufficient for this project - but for future projects, we may have to take into account this slowdown.  Our theory is Javascript takes a while to parse the data packet. I am not sure if that is true.

We went a step further and let the user upload a VDB file from the browser. This required additionally transmitting an image of the VDB (which we rendered on the fly).

### Results

The teams completed the problem statement to varying levels of completeness after the designated 24 hours. We were able to shortlist five teams who were neck-and-neck with each other (we had asked them to submit gifs of their results).

One after the other these teams presented their results and approach to solving the problem in front of the entire audience. A hard decision - but the judges (ie, us!) reached consensus and two teams were picked as the winner and runner up!

[![The hackathon winners!](/images/hackathon-winners.png)](/images/hackathon-winners-big.png)

### Aftermath
The two days leading up to the hackathon and the two days of the hackathon - were intense and fun! The day afterwards felt a bit uneasy with thoughts like "I don't have anything to do" and "What do I do?".

It would not have been possible to pull off something of this scale without the awesome team that we had! 

[![The team](/images/hackathon-post.png)](/images/hackathon-post-big.jpg)
