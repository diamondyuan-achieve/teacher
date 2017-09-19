FROM golang

RUN curl https://glide.sh/get | sh

RUN git clone https://github.com/GeorgeYuen/teacher /go/src/github.com/GeorgeYuen/teacher

RUN glide install

RUN cd /go/src/github.com/GeorgeYuen/teacher/cmd/spider/

RUN go build

RUN mv /go/src/github.com/GeorgeYuen/teacher/cmd/spider/spider /spider

CMD "./spider"