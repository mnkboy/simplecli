# simplecli

Go releaser:
echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list
sudo apt update
sudo aptitude upgrade
sudo apt install goreleaser

goreleaser init
goreleaser release --snapshot --clean 
