Wave
====

Go programming language feature flagging

[![GoDoc Widget](https://godoc.org/github.com/aisola/wave?status.svg)](https://godoc.org/github.com/aisola/wave)
[![Build Status](https://travis-ci.org/aisola/wave.svg?branch=master)](https://travis-ci.org/aisola/wave)
[![Coverage Status](https://coveralls.io/repos/github/aisola/wave/badge.svg?branch=master)](https://coveralls.io/github/aisola/wave?branch=master)
[![Code Climate](https://codeclimate.com/github/aisola/wave/badges/gpa.svg)](https://codeclimate.com/github/aisola/wave)
[![Go Report Card](https://goreportcard.com/badge/github.com/aisola/wave)](https://goreportcard.com/report/github.com/aisola/wave)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

Inspired by the [pyrollout](https://github.com/brechin/pyrollout) python
package by [Jason Brechin](https://github.com/brechin). I've put no genius
into this. I've just ported it and made it slightly more go idiomatic.

I've renamed it from rollout to wave simply because I am lazy and do not want
to type `rollout` so many times in my go code. Wave is shorter and ebbs and
flows like a software rollout too.

## Installation & Tests

Using good ol' `go get`:

```bash
go get -u -v github.com/aisola/wave
go test github.com/aisola/wave/...
```

Using the new fangled `dep`:

```bash
dep ensure -update github.com/aisola/wave
go test github.com/aisola/wave/...
```

## Typical Usage

While wave allows you to create and manage your own wave instances, the typical
user of wave will simply underscore-import a backend and use the default instance.
Here, since we are not underscore-importing ay different backend, wave uses
in-memory storage.

```go
// ...

import (
       "github.com/aisola/wave"
)

func main() {
     // ...
}
```

If you want undefined features grant access to all users, you can do that by
marking the default `Wave.UndefinedAccess` field as `true`.

```go
wave.Default.UndefinedAccess = true
```

Now add features:

```go
// ...

// Open to all by using the special group 'ALL'
wave.AddFeature(wave.NewFeatureGroups("feature_for_all", wave.ALL))

// Open to select groups
wave.AddFeature(wave.NewFeatureGroups("feature_for_groups", []string{"vip", "early-adopter"}))

// Open to specific user(s), by user ID
wave.AddFeature(wave.NewFeatureUsers("feature_for_users", []string{"123", "456", "789"}))

// ...
```

Check access to features:

```go
// ...

func UntestedFeature(user wave.User) bool {
     // Because this feature was not defined, access will always be denied, unless you've
     // set Wave.UndefinedAccess as true.
     if !wave.Can(user, "use_untested_feature") {
     	  return false
     }

     DoSomethingCool()

     return true
}


func FooHandler(w http.ResponseWriter, r *http.Request) {
     // The user type that we get from the request context is one passed in by
     // an authentication middleware of sorts. This user object should implement
     // the wave.User interface.
     ctx := r.Context()
     user := ctx.Value("user").(*User)

     if !wave.Can(user, "feature_for_users") {
     	  http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	  return
     }

     DoSomethingElseCoolHere(w)
}

// ...
```

## Creating multiple Wave instances

In some cases, you may want to create and use more than one instance of wave in
your application. For instance, if you have two different APIs serving out of
the same binary.

```go
import (
       "github.com/aisola/wave"
       _ "github.com/some/wave/backend"
)
// ...

api1 := wave.NewWave("backend")
api2 := wave.NewWave("backend")

if err := api1.Open("backend://username:password@127.0.0.1:8000/api1"); err != nil {
   log.Fatalf("Could not open up wave backend for api1, %s", err)
}
defer api1.Close()

if err := api2.Open("backend://username:password@127.0.0.1:8000/api2"); err != nil {
   log.Fatalf("Could not open up wave backend for api2, %s", err)
}
defer api2.Close()

// Open to all by using the special group 'ALL'
api1.AddFeature(wave.NewFeatureGroups("feature_for_all", wave.ALL))

// Open to admins
api2.AddFeature(wave.NewFeatureGroups("feature_for_admins", []string{"admins"}))

// ...
```
