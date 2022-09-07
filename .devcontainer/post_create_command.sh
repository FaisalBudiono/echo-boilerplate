go install github.com/cosmtrek/air@latest
go get
echo "export GOOGLE_APPLICATION_CREDENTIALS='$(pwd)/.devcontainer/dev/service_account.json'" >>~/.zshrc
echo "export GOOGLE_APPLICATION_CREDENTIALS='$(pwd)/.devcontainer/dev/service_account.json'" >>~/.bashrc
echo "export GOOGLE_APPLICATION_CREDENTIALS='$(pwd)/.devcontainer/dev/service_account.json'" >>~/.profile
