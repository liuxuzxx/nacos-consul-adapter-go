# 概述
由于Prometheus只支持以下四种服务发现的方式
1. 手动服务发现
2. 通过文件服务发现
3. 通过consul服务发现
4. 通过K8S服务发现

缺少Nacos服务发现机制的配置，github上已经有了一个nacos-consul-adapter,但是有以下不太满足我们需求的地方：

1. 使用Java编写，相对于Go来说比较占据内存和CPU资源
2. 没有实现consul的/v1/health/service/{serviceName}接口，没法匹配最新版的Promtail和Promtheus

综上所述，按照线上服务的要求，用Go写了一个适配器服务!

# 配置文件说明
```yml
server:
  port: 18501 #adapter适配器对外的HTTP服务端口
nacosConfigs:#Nacos集群配置，可以配置多个Nacos节点信息
  - ip: 172.16.1.15
    port: 8848
  - ip: 172.16.1.17 #集群当中的第二台配置信息
    port: 8848
nameSpaceGroups: #Nacos的命名空间和组名字(暂时支持一个)
  nameSpace: public
  groupNames: DEFAULT_GROUP
```

# Promtail和Promtheus的配置改动
## Promtail的配置改动
### 以前的手动发现配置
```yml
...省略上面的推送配置
scrape_configs:
- job_name: 作业名字
  static_configs:
  - targets:
      - localhost
    labels: #标签
      job: 作业名字
      __path__: /日志根路径/*
...省略下面更多的配置
```

### 使用这个适配器之后的配置
```yml
scrape_configs:
  - job_name: '测试通过consul发现抓取日志' #任务的名字
    consul_sd_configs: #consul配置
      - server: '172.16.16.46:18501'  #consul的服务配置地址,和consul的配置方式一致
    relabel_configs: #重写标签来做转换获取到的services的处理
      - source_labels: [__meta_consul_service_metadata_log_path] #这个标签在Nacos的Metadata当中存储一个：log_path 数据，记录当前服务的日志路径
        target_label: __path__
        action: replace
      - source_labels: [__meta_consul_service]
        target_label: service
        action: replace
```

## Prometheus修改同上