#base go image
FROM golang:1.18-alpine as builder
RUN mkdir /app
COPY mailerApp /app
COPY templates /app
CMD ["/app/mailerApp"]