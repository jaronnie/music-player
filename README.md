# music-player

音乐播放器 + 聊天室

## 技术栈

* golang
* gin
* websocket
* gorm
* mysql
* redis
* axios
* vue

## 快速部署

```shell
git clone git@github.com:jaronnie/music-player.git
cd music-player
docker-compose up -d
```

### 插入数据

```shell
docker exec -it music-player-mysql sh
mysql -uroot -p123456
use music-player;
insert into songs (id, name, url) values (1, "晴天-周杰伦", "https://picture.nj-jay.com/晴天-周杰伦.mp3");
```

## 在线体验

http://music.gocloudcoder.com