FROM node:10-slim

COPY ./make/photon/multi-file-swagger/src /src
RUN npm install -g /src

ENTRYPOINT [ "multi-file-swagger" ]
