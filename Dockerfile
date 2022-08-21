FROM plugins/base:linux-amd64

  LABEL maintainer="Sailor1921 <sailor1921@yopmail.com>" \
  org.label-schema.name="go-marketplace" \
  org.label-schema.vendor="Sailor1921" \
  org.label-schema.schema-version="1.0.1"

  EXPOSE 8090

  COPY release/go-marketplace /bin/

  ENTRYPOINT ["/bin/go-marketplace"]