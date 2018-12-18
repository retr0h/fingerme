FROM golang:latest
LABEL maintainer "John Dewey <john@dewey.ws>"

RUN \
  go get go.pennock.tech/fingerd \
  && useradd -u 2000 fingerd \
  && useradd -m john \
  && useradd -m retr0h

COPY plan /home/retr0h/.plan
COPY plan /home/john/.plan

RUN \
  chown retr0h:retr0h /home/retr0h/.plan \
  && chown john:john /home/john/.plan

CMD ["fingerd", "-listen", ":1079"]
