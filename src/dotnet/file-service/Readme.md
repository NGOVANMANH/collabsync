## FFmpeg Setup

Before using this project, ensure FFmpeg is installed and available in your system's PATH.

### Installation

#### Windows

1. Download FFmpeg from [ffmpeg.org/download.html](https://ffmpeg.org/download.html).
2. Extract the archive.
3. Add the `bin` folder to your system's PATH environment variable.

#### macOS

```sh
brew install ffmpeg
```

#### Linux (Debian/Ubuntu)

```sh
sudo apt update
sudo apt install ffmpeg
```

### Verify Installation

Run the following command to verify FFmpeg is installed:

```sh
ffmpeg -version
```