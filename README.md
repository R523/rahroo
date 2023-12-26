# Rahroo 🛣️

## Introduction

This program is designed to differentiate between white and black surfaces using a light sensor.
The sensor operates as a variable resistor that alters voltage levels in response to light intensity, necessitating the use of an analog signal for accurate detection.
To facilitate this, we’re employing an MPC3008 ADC, known for its SPI (Serial Peripheral Interface) compatibility.
The ADC has been configured to integrate seamlessly with the Raspberry Pi’s hardware SPI capabilities, allowing for precise and real-time color sensing critical for the application’s requirements.
