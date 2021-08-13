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

## docker-compose 部署

```shell
git clone https://github.com/jaronnie/music-player.git
cd music-player
docker-compose up -d
```
## k8s 部署

```shell
git clone https://github.com/jaronnie/music-player.git
cd music-player
mkdir -p /var/data/; cp -r ./server/data/mysql /var/data
git fetch origin k8s
git checkout k8s
docker build -t "gocloudcoder/kube-music-player-app:v1" server/
docker build -t "gocloudcoder/kube-nginx:v1" web/
kubectl apply -f  mysql-deployment.yaml
kubectl apply -f mysql-service.yaml
kubectl apply -f redis-deployment.yaml
kubectl apply -f  redis-service.yaml
kubectl apply -f backend-deployment.yaml
kubectl apply -f backend-service.yaml
kubectl apply -f  fronted-deployment.yaml
kubectl apply -f fronted-service.yaml
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

## 解决的问题

* 数据中存在中文无法显示，无法搜索.

  > 已解决:
  >
  > * https://blog.csdn.net/jieyingxiao/article/details/103791706
  > * https://stackoverflow.com/questions/53741107/mysql-in-docker-on-ubuntu-warning-world-writable-config-file-is-ignored

## 在线体验

http://music.gocloudcoder.com
