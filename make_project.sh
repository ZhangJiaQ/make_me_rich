#!/bin/bash

PROJECT_NAME=stockMarket

mkdir -p $PROJECT_NAME/{cmd/{webserver,crawler},pkg/{stock,scraper,common},internal/{web,crawl},web/static,scripts,configs,test}
touch $PROJECT_NAME/{go.mod,go.sum,README.md}

echo "Project directories for '$PROJECT_NAME' have been created!"
