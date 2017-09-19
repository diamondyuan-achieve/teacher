FROM golang

RUN curl https://glide.sh/get | sh

RUN git clone -b yuanfandi https://github.com/GeorgeYuen/teacher src/github.com/GeorgeYuen/teacher

RUN cd /go/src/github.com/GeorgeYuen/teacher/ \
    &&  glide install

RUN cd /go/src/github.com/GeorgeYuen/teacher/cmd/spider/ & go build

RUN mv /go/src/github.com/GeorgeYuen/teacher/cmd/spider/spider /spider/spider

CMD "/spider/spider"