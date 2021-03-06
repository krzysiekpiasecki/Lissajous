# Lissajous

HTTP server deployed on Heroku, generating Lissajous curves animations (GIF), written in GO programming language. Implementation inspired by the example from the book "The Go Programming Language".

<p align="center">
  <br>
  <img src="https://github.com/krzysiekpiasecki/lissajous/blob/master/lissajousCurves.gif" align="center">
  <br>
</p>

## Web Demo

https://lissajous-go.herokuapp.com/?cycles=2&size=200&delay=8&frames=64&res=0.002

Use query string to change parameters

## Build and run locally

```bash
git clone https://github.com/krzysiekpiasecki/Lissajous.git
cd Lissajous
export PORT=8080
go run ./src/lissajous/lissajous.go
```

## Links

[Lissajous curves](https://en.wikipedia.org/wiki/Lissajous_curve)

[Book "The Go Programming Language"](https://books.google.pl/books/about/The_Go_Programming_Language.html?id=SJHvCgAAQBAJ&source=kp_cover&redir_esc=y)

[Open source programming language GO](https://golang.org/)
