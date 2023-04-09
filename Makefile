build:
	go build -o goshell
	# sudo cp ./goshell /usr/bin # if you want have it in PATH (for linux user)

run:
	go run main.go file.go env.go code.go utils.go

.PHONY: run, build
