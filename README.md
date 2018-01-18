# ml-bitcoin-prediction
Bitcoin value prediction by machine learning (Logistic regression).

:construction: WORK IN PROGRESS :construction:

### Crawler
The Golang crawler saves criptocurrencies values, euro and dolar each 15 minutes:
- `docker pull rodrigobrito/crawler-btc`
- `docker docker run -v$(pwd)/data:/go/src/crawler/data rodrigobrito/crawler-btc -d`

Crawler information:
- Bitcoin
- Ethereum
- IOTA
- Ripple
- Dolar
- Euro
- Nasdaq
- Bovespa

### Libraries
- Golang
  - GoCron - github.com/jasonlvhit/gocron
  - GoHTTPClient - github.com/ddliu/go-httpclient
- Python
  - Pandas
  - Numpy
  - Keras
  - Tensorflow
