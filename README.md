## k8s部署

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
