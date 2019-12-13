# Install OPS

source: https://ops.city/

```bash
curl https://ops.city/get.sh -sSfL | sh
```

# Install KVM

```bash
sudo apt-get install qemu-kvm libvirt-bin virtinst bridge-utils cpu-checker
```

## Compile Go app

```bash
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo  -ldflags '-extldflags "-static"'
```

## Run unikernel
```bash
ops run unikernels
```

```bash
booting /home/wkm0108a/.ops/images/unikernels.img ...
qemu-system-x86_64: warning: TCG doesn't support requested feature: CPUID.01H:ECX.vmx [bit 5]
assigned: 10.0.2.15
2019/12/13 11:03:38 Listening...on 8080
```

## Expose unikernel kvm ports

```bash
ops run -v -p 8080 -c config.json unikernels
```

## Verbose output

```bash
booting /home/wkm0108a/.ops/images/unikernels.img ...
qemu-system-x86_64 -drive file=/home/wkm0108a/.ops/images/unikernels.img,format=raw,if=none,id=hd0 \
	-device virtio-blk,drive=hd0 \
	-device virtio-net,netdev=n0 \
	-netdev user,id=n0,hostfwd=tcp::8080-:8080 \
	-nodefaults -no-reboot -device isa-debug-exit \
	-m 2G -display none \
	-serial stdio
qemu-system-x86_64: warning: TCG doesn't support requested feature: CPUID.01H:ECX.vmx [bit 5]
assigned: 10.0.2.15
2019/12/13 11:13:15 Listening...on 8080
```

## Optimizations

### Enable all CPUs and KVM virtualization

Add following flags to `qemu` command

```bash
-smp $(nproc) -enable-kvm -cpu host
```

```bash
qemu-system-x86_64 -drive file=/home/wkm0108a/.ops/images/unikernels.img,format=raw,if=none,id=hd0 \
	-smp $(nproc) -enable-kvm -cpu host \
	-device virtio-blk,drive=hd0 \
	-device virtio-net,netdev=n0 \
	-netdev user,id=n0,hostfwd=tcp::8080-:8080 \
	-nodefaults -no-reboot -device isa-debug-exit \
	-m 2G -display none \
	-serial stdio
```