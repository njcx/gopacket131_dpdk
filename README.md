# gopacket

This library provides packet decoding capabilities for Go.

Forked from the popular gopacket [repo](https://github.com/google/gopacket) by Google, this fork was created to ensure the project doesn't become stale and bugfixes, new protocols and performance improvements can be merged into it. submit your PRs here :) 


Minimum Go supported is 1.19

```bash

dpdk >= DPDK 20.02.1

kernel >= 3.10.0

CentOS
#  yum install -y libpcap-devel gcc gcc-c++ make meson ninja  numactl-devel  numactl  net-tools pciutils
#  yum install -y kernel-devel-$(uname -r) kernel-headers-$(uname -r)

Debian + Ubuntu
# apt install -y libpcap-dev gcc g++ make meson ninja-build libnuma-dev numactl net-tools pciutils
# apt install -y linux-headers-$(uname -r)


#  wget http://fast.dpdk.org/rel/dpdk-20.11.10.tar.xz
#  tar -Jxvf dpdk-20.11.10.tar.xz
#  cd dpdk-stable-20.11.10 && meson build && cd build && ninja && ninja install
#  export LD_LIBRARY_PATH=/usr/local/lib64:$LD_LIBRARY_PATH
#  git clone https://github.com/njcx/dpdk-kmods
#  cd dpdk-kmods && make
#  modprobe uio  &&  insmod igb_uio.ko
#  /xxx/dpdk-stable-20.11.10/usertools/dpdk-devbind.py -b igb_uio 0000:03:00.0(pci-)
#  CGO_CFLAGS="-msse4.2 -fno-strict-aliasing " CGO_LDFLAGS=" -lrte_eal -lrte_mbuf -lrte_mempool -lrte_ethdev -lpcap" go build


```