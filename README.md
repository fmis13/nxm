<p align="center">
    <img src="https://avatars.githubusercontent.com/u/120988076?s=200&v=4" width=128 height=128>

<h2 align="center">nxm</h2>

<p align="center">
    The package manager for linux (home-)servers.
    <br>
    <a href="https://github.com/nixmember/nixmember/issues/new">Report bug(s)</a>
    ·
    <a href="https://github.com/nixmember/nixmember/issues/new">Feature request(s)</a>
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

In order to compile `nxm` you need to have `go` and `git` installed, and run the following commands:
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