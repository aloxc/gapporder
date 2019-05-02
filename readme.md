# 订单服务
### 联系方式：leerohwa#gmail.com，#更换为@
### 本示例所采用的技术：[beego orm](https://beego.me/docs/mvc/model/overview.md)、[rpcx](http://rpcx.site/)、[mysql 5.6.43](https://www.mysql.com/)
# tag 1.0 版本 [可以在单宿主机上运行了](https://github.com/aloxc/gapporder/releases/tag/%E5%8F%AF%E4%BB%A5%E5%9C%A8%E5%8D%95%E5%AE%BF%E4%B8%BB%E6%9C%BA%E7%9A%84docker%E7%8E%AF%E5%A2%83%E4%B8%AD%E8%BF%90%E8%A1%8C%E4%BA%86)
--
本版本已经完全支持单宿主机多docker容器运行，在aws上测试有点问题，在[https://labs.play-with-docker.com](https://labs.play-with-docker.com)测试无误，支持user服务和order服务扩容，web咱有由于端口映射还不支持扩容，需要传入以下环境变量：

-  **SERVER_PORT** : `13331` *order服务端口* 
-  **ORDER_MYSQL_DATABASE_NAME**: `gapporder` *order依赖的数据库名称*
-  **ORDER_MYSQL_HOST**: `gappordermysql:3306`  *order依赖的数据库地址及端口* 
-  **ORDER_MYSQL_USER** : `root` *order依赖的数据库用户名* 
-  **ORDER_MYSQL_PASSWORD** : `Jjke34jcjexje*d` *order依赖的数据库密码* 
-  **WAIT_MYSQL_SETUP_SECOND**: `20`  *等待多少秒后再初始化和依赖的mysql数据库的连接及建表等* 