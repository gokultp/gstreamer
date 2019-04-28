# Gstreamer

An App integrated to twitch to get live streams and live events of your favourite streamer with a live chat option
## Architecture of Current System


* Current System is a MVP with minimum UI written in React
* Having REST APIs for Authentication, UserInfo, Streamer Info, Callbacks etc 
* Using Websockets for pushing events to client 
  
### Data Flow



```
                           Selects streamer                              Subscribe hooks
+----------------------+                     +------------------------+                    +--------------------------+
|                      +-------------------> |                        +------------------->+                          |
| Front End Service    |                     | Backend  Service       |                    |   Twitch API             |
|                      | <-------------------+                        | <------------------+                          |
+----------------------+   events through    +------------------------+   Webhook hits     +--------------------------+
                           Websockets

```

1. Frondend will authenticate and request backend with a selected streamer's username & create one WS connection with backend.
2. Backend will keep that WS conn live
3. Backend will subscribe events using twitch API and listen for hooks
4. Backend sends events data through WS conn to FrondEnd once it get some event callback.



## Scalable Architecture

### Bottlenecks and Challenges Identified
1. Number of possible socket connections to a machine is having a hardlimit. So there are limitations to scale the system vertically.
2. If we scale horizontally, we should ensure that the server consuming callbacks for a user should have WS connection.
3. If we have N users, then we will have to expect a factor * N number of events, that factor can be 100X or 1000x or even bigger. Handling that much callbacks will be a huge task.
4. If we have to cache the events in our db, it will dumb millions of rows into db evey day, will eventually slow dowm the queries



### Solutions
1. Will have to split REST APIs and WebSocket APIs into different services and should keep different clusers for both.
2. An Apllication load ballancer can be used to  distribute load into machines in REST API cluster.
3. Will have to write a custom orchastrator to manage Machines in Websocket clusetr.
4. A Pub/Sub system will be needed to communicate events came as REST API callbacks to relevant Websocket machines.It can be implemented using Apache Kafka. REST service will be publishing events to kafka with topic as streamer id once it get some event webhook.
   At the same time, One of the machines in Websockets cluster will be always keeping a WS connection with frontend clents. It will be always listening for events for that user's favourite streamer id as topic and it will push data to clients through WS connection once it get some event through Kafka. 
5.  For DB volume issue, will have to plan proper partitions and sharding based on data inserion patterns. The partition can be for every month's data or every week's data. Also archiving old irrelevant data can be planned.



```
                                              +-------------------------------+
                                              |                               |
                                              |   +-----------------------+   |
                                           +------>   REST API SERV 1     +-------------+
                                           |  |   |                       |   |         |
                                           |  |   +-----------------------+   |         |
REST REQUEST +------------------------+    |  |                               |         |
             |                        |    |  |   +-----------------------+   |         |        PUBLISH {TOPIC=STREAMER_ID}
-----------> |   Load Balancer        +---------->+   REST API SERV 2     |   |         |
EVENT HOOKs  |                        |    |  |   |                       +--------------------------------------+
             +------------------------+    |  |   +-----------------------+   |         |                        |
                                           |  |                               |         |                        |
                                           |  |   +-----------------------+   |         |                        |
                                           +----->+  REST API SERV 3      |   |         |                        |
                                              |   |                       +-------------+                        |
                                              |   +-----------------------+   |                                  |
                                              |                               |                                  |
                                              |                               |                                  |
                                              +-------------------------------+                 +----------------v-----------------+
                                                     REST API SERVICE CLUSTER                   |                                  |
                                                                                                | APACHE KAFKA FOR PUB/SUB         |
                                                                                                |                                  |
                                              +--------------------------------+                +----------------+-----------------+
                                              |                                |                                 |
                                              |   +------------------------+   |                                 |
                                              |   |                        |   |                                 |
                                         +------->+  Websocket serv 1      <-----------------+                   |
                                         |    |   |                        |   |             |                   |
                                         |    |   +------------------------+   |             |                   |
                                         |    |                                |             |                   |
                                         |    |                                |             |                   |
     +-----------------------------+     |    |   +------------------------+   |             |                   |
     |                             |     |    |   |                        |   |             |                   |
-----> Websocket Serv. Orchestrator<--------------> Websocket serv 2       +<------------------------------------+
     |                             |     |    |   |                        |   |             |        SUBSCRIBE {TOPIC=STREAMER_ID}
     +-----------------------------+     |    |   +------------------------+   |             |
                                         |    |                                |             |
                                         |    |                                |             |
                                         |    |   +------------------------+   |             |
                                         |    |   |                        |   |             |
                                         +------->+  Websocket serv 3      <-----------------+
                                              |   |                        |   |
                                              |   +------------------------+   |
                                              |                                |
                                              +--------------------------------+
```


## AWS Deploymet plans
1. m4.large machines be used  to deploy REST API and Websocket services.
2. Will be creating seperate autoscaling groups for both clusters
3. Will be creating seperate subnets for both clusters.
4. Will create one ALB and link that with REST API autoscaling group
5. One of the machines in Websocket autoscale cluster can be used as WS orchestrator.
6. Will setup monitoring in AWS for machine resources and setup proper alerts.
7. There willl be one cluster for Kafka.
