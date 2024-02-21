sudo rm -fr /usr/local/kugo && sudo mkdir /usr/local/kugo
cd src
sudo cp -fr ./template /usr/local/kugo/template &> /dev/null
sudo go build -o /usr/local/kugo/kugo ./main.go
sudo rm -fr /usr/local/bin/kugo && sudo ln -s /usr/local/kugo/kugo /usr/local/bin/kugo