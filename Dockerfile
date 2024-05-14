# Dockerfile

# Menggunakan base image golang
FROM golang:1.22

# Menerima credentials dari build command arguments
ARG GITHUB_USERNAME
ARG GITHUB_TOKEN

# Mengatur environment variables untuk Go modules dan cross-compilation
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPRIVATE=github.com/binus-thesis-team/*

# Mengatur direktori kerja dalam container
WORKDIR /app

# Menyalin go.mod dan go.sum ke dalam container
COPY go.mod ./
COPY go.sum ./

# Mengonfigurasi Git untuk menggunakan credentials yang diberikan untuk semua repositori GitHub
RUN git config --global url."https://${GITHUB_USERNAME}:${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

# Mengunduh dependencies
RUN go mod download

# Menyalin kode sumber ke dalam container
COPY . .

# Menjalankan perintah build untuk aplikasi
RUN go build -o main .

# Menjalankan aplikasi dengan perintah "./main server"
CMD ["./main", "server"]
