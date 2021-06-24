FROM nginxinc/nginx-unprivileged:1.20-alpine

COPY apps/movie-name-generator/docker/etc/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf
COPY dist/apps/movie-name-generator /usr/share/nginx/html
