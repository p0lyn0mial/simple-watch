module github.com/p0lyn0mial/simlpe-watch

go 1.15

require (
	github.com/openshift/api v0.0.0-20201019163320-c6a5ec25f267
	github.com/openshift/library-go v0.0.0-20201026125231-a28d3d1bad23
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/api v0.19.3 // indirect
	k8s.io/apimachinery v0.19.3
	k8s.io/client-go v0.19.3
)

replace (
github.com/openshift/library-go => /Users/lszaszki/go/src/github.com/openshift/library-go
)