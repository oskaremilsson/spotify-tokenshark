# spotify-controller

### What is this?
This is a token lending service for Spotify API tokens.

### Why?
While using my phone as the main driver for Android Auto, I want my girlfriend to be able to control the music.
For this she should not need to use my phone that's tethered to the car.

### What about Spotify on the car display?
Yes, it's possible to control from there.
I can also play/pause and skip songs from the wheel.
But browsing playlists, choosing artists or searching for music is a real pain there while driving.
It would be much nicer experience to be able to use another phone.

### What about Spotify Group Sessions?
Sure, we could set up a session each time. 
I don't wish to this setup every time we're out for a drive.

### So what is the real goal?
I want the possibility to give certain people access to use my Spotify API token in order to control my music.
This service will be that broker.

### Cool, but how?
The idea is quite simple _(maybe)_:
* User A store their `refresh token` in the service.
* User A gives User B the right to use it
* User B requests an `access token` for User A by providing a personal `access token` to Spotify API 
* Service validates who User B is against Spotify API _(instead of having own auth)_
* Service validated that User B indeed have the rights to get a token
* Service send User B an `access token` issues for User A
* Finally User B uses that token towards Spotify API controlling music for User A

### What about UX?
Yes, there will be some kind of client in a separate repo. Probably 

---------------

_I'm new to Go, so keep that in mind while reading the code._
_HUGE thanks to @sikevux for answering my stupid questions._