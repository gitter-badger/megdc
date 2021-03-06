#!/bin/bash

if [ -f /etc/lsb-release ]; then
    . /etc/lsb-release
    DISTRO=$DISTRIB_ID
elif [ -f /etc/debian_version ]; then
    DISTRO=Debian
    # XXX or Ubuntu
elif [ -f /etc/redhat-release ]; then
    DISTRO="Red Hat"
    # XXX or CentOS or Fedora
else
    DISTRO=$(uname -s)
fi



if [ "$DISTRO" = "Red Hat" ]  || [ "$DISTRO" = "Ubuntu" ] || [ "$DISTRO" = "Debian" ]
then

cat > //usr/share/megam/megamgulpd/conf/gulpd.conf << 'EOF'

### Welcome to the Gulpd configuration file.

  ###
  ### [meta]
  ###
  ### Controls the parameters for the Raft consensus group that stores metadata
  ### about the gulp.
  ###

  [meta]
    debug = true
    hostname = "localhost"
    bind_address = "192.168.1.105:7777"
    dir = "/var/lib/megam/gulp/meta"
    riak = ["192.168.1.105:8087"]
    api  = "https://api.megam.io/v2"
    amqp = "amqp://guest:guest@192.168.1.105:5672/"

  ###
  ### [gulpd]
  ###
  ### Controls which assembly to be deployed into machine
  ###

  [gulpd]
    name_gulp = "hostname"
    cats_id = "AMS1259077729232486400"
    cat_id = "ASM1260230009767985152"
	provider = "chefsolo"
	cookbook = "megam_run"
	repository = "github"
	repository_path = "https://github.com/megamsys/chef-repo.git"

  ###
  ### [http]
  ###
  ### Controls how the HTTP endpoints are configured. This a frill
  ### mechanism for pinging gulpd (ping)
  ###

  [http]
    enabled = true
    bind-address = "localhost:6666"

EOF

sed -i "s/^[ \t]*name_gulp.*/    name = \"$NODE_NAME\"/" /usr/share/megam/megamgulpd/conf/gulpd.conf
sed -i "s/^[ \t]*cats_id.*/    cats_id = \"$ASSEMBLIES_ID\"/" /usr/share/megam/megamgulpd/conf/gulpd.conf
sed -i "s/^[ \t]*cat_id.*/    cat_id = \"$ASSEMBLY_ID\"/" /usr/share/megam/megamgulpd/conf/gulpd.conf

fi




case "$DISTRO" in
   "Ubuntu")
stop megamgulpd
start megamgulpd
   ;;
   "Debian")
#systemctl stop megamgulpd.service
systemctl start megamgulpd.service
   ;;
   "Red Hat")
#systemctl stop megamgulpd.service
systemctl start megamgulpd.service
   ;;
   "CoreOS")
if [ -f /mnt/context.sh ]; then
  . /mnt/context.sh
fi

sudo cat >> //home/core/.ssh/authorized_keys <<EOF
$SSH_PUBLIC_KEY
EOF

sudo -s

sudo cat > //etc/hostname <<EOF
$HOSTNAME
EOF

sudo cat >> //etc/hosts <<EOF
$IP_ADDRESS $HOSTNAME localhost

EOF

sudo cat > //etc/systemd/network/static.network <<EOF
[Match]
Name=ens3

[Network]
Address=$IP_ADDRESS/24
Gateway=$GATEWAY
EOF

sudo systemctl restart systemd-networkd

   ;;
esac
