# goldmark-links

[Goldmark](https://github.com/yuin/goldmark) link replacer extension.

support link add `rel = "nofollw"`, `target = "_blank"` and more customize attribute.

## code

```go
var source = []byte(`[website](https://fungo.dev)
[source](https://github.com/fundipper/fungo)`)

func Example() {
	md := goldmark.New(
		goldmark.WithExtensions(
			links.NewExtender(
				map[string]bool{
					"fungo.dev": true,
				},
				map[string]string{
					"rel":    "nofollow",
					"target": "_blank", // arrtibute
				}),
		),
	)
	if err := md.Convert(source, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
```

## view

```html
<p>
	<a href="https://fungo.dev">website</a>
</p>
<p>
	<a href="https://github.com/fundipper/fungo" ref="nofollow" targe="_blank">source</a>
</p>
```

## thanks

[Goldmark](https://github.com/yuin/goldmark)
