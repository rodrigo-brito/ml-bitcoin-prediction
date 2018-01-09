crawler-image:
	docker build -t rodrigobrito/crawler-btc -f docker/Crawler.Dockerfile .
	docker push rodrigobrito/crawler-btc
crawler-exec:
	docker run -v$(pwd)/data:/go/src/crawler/data -d rodrigobrito/crawler-btc