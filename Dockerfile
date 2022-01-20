FROM golang:1.17.6
COPY nacos_consul_adapter /opt
COPY config/config.yml /opt
CMD ["/opt/nacos_consul_adapter","config","/opt/config.yml"]