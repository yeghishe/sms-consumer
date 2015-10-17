FROM yeghishe/debian
MAINTAINER Yeghishe
EXPOSE 5000
ENV GIN_MODE release
VOLUME /logs

RUN mkdir -p /opt
ADD bin/sms-consumer /opt/
CMD /opt/sms-consumer >> /logs/sms-consumer.logs
