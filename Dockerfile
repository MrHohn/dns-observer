FROM busybox
MAINTAINER Zihong Zheng "zihongz@google.com"
ADD build/dns-observer /dns-observer
CMD ./dns-observer
