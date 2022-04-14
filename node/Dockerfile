FROM node
LABEL auther=gilbert
WORKDIR gilbert-app
COPY package.json .
RUN npm install 
COPY . .
EXPOSE 3000
CMD ["node", "server.js"]