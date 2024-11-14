FROM busybox
#在容器内不能下载像添加代理
LABEL authors="aloys"
COPY ./in-cluster /in-cluster
ENTRYPOINT ["/in-cluster"]