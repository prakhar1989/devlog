---
title: "Contrive"
date: 2010-01-10
brief: "An online realtime strategy/quiz game built with GWT I built for my university's national technical festival. This was my first foray into the realm of web applications."
type: project
thumbnail: "/img/logo-contrive.jpg"
---

Contrive was an online quiz + strategy game held at the innovation festival, Quark 2010 at BITS-Pilani, Goa. I was to make the software for the entire thing. It seemed like a daunting task initially, considering the required functionality. But after quite a bit of effort, I was able to create the software that worked perfectly.

### What is Contrive?
Contrive was an online science quiz, with a twist of strategy. You play against three other teams. You start off at one corner of the grid. The goal is to reach the center of the grid. To do that, you need to answer a series of questions. Once you answer a question, you can move in one of the four directions.

{{< figure title="The contrive grid" src="/img/contrive-board.jpg" >}}

Now, the colour of your location decides the genre of the question. Since this is a science quiz, there are four different colours: yellow, red, blue and green. Each representing physics, mathematics, biology and chemistry. So, as you move, you can decide which type of question you want next.

Because three other teams are also playing, they’ll be doing the same thing. Answering questions and moving towards the center.

That’s the basic part. Now for the strategy part. Positions where you answer correctly are blocked to everyone else. This way, you can block other players… essentially “checkmating” them. Or you can circle around the center to make it inaccessible to other players.

### Technical challenges
The very first challenge was to decide which programming language to use. Options being considered were Silverlight, Flash, Java applets, and JavaScript. I decided to use GWT. That is, the Google Web Toolkit. This is a set of tools that compile Java into JavaScript. Neat.

The next challenge was to implement this web application. It was required that at least 20 teams play simultaneously during a 2 hour period. Now that turns out to be a lot of load for the server. There’s the constant ping. And then other timers, polling for the position of other teams, downloading questions, updating moves.

Things had to be made as efficient as possible. Contrive was hosted on the Quark website. So it had a lot of traffic already. All this pinging/etc could slow down the entire thing. So, SQL queries were optimized. The code was made as short as possible. The “output” the Contrive APIs gave were made as precise and concise a possible.

### Final results
Up till now, the application has been functioning perfectly. It has been able to support 20 teams. We haven’t tried crossing the limit, at least as of now. Though I’m quite sure it will scale well even if around 100 teams login simultaneously. But then the Quark website would go down

Oh, and after this application, I feel creating something like Facebook should be simple as well :P
