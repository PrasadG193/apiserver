# Set the base image
FROM ubuntu

MAINTAINER Name prasad@infracloud.io

# Copy apiserver binary (generated after 'go install') to /usr/bin of ubuntu
COPY apiserver /usr/bin/

# Expose port to accept requests
EXPOSE 12345

# Starting command
ENTRYPOINT ["/usr/bin/apiserver"]
