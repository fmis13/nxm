<p align="center">
    <img src="https://avatars.githubusercontent.com/u/120988076?s=200&v=4" width=128 height=128>

<h2 align="center">nxm</h2>

<p align="center">
    The package manager for Linux (home-)servers.
    <br>
    <a href="https://github.com/nixmember/nixmember/blob/master/README.md#installing-via-the-installation-script">Install</a>
    ·
    <a href="https://github.com/nixmember/nixmember/issues/new/choose">Request feature or report bug</a>
    ·
    <a href="https://github.com/nixmember/nixmember/blob/master/CONTRIBUTING.md">Contributing</a>
    </p>
</p>

## Getting Started

### Installing via the installation script
The installation script requires `docker`, `wget`, `grep`, `jq` and `curl`

##### System-wide installation:
```bash
curl https://raw.githubusercontent.com/nixmember/nixmember/master/install.sh | sudo bash -s
```

**Note:** Always check bash scripts before running them as sudo. Feel free to check out our installation script, it's safe.

##### Local installation
```bash
curl https://raw.githubusercontent.com/nixmember/nixmember/main/install.sh | bash -s
```

### Compiling

To compile `nxm` you need to have `go` and` `git` installed, and run the following commands:

```bash
git clone https://github.com/nixmember/nixmember && cd nixmember
sudo make install
```
The binary will now build in the directory, you can move the binary to the `/usr/bin` directory by running the following command:

```bash
sudo mv nxm /usr/bin
```
or for local installation:

```bash
mkdir ~/.local/bin
mv nxm ~/.local/bin
```