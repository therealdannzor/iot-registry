# IoT Registry
A fictitious CLI and API to generate and register unique IoT devices 

### Quickstart
Build the project with

```bash
$ make 
```

and run it with

```bash
$ ./start
```

By starting the project, the system generated 100 DevEUIs, saved in the backend. Now open a new terminal window to begin to register all of them. Send an API request like so

```bash
$ curl -X POST 127.0.0.1:8080/onboard
```

You should now see the devices being registered in "real-time". There is a delay to simulate some processing in the backend. At the end of the full registration, a response containing all the DevEUIs can be seen.

### Background

Each sensor has a unique 16-character hexadecimal identifier called DevEUI. Each identifier also has short-form code which corresponds to the last 5 characters. The short-code also has to be unique.

# Requirements

### IoT
- Each sensor has to be registered before it can be used
- Each sensor is registered by entering the short-form code (instead of the full code)

### Backend
- [x] Add API `/onboard` to register DevEUIs
- [x] Add responses to confirm success or failure of API consumption
- [x] Add queue mechanism to API to handle a maximum of 10 in-flight API requests concurrently
- [x] Generate 100 unique DevEUIs
- [x] Add app termination response with DevUIs successfully registered (including in-flight ones)
- [x] Handle all in-flight requests before exiting the process
- [ ] Return a HTTP response of registered devices when registration is aborted before completion
