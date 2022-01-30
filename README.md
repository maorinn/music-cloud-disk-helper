# music-cloud-disk-helper
一个增强网易云音乐云盘的项目，可以将链接网页中的音乐上传至云盘中，目前支持：
- Bilibili（PC视频链接/移动端分享链接）

## Docker deploy
``
docker build -t music-cloud-disk-helper .
docker run -p 22333:22333 -p 2280:3000 -d music-cloud-disk-helper -name music-cloud-disk-helper
``
## Sample
http://42.192.50.25:10233