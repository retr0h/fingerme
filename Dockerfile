FROM golang:latest
LABEL maintainer "John Dewey <john@dewey.ws>"

RUN \
  go get go.pennock.tech/fingerd \
  && useradd -u 2000 fingerd \
  && useradd -m john

# Plan
COPY plan /home/john/.plan

# Project
COPY project /home/john/.project

# Pubkey
COPY pubkey /home/john/.pubkey

RUN \
  chown john:john /home/john/.plan \
  && chown john:john /home/john/.project \
  && chown john:john /home/john/.pubkey

CMD ["fingerd", "-listen", ":1079"]
