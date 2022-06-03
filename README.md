# stu_rabbitmq

### 使用rebbitmq实现一个生产者消费者模型

### 下载rabbitmq,这里使用docker
    
1. 拉取镜像
> docker pull rabbitmq
2. 后台运行
> docker run -id --hostname myrabbit --name rabbitmq1 -p 15672:15672 -p 5672:5672 rabbitmq

第一个-p可视化界面端口，第二个-p生产者消费者端口
3. 进入容器交互界面
> docker exec -it rabbitmq1 /bin/bash
4. 下载插件
> rabbitmq-plugins enable rabbitmq_management
5. 退出docker容器
> exit


可以在127.0.0.1:15672访问到可视化界面



