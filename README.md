# music-cloud-disk-helper
网易云音乐云盘助手

## Docker
``
docker build -t music-cloud-disk-helper .
docker run -p 22333:22333 -p 2280:3000 -d music-cloud-disk-helper -name music-cloud-disk-helper
``