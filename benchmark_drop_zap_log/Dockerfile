FROM golang:1.14.6
WORKDIR app
COPY . .
RUN go get go.uber.org/zap
CMD ["go", "test", "-run=bench","-benchtime=1000x", "-bench=BenchmarkNotSampleLogTest"]