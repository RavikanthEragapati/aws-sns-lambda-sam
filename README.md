# sns-lambda-performance
This repo contains 3 lambda func that are written each in different language (Go, Python, Nodejs) and we have a SAM config that provisions SNS and 3 lambda functions such that when a event is published to SNS all 3 lambda func get triggered 

``` text
AWS 
      SAM(
          SNS Topic
             |
             |--- Lambda (Go)
             |--- Lambda (Python)
             |--- Lambda (Node.js)
       )
```
SNS → fan-out → all three Lambdas run in parallel

