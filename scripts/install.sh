os=""
arch=""

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
	os="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
	os="darwin"
else
	echo "error: os type '$OSTYPE' not supported"
	exit 1
fi

if [[ $(uname -m) == x86_64* ]]; then
    arch="amd64"
elif [[ $(uname -m) == i*86 ]]; then
    arch="386"
elif [[ $(uname -m) == aarch64 ]]; then
    arch="arm64"
elif [[ $(uname -m) == arm64 ]]; then
    arch="arm64"
elif  [[ $(uname -m) == arm* ]]; then
    arch="arm"
else
	echo "error: os arch '$(uname -m)' not supported"
	exit 1
fi

bin="/usr/local/bin/pve-api-wrapper"
url="http://github.com/iolave/go-proxmox/releases/latest/download/pve-api-wrapper-$os-$arch"
echo "info: downloading pve-api-wapper-${os}-${arch} binary"
status_code=$(curl -L --write-out %{http_code} --silent --output /dev/null $url)
if [[ "$status_code" != "200" ]]; then
	echo "error: got status code $status_code while downloading binary"
	exit 1
fi

sudo curl -L -o $bin $url 
chmod +x $bin
