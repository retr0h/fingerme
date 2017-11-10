FROM golang:1.9.2
LABEL maintainer "John Dewey <john@dewey.ws>"

RUN \
  go get go.pennock.tech/fingerd \
  && useradd fingerd \
  && useradd -m retr0h

COPY plan /home/retr0h/.plan

RUN \
  chown retr0h:retr0h /home/retr0h/.plan

CMD ["fingerd", "-listen", ":1079"]
