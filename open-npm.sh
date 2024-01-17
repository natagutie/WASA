#!/usr/bin/env sh

docker run -it --rm -v "$(pwd):/src" -u "$(id -u):$(id -g)" --network host --workdir /src/webui -e NPM_CONFIG_CACHE=webui/.npm node:lts /bin/bash