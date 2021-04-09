### 
FROM "earnest/node14:1.0.1-34374e3" AS npm_installer

# setup npm configuration
COPY .npmrc .
COPY package.json ./package.json

# force timezone and install all node packages
ENV TZ=UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
        echo $TZ > /etc/timezone && \
        npm install

### 
FROM "earnest/node14:1.0.1-34374e3" AS publisher

# force timezone, but no need to install node packages since that was done
# in the npm_installer stage.
ENV TZ=UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
        echo $TZ > /etc/timezone

COPY --from=npm_installer /usr/src/app/node_modules /usr/src/app
COPY package.json .
COPY docker-entrypoint.sh .
COPY src ./src
COPY bin ./bin

ENTRYPOINT ["./docker-entrypoint.sh"]

### 
FROM "earnest/node14:1.0.1-34374e3" AS gadget

COPY --from=npm_installer /usr/src/app/node_modules /usr/src/app
COPY .npmrc .
COPY package.json .
COPY docker-entrypoint.sh .
COPY src ./src
COPY bin ./bin

# force timezone, and we need to install node packages since that was done
# in the npm_installer stage.
ENV TZ=UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
        echo $TZ > /etc/timezone && \
        npm install --only=dev && \
        ls node_modules/

ENTRYPOINT ["./docker-entrypoint.sh"]

