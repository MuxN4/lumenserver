# lumenserver

This project implements a robust custom HTTP server in Go, featuring a modular structure, basic routing, middleware support, and request handlers.

## What's this all about?

So, I wanted to dig deeper into how web servers work, and I figured the best way to learn was to build one from scratch. This project is the result of that curiosity - a basic but functional HTTP server that does a bit more than just say "hello world".

Here's what I've managed to pack into this thing:

- A server that actually, you know, serves stuff
- Some basic routing (because we all love to navigate)
- A middleware system (for all that behind-the-scenes magic)
- A structure that doesn't make me want to pull my hair out

## Project Anatomy

Here's how I've organized this beast:

```
custom-http-server/
├── main.go (where the magic begins)
├── server/
│   ├── server.go (the heart of our operation)
│   ├── router.go (traffic control)
│   └── middleware.go (our bouncer)
└── handlers/
    └── handlers.go (does what it says on the tin)
```

## Getting This Thing Running

First things first, you'll need Go installed. I used version 1.16, but anything newer should work too.

1. Grab the code:
   ```
   git clone https://github.com/yourusername/custom-http-server.git
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


## Make It Your Own

Feel free to tinker with this thing. Add new routes in `main.go`, cook up new handlers in `handlers/handlers.go`, or beef up security with some middleware. Sky's the limit!

## Wanna Contribute?

Found a bug? Think you can make this even better? Go for it! Fork the repo, make your changes, and hit me with a pull request. Just be nice, okay?

## Legal Stuff

This project is under the MIT License. Basically, do whatever you want with it, but don't blame me if something goes wrong. Check out the [LICENSE](LICENSE) file for the boring details.

---

That's it! Hope you have as much fun playing with this as I did building it. If you've got questions, feel free to reach out. Happy coding! 🚀