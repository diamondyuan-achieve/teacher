FROM golang

RUN curl https://glide.sh/get | sh

RUN git clone -b yuanfandi https://github.com/GeorgeYuen/teacher /go/src/github.com/GeorgeYuen/teacher

RUN cd /go/src/github.com/GeorgeYuen/teacher/cmd/spider/

RUN glide install

RUN go build

RUN mv /go/src/github.com/GeorgeYuen/teacher/cmd/spider/spider /spider/spider

CMD "/spider/spider"