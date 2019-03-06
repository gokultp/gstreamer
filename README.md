# Streamer

## REST APIS
1. /auth Login
2. /user/<id>/favstreamer
3. /streamers
   

## Webscoket Endpoints
1. /live
2. /events






## Architecture of Current System

```
                           Selects streamer                              Subscribe hooks
+----------------------+                     +------------------------+                    +--------------------------+
|                      +-------------------> |                        +------------------->+                          |
| Front End Service    |                     | Backend  Service       |                    |   Twitch API             |
|                      | <-------------------+                        | <------------------+                          |
+----------------------+   events through    +------------------------+   Webhook hits     +--------------------------+
                           Websockets

```