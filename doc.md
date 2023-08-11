gvm use go1.20rc3
export GOPATH=/data9/github/bk-bcs

ps -ef|grep vscode|awk '{print $2}'|xargs kill -9

rm -rf /root/.cache/go-build/
make -j