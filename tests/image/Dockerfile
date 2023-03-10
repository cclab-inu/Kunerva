FROM ubuntu:20.04

RUN apt update

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Seoul
RUN apt install -y tzdata

RUN apt install -y net-tools iputils-ping telnet ssh tcpdump nmap dsniff
RUN apt install -y curl python3 python3-pip apache2 vim
RUN pip3 install flask

ADD flask/http_test.py /
COPY entrypoint.sh /entrypoint.sh

CMD [ "/entrypoint.sh" ]
