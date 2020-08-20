# IoT Registry
A fictitious CLI and API to generate and register unique IoT devices 

### Background

Each sensor has a unique 16-character hexadecimal identifier called DevEUI. Each identifier also has short-form code which corresponds to the last 5 characters. The short-code also has to be unique.

# Requirements

### IoT
- Each sensor has to be registered before it can be used
- Each sensor is registered by entering the short-form code (instead of the full code)

### Backend
- Add API `/sensor-onboarding-example` to register DevEUIs
- Add responses to confirm success or failure of API consumption
- Add queue mechanism to API to handle a maximum of 10 in-flight API requests concurrently
- Generate 100 unique DevEUIs
- Create function to handle all in-flight requests before exiting the process
- Add app termination response with DevUIs successfully registered (including in-flight ones)
