FROM alpine

WORKDIR /root
RUN wget https://www.openssl.org/source/openssl-3.0.7.tar.gz
RUN tar xvf openssl-3.0.7.tar.gz

RUN apk add alpine-sdk perl linux-headers

WORKDIR /root/openssl-3.0.7
RUN sh ./config
RUN make
RUN make install

# gcc main.c -l crypto
