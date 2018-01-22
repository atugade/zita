FROM centos

ADD zita zita

RUN chmod +x zita

CMD ["./zita"]
