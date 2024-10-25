# [Let's Go](https://lets-go.alexedwards.net/sample/00.00-front-matter.html)

## Why this book?


--- 

> Module paths for downloadable packages
If you’re creating a project which can be downloaded and used by other people and programs, then it’s good practice for your module path to equal the location that the code can be downloaded from.

> For instance, if your package is hosted at https://github.com/foo/bar then the module path for the project should be github.com/foo/bar.

---

### Creating the handler, router, and web server

```
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Snippetbox"))
}
```

The home handler is just a regular function that takes in two parameters, the ResponseWriter interface and Request pointer struct, and sends the byte slice to the writer.

Here are the links for the http library, the interface [ResponseWriter](https://pkg.go.dev/net/http#ResponseWriter), and the pointer struct [Request](https://pkg.go.dev/net/http#Request)

