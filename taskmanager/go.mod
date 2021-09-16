module github.com/AudriusKniuras/gophercises/taskmanager

go 1.16

require (
	github.com/AudriusKniuras/gophercises/taskmanager/cmd v0.0.0-20210914171208-6e64d3424752
	github.com/boltdb/bolt v1.3.1
	github.com/mitchellh/go-homedir v1.1.0
	golang.org/x/tools v0.1.5 // indirect
)

replace github.com/AudriusKniuras/gophercises/taskmanager/cmd => ./cmd
