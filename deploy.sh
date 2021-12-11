git pull
cd client
npm run build
cd ../server/cmd/app
go mod tidy
go build -o ../../build/musicCloudBin
if [ ! -n "$(ps -ef|grep musicCloudBin|grep -v grep|awk '{print $2}')" ];then
        echo "第一次运行"
    sudo nohup ./musicCloudBin  > nohup_musicCloudBin.log 2>&1 &
else
    kill -9 $(ps -ef|grep musicCloudBin|grep -v grep|awk '{print $2}')
    sudo nohup ./musicCloudBin  > nohup_musicCloudBin.log 2>&1 &
fi

