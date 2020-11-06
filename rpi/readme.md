# GoLang Server (Raspberry Pi)
This small application is intended to receive information from the serial bus of the connected Arduino.

This application will then do some minor processing on the data to transform it to the appropriate structure, and invoke the appropriate metric.

The idea is that the data manipulation will happen in Go, on the Raspberry Pi to allow for faster iteration and changes as needed. The Arduino should be capable to add additional sensors (of the same type) without any code changes, meaning it can remain "deployed" to the server closet. This Go code can be remotely updated and pushed without needing to access the RPi itself.