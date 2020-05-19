# spotify-controller

### What is this?
This is a hobby project with the aim fix a need that I've had while connecting my phone to Android Auto.

### What need?
Well, I use my phone as the main "driver" to Android Auto. But then I want my girlfriend to be able to control the musik. Ultimately she should not need to use my phone that's tethered to the car.

### What about Spotify Group Sessions?
Yes, you can set up a session each time to let someone control the music. 
While this is true I don't wish to so this every time.

### So what is the real goal?
I want to give certain people constant access to use my API tokens in order to control my music. This service will allow her to login to a client and "inpersonate" me. 


### Cool, how does it work?
The idea is quite simple _(maybe)_:
* User A upload their refreshtoken to the service.
* User A gives User B access to use it
* User B requests access token for User A
* Server validated that User B have the rights to do so
* User B uses that token towards Spotify API controlling music for User A

---------------

_I'm new to Go, so keep that in mind while reading the code._

_HUGE thanks to @sikevux for answering my stupid questions._