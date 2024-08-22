# 🌟 lumenserver

This project implements a robust custom HTTP server in Go, featuring a modular structure, basic routing, middleware support, and request handlers.

## What's this all about?

So, I wanted to dig deeper into how web servers work, and I figured the best way to learn was to build one from scratch. This project is the result of that curiosity - a basic but functional HTTP server that does a bit more than just say "hello world".

Here's what I've managed to pack into this thing:

- A server that actually, you know, serves stuff
- Some basic routing (because we all love to navigate)
- A middleware system (for all that behind-the-scenes magic)
- A structure that doesn't make me want to pull my hair out

## Getting This Thing Running

First things first, you'll need Go installed. I used version 1.16, but anything newer should work too.

1. Grab the code:
   ```
   git clone https://github.com/MuxN4/lumenserver.git
   ```
2. Head into the project folder:
   ```
   cd lumenserver
   ```
3. Build it:
   ```
   go build
   ```
4. Run it:
   ```
   ./lumenserver
   ```

If all goes well, you should see the server running on `localhost:8080`. Exciting, right?

## 🔨 Let's Break It... I Mean, Test It!
Alright, now that you've got this bad boy up and running, let's see what it can do! Time to poke it with a stick (or in this case, curl commands) and see how it reacts.

## Curl-ing Up with Your Server
Open up a new terminal window. Your server's feeling lonely in that other one, so let's give it some company:

1. Say hello to your home page:
   ```
   curl http://localhost:8080/
   ```
If it replies with "Welcome to our custom HTTP server!", you're golden!

2. Ask it for the time (because why not?):
   ```
   curl http://localhost:8080/time
   ```
It should spit out the current time. Fancy, huh?

3. Get a list of users (imaginary friends, perhaps?):
   ```
   curl http://localhost:8080/users
   ```
You should get a JSON array. Empty? Time to make some friends!

4. Create a new user (finally, a friend!):
   ```
   curl -X POST -H "Content-Type: application/json" -d '{"name":"Eve"}' http://localhost:8080/users
   ```
If Eve appears in JSON form, congrats! You've made a digital friend.

5. Try to visit a non-existent page (let's get lost):
```
curl http://localhost:8080/narnia
```
"404 page not found"? Perfect! Your server isn't making up imaginary places.

## Middleware Madness
Now, let's check if our middleware is doing its job (or just middleware-ing through life):

LoggingMiddleware: Check your server logs. If you see messages popping up like a chatty neighbor, it's working!

RecoveryMiddleware: Time to cause some chaos! Temporarily add this to your TimeHandler:
   ```
   func TimeHandler(w http.ResponseWriter, r *http.Request) {
      panic("It's panic time!")
   }
   ```
   
Now, try to get the time again:
   ```
   curl http://localhost:8080/time
   ```
If you see "Internal Server Error" instead of your server curling up in a corner, the RecoveryMiddleware is doing its job!

## Make It Your Own

Feel free to tinker with this thing. Add new routes, cook up new handlers, or beef up security with some middleware. Sky's the limit!

## Wanna Contribute?

Found a bug? Think you can make this even better? Go for it! Fork the repo, make your changes, and hit me with a pull request. Just be nice, okay?

## Legal Stuff

This project is under the MIT License. Basically, do whatever you want with it, but don't blame me if something goes wrong.

---

That's it! Hope you have as much fun playing with this as I did building it. If you've got questions, feel free to reach out. Happy coding! 🚀
