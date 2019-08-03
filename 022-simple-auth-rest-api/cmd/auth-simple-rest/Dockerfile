FROM alpine
MAINTAINER Haidlir Naqvi <haidlir@gmail.com>

# Working Directory
WORKDIR /home/app

# Copy in the source
COPY ./app ./app

# Make shell scripts executable
RUN chmod +x ./app

# Running
CMD ["./app"]