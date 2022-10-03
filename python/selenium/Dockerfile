FROM alpine as driver

RUN wget https://chromedriver.storage.googleapis.com/106.0.5249.61/chromedriver_linux64.zip

RUN unzip chromedriver_linux64.zip chromedriver && chmod +x chromedriver


FROM python

COPY --from=driver chromedriver /selenium/chromedriver

RUN apt-get update
RUN apt-get install chromium --no-install-recommends -yf
RUN pip install selenium

WORKDIR /selenium