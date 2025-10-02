# Ivan

![Go Report Card](https://goreportcard.com/badge/github.com/deniskorbakov/ivan)
![Release](https://img.shields.io/github/release/deniskorbakov/ivan?status.svg)
![Action Lint](https://github.com/deniskorbakov/ivan/actions/workflows/lint.yml/badge.svg)
![GitHub Repo stars](https://img.shields.io/github/stars/deniskorbakov/ivan)

Ivan - Auto build and deploy pipline on Go with Ansible

Made using data from packages:

* [cobra](https://github.com/spf13/cobra)
* [fang](http://github.com/charmbracelet/fang)
* [huh](https://github.com/charmbracelet/huh)

## ✨ Install

You will need [Go](https://go.dev/doc/install) and [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html) to work.

You can also install the package from the release via the [link](https://github.com/deniskorbakov/ivan/releases)

Local:

clone the repository

```bash
git clone https://github.com/deniskorbakov/ivan.git
````

go to the project folder

```bash
cd ivan
````

build the app

```bash
make build
```

launch the app

```bash
./ivan
```

## 📖 Examples & Usage

The list of commands that you can use in this Build and Deploy 

### 🔌 Build

The command to get rep info and build without ansible

For ansible to work over ssh, you must transfer your key to the server.

```bash
ivan build
```

<img src=".assets/build.svg" width="650">

## 🧪 Testing

The command to launch the linter:

```bash
make lint
```

## 🤝 Feedback

We appreciate your support and look forward to making our product even better with your help!

[@Denis Korbakov](https://github.com/deniskorbakov)
