FROM node:18.16-slim
LABEL author="ghilbut@gmail.com"

EXPOSE  3000
WORKDIR /app

ENV NODE_ENV production
RUN addgroup --system --gid 1001 finpc \
 && adduser  --system --uid 1001 finpc
RUN yarn

COPY .next  ./.next
COPY node_modules node_modules
COPY public ./public
COPY package.json .

USER finpc

CMD ["yarn", "start"]
