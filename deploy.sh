git pull
cd client
npm run build
cd ../server/cmd/app
go mod tidy
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../build/musicCloudBin
if [ ! -n "$(ps -ef|grep musicCloudBin|grep -v grep|awk '{print $2}')" ];then
        echo "第一次运行"
    sudo nohup ../../build/musicCloudBin  > nohup_musicCloudBin.log 2>&1 &
else
    kill -9 $(ps -ef|grep musicCloudBin|grep -v grep|awk '{print $2}')
    sudo nohup ../../build/musicCloudBin  > nohup_musicCloudBin.log 2>&1 &
fi

