VENV="/tmp/venv/go-proxmox"
mkdir -p $VENV

python3 -m venv ${VENV}
source /tmp/venv/go-proxmox/bin/activate

pip3 install mkdocs mkdocs-material
