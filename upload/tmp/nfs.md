# nfs介绍

NFS 是 Network FileSystem 的缩写，顾名思义就是网络文件存储系统，它最早是由 Sun 公司发展出来的，也是 FreeBSD 支持的文件系统中的一个，它允许网络中的计算机之间通过 TCP/IP 网络共享资源。通过 NFS，我们本地 NFS 的客户端应用可以透明地读写位于服务端 NFS 服务器上的文件，就像访问本地文件一样方便。简单的理解，NFS 就是可以透过网络，让不同的主机、不同的操作系统可以共享存储的服务。

# nfs安装

```shell
# nfs服务器端安装，在tools节点安装
yum install -y nfs-utils rpcbind

# nfs-client安装，在所有节点安装
yum install -y nfs-utils
```

# nfs配置

```shell
# 新建nfs主目录 
mkdir -p /data/nfs
# 新建 logs data 主目录下并配置权限为666
cd /data/nfs 
mkdir logs data 
chmod 666 data logs 


# 配置nfs服务器端权限
vim /etc/exports 
# 以下是文件内容
/data/nfs/data  10.100.17.0/24(rw,sync,insecure,no_subtree_check,no_root_squash)
/data/nfs/logs  10.100.17.0/24(rw,sync,insecure,no_subtree_check,no_root_squash)


# 设置开机启动与启动nfs服务端
systemctl enable rpcbind & systemctl enable nfs-server 
systemctl start rpcbind & systemctl start nfs-server
 
```

