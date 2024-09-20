# YourApp

**YourApp** is a minimal command-line tool built in Go that allows users to process both video and image files with ease. It leverages the power of [FFmpeg](https://ffmpeg.org/) under the hood via the [ffmpeg-go](https://github.com/u2takey/ffmpeg-go) library, offering a clean and modular command structure for common media processing tasks.

## Features

- **Video Conversion**: Convert video files between different formats.
- **Image Conversion**: Convert image files between different formats (e.g., JPG to PNG).
- **Image Rotation**: Rotate images by a specified degree.
- Built with the **Cobra** framework for intuitive CLI commands.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go** 1.17 or higher installed.
- **FFmpeg** installed and accessible in your system's `PATH`.

You can install FFmpeg by following the instructions [here](https://ffmpeg.org/download.html).

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/yourusername/yourapp.git
   cd yourapp
