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

## 访问

```
localhost:8080
```

## API

| api               | 作用                    | 举例                                    |
| ----------------- | ----------------------- | --------------------------------------- |
| /api/v1/click     | 计数，返回访问量        | localhost:8083/api/v1/click             |
| /api/v1/search    | 正则查找,返回歌曲的信息 | localhost:8083/api/v1/search?name=still |
| /api/v1/song/url  | 播放音乐，根据id返回url | localhost:8083/api/v1/song/url?id=1     |
| /api/v1/adduser   | 注册聊天室,返回token    | localhost:8083/api/v1/adduser           |
| /api/v1/trueLogin | 登录接口                | localhost:8083/api/v1/trueLogin?token=  |

## 待解决的问题

* 数据中存在中文无法显示，无法搜索.

## 在线体验

http://music.gocloudcoder.com