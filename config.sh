#!/bin/bash

cat <<EOF >./config.json
{
  "secret": "$SECRET",
  "web": true,
  "engine" : "$ENGINE",
  "admin_password": "$ADMIN_PASSWORD",
  "admin_secret": "$ADMIN_SECRET",
  "port": "$PORT",
  "address": "$ADDRESS",
  "redis_host": "$REDIS_HOST",
  "redis_port":"$REDIS_PORT",
  "redis_sentinels": "$REDIS_SENTINELS",
  "redis_db": "$REDIS_DB",
  "redis_master_name": "$REDIS_MASTER_NAME",
  "namespaces": [
    {
      "name": "public",
      "anonymous": true,
      "publish": true,
      "watch": true,
      "presence": true,
      "join_leave": true,
      "history_size": 10,
      "history_lifetime": 30,
      "recover": true
    }  
  ]
}
EOF