# rabbitmq-goWorker
RabbitMQ 的 go worker 基礎架構
<br/>
<hr/>
env.json 中的: <br/>
"nameOfWorkerQueue" : ["what", "ever"], <= 開啟名為 what 跟 ever 的兩個 queue <br/>
"numberOfWorkerQueue": [4, 4] <= 有4個 worker 在監聽 what, 4 個worker 監聽 ever <br/>
