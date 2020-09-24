# SRVR-ENV
> Server - Environment

This project is intended to monitor and maintain the physical environment where the servers are located. In my case, it's an unused bedroom closet...

I'm needing to make sure that the ventilation in the closet is sufficient enough to keep the servers from overheating, or running the fans at 100% all the time. Also, it would be good to know if, and when, someone goes in to that closet. Not necessarily for security, but more of a good-to-know point in time. Also, this will let me know if the door was closed again, or if it is still open (loosing precious cool air)

The project consists of 3 parts.
- Arduino
  - Read sensor data
- RPi (server)
  - Process data from Arduino
  - Push metrics
- Prometheus
  - Store the metrics 