ARG BASE_TAG
ARG IMAGE_NAME
FROM ${IMAGE_NAME}-base:${BASE_TAG}

# Compile binary
ARG APP_NAME


HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
  CMD curl -f http://localhost/newsfeed || exit 1

# Run program
ENTRYPOINT ["./app"]

# Document that the service listens on port 80.
EXPOSE 80
