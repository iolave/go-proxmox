VENV="${TMPDIR}/venv/go-proxmox"
mkdir -p $VENV

python3 -m venv ${VENV}
source "${VENV}/bin/activate"

pip3 install \
    mkdocs \
    mkdocs-material \
    pymdown-extensions \
    markdown-callouts \
    mkdocs-render-swagger-plugin
