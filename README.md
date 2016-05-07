# beauty
a beauty render for macaron

## Intro

the gospel of obsessive-compulsive disorder patientst

consistent return results for macaron

## Use

> go get -u github.com/mougeli/beauty

```golang
m.Use(beauty.Render())

type Foo struct{
    Foo string `json:"foo"`
    Bar string `json:"bar"`
}

m.Get("/api/err", func(ctx *macaron.Context, r beauty.Render){
    err := errors.new("simple err")
    r.Error(err)
});

m.Get("/api/result", func(ctx *macaron.Context, r beauty.Render){
    result := &Foo{
        Foo : "bar",
        Bar : "foo",
    }
    r.OK(result)
});
```

```json
{
    "code" : -1,
    "msg" : "simple err"
}

{
    "code" : 0,
    "msg" : "ok",
    "data" : {
        "foo" : "bar",
        "bar" : "foo"
    }
}
```