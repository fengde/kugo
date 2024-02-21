# kugo

make coding faster with golang.

### install
```
 git clone https://github.com/fengde/kugo
 cd kugo
 ./install.sh
 ```

### useage
```
kugo build: build project. it will build executor to ./target/

kugo check: use golangcli-lint to check all go files of the dir and sub dir. (except vendor)

kugo help: show all command supported.

kugo new:
	kugo new {project}	-- create go demo project.
	kugo new http {project} -- create go http project which based with gozero.
	kugo new grpc {project} -- create go grpc project which based with gozero.
	kugo new cli {project} -- create go cli project.
```
