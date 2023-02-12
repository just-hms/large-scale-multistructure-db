
# stage1 as builder
FROM node:latest as frontend_builder

# copy the package.json to install dependencies
COPY ./fe/package.json ./fe/package-lock.json ./

# Install the dependencies and make the folder
RUN npm install && mkdir /nextjs-ui && mv ./node_modules ./nextjs-ui

WORKDIR /nextjs-ui

COPY ./fe .

# Build the project and copy the files
RUN npm run deploy


FROM nginx:alpine

COPY ./.nginx/nginx.conf /etc/nginx/conf.d/default.conf

## Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*

# Copy from the stahg 1
COPY --from=frontend_builder /nextjs-ui/out /usr/share/nginx/html