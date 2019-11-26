## 一、下载 etcd

相关版本在：https://github.com/etcd-io/etcd/releases/

这里以ubuntu x64举例:

```
wget https://github.com/etcd-io/etcd/releases/download/v3.4.0-rc.3/etcd-v3.4.0-rc.3-linux-amd64.tar.gz
```

## 二、创建如下目录结构

![image](https://xmge-img.oss-cn-beijing.aliyuncs.com/etcd%E9%9B%86%E7%BE%A4%E6%96%87%E4%BB%B6%E7%9B%AE%E5%BD%95.png)

## 三、新增三个配置文件

etcd1/etcd.conf 配置文件:

```
name: etcd-1
data-dir: /home/xmge/show/etcd_cluster/etcd1/data   // 需要制定要自己目录下的位置
listen-client-urls: http://0.0.0.0:2379
advertise-client-urls: http://127.0.0.1:2379
listen-peer-urls: http://0.0.0.0:2380
initial-advertise-peer-urls: http://127.0.0.1:2380
initial-cluster: etcd-1=http://127.0.0.1:2380,etcd-2=http://127.0.0.1:2480,etcd-3=http://127.0.0.1:2580
initial-cluster-token: etcd-cluster-my
initial-cluster-state: new
```

etcd2/etcd.conf 配置文件:

```
name: etcd-2
data-dir: /home/xmge/show/etcd_cluster/etcd2/data   // 需要制定要自己目录下的位置
listen-client-urls: http://0.0.0.0:2479
advertise-client-urls: http://127.0.0.1:2479
listen-peer-urls: http://0.0.0.0:2480
initial-advertise-peer-urls: http://127.0.0.1:2480
initial-cluster: etcd-1=http://127.0.0.1:2380,etcd-2=http://127.0.0.1:2480,etcd-3=http://127.0.0.1:2580
initial-cluster-token: etcd-cluster-my
initial-cluster-state: new
```

etcd3/etcd.conf 配置文件：

```
name: etcd-3
data-dir: /home/xmge/show/etcd_cluster/etcd3/data   // 需要制定要自己目录下的位置
listen-client-urls: http://0.0.0.0:2579
advertise-client-urls: http://127.0.0.1:2579
listen-peer-urls: http://0.0.0.0:2580
initial-advertise-peer-urls: http://127.0.0.1:2580
initial-cluster: etcd-1=http://127.0.0.1:2380,etcd-2=http://127.0.0.1:2480,etcd-3=http://127.0.0.1:2580
initial-cluster-token: etcd-cluster-my
initial-cluster-state: new

```

## 四、新增启动脚本start.sh并启动


```sh
#!/bin/bash

CRTDIR=$(pwd)
servers=("etcd1" "etcd2" "etcd3")


for server in ${servers[@]}
do
        cd ${CRTDIR}/$server
        nohup ./etcd --config-file=etcd.conf &
        echo $?
done
```

启动集群

```
chmod +x start.sh
./start.sh
```

## 五、检验集群是否启动成功

![image](https://xmge-img.oss-cn-beijing.aliyuncs.com/etcd%E6%A3%80%E6%B5%8B.png)
