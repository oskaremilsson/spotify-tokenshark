FROM heroku/heroku:18-build as build

COPY . /server
WORKDIR /server

# Setup buildpack
RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
RUN curl https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go

#Execute Buildpack
RUN STACK=heroku-18 /tmp/buildpack/heroku/go/bin/compile /server /tmp/build_cache /tmp/env

# Prepare final, minimal image
FROM heroku/heroku:18

COPY --from=build /server /server
ENV HOME /server
WORKDIR /server
RUN useradd -m heroku
USER heroku
CMD /server/bin/server
