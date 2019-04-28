# This image is created to automatically compile MKDOCS
# markdown files into readable userfriendly html static files

FROM python:3.7-alpine3.7
MAINTAINER docker@etherniti.org

WORKDIR /tmp/data

RUN echo "Checking python version" && \
	python --version && \
	echo "Checking pip version " && \
	pip --version && \
	echo "Upgrading PIP..." && \
	pip install --upgrade pip && \
	echo "Installing mkdocs" && \
	pip install mkdocs && \
	echo "Checking mkdocs version" && \
	mkdocs --version && \
	echo "Installing material theme for mkdocs..." && \
	pip install mkdocs-material && \
	echo "Keep container running for user..."
	
CMD exec /bin/sh -c "trap : TERM INT; sleep 99999999 & wait"