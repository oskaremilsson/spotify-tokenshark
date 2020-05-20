# spotify-controller

### What is this?
This is, soon _(this is super early in the project)_, an `access token` lending service for **Spotify API tokens**.

### Why?
While using _my_ phone as the driver for example Android Auto. I want _someone_ to be able to control the music on _my device_.
_Someone_ should not need to use _my_ phone that's tethered to the car. **Let's lend them my `access token`**.

### What about Spotify on the car display?
Yes, it's possible to control from there.
Maybe even play/pause or skip songs from the wheel.

But **browsing playlists**, **choosing artists** or **searching for music** is often a real pain on those screens.
It would be much nicer experience for _someone_ to be able to use another device.

### What about Spotify Group Sessions?
We could set up a session each time, it would be nice to just setup once.
It's great for other scenarios.

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
