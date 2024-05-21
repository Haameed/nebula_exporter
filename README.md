# OpenNebula

OpenNebula is a tool that enables monitoring of various objects in a Prometheus-friendly manner.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
  - [Clone and Build](#clone-and-build)
  - [Download from Releases](#download-from-releases)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Features

- Prometheus-friendly monitoring
- Easy setup and configuration
- Supports objects like:
  - Hosts:
    - Host status
    - Max CPU
    - Total CPU
    - Used CPU
    - Memory Max
    - Memory Total
    - Memory allocation
  - VMs:
    - VM status
    - VM vCPUs
    - VM virtual memory
  - Datastores:
    - Datastore status
    - Datastore Total capacity MB
    - Datastore Used MB
    - DataStore Free MB
  - Virtual Networks
    - vNet total IPs
    - vNet total leased
## Installation

### Clone and Build

To clone the project and build it locally, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/Haameed/nebula_exporter.git
    ```
2. Navigate to the project directory:
    ```bash
    cd nebula_eexporter
    ```
3. Build the project:
    ```bash
    go build .
    ```

### Download from Releases
#### 
To download the latest release, follow these steps:

1. Go to the [releases page](https://github.com/Haameed/nebula_exporter/releases).
2. Download the desired version for your platform.
3. Extract the downloaded file and follow the usage instructions below.

## Usage


## Configuration
To generate a sample configuration use the following command 
```bash
  nebula_exporter -generate-config true
```
This command creates a sample configuration in /tmp/sample.json
## Contributing

I welcome contributions from everyone! Whether fixing a bug, improving documentation, or adding new features, your help is greatly appreciated.

To contribute, please follow these steps:

1. **Fork the repository**: Click the "Fork" button at the top of this page to create a copy of the repository under your GitHub account.

2. **Clone your fork**: Use the following command to clone your forked repository to your local machine.
    ```bash
    git clone https://github.com/Haameed/nebula_exporter.git
    ```

3. **Create a branch**: Create a new branch for your changes.
    ```bash
    git checkout -b feature-branch
    ```

4. **Make your changes**: Make your changes to the codebase.

5. **Commit your changes**: Commit your changes with a descriptive commit message.
    ```bash
    git add .
    git commit -m "Description of your changes"
    ```

6. **Push to your fork**: Push your changes to your forked repository.
    ```bash
    git push origin feature-branch
    ```

7. **Create a pull request**: Go to the original repository and create a pull request from your forked repository. Please clearly describe the changes you have made and any related issue numbers.

I appreciate your effort and will review your pull request as soon as possible


After installing OpenNebula, you can start using this exporter with the following command:

```bash

./nebula_exporter -config config.json
```
## License
MIT
