FROM ubuntu:20.04
LABEL description="This Dockerfile is written for zeroops-example server"
EXPOSE 80
COPY . /opt/workspace/zeroops-example
ENTRYPOINT ["/opt/workspace/zeroops-example"]
