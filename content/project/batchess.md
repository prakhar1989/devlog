---
title: "Batchess"
date: 2008-06-12
brief: "A remake of the epic chess game, Batchess from the olden days. Complete with animated pieces and cinematic camera movements"
type: project
thumbnail: "/img/logo-chess.jpg"
---

A quite simple game of chess, created from scratch using C++. No rendering engine (like irrlicht, or 3DGS) was used in any way. It used OpenGL for producing 3D graphics, and OpenAL for creating sound.

{{< figure title="The chessboard" src="/img/chess_board.jpg" >}}

The rules of the game are the same as the original chess invented thousands of years back. The only difference is that the “pieces” are alive. The king and queen are actual chinese characters, who walk, attack and die. The pawns are battle droids from the Star Wars movies. The knight is Darth Vader. The rook and bishop are some wierd characters that even I don’t recognize.

And the game is between the red and the blue, instead of the traditional black and white.

{{< figure title="A red knight attacking the blue pawn" src="/img/chess_knight_attack.jpg" >}}

If all this reminds you of the old Batchess, you’re right. That game was the “inspiration” for this one.

### Features
* OpenGL for graphics
* OpenAL for audio
* Uses Quake 2 models (MD2)
* Uses .tga, .bmp and .jpg files to skin models
* Uses Object Oriented Concepts of C++ (e.g: a renderer object handles all display, a chess piece object handles all move validation, etc)
* Implements a Camera Animator class to move the camera over a specified path. Used at multiple places. e.g: to create dramatic kill sequences, turning the tables, etc

{{< figure title="The queen can attack - shown by the red square" src="/img/chess_queen_can_attack.jpg" >}}

### Fin
Overall, a quite basic game programming project. But much of an achievement for me – not using any pre-made game engine – everything made from scratch!


