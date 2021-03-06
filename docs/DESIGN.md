- Management Server
- Nodes
    - Server - communicates with the management server using gRPC
        API:
            - Start VM (cpu, memory, network)
            - Stop VM

- Upon adding a node firecracker and additional tools (gRPC client/server) will be installed

- VM
    - will be assigned a uuid
    - the uuid will be used for the unix socket path as well
    - create a tap device for the VM, same uuid will be used
    - need to save VMs on the host to keep track (sqlite?)


# Example requests
```
curl -XPOST http://localhost:8888/hosts\?install\=true -d '{ "name":"hosto", "address": "192.168.122.45", "user": "root", "password": "centos", "rootfs": "rootfs", "kernel": "vmlinux", "local_node_path": "/home/benny/Development/go/src/github.com/PUMATeam/catapult-node", "port": 8001}'
```
```
curl -XPOST http://localhost:8888/vms -d '{ "name": "hello", "vcpu": 1, "memory": 128, "kernel": "vmlinux", "rootfs": "rootfs" }'
```

# Stuff to think about
- How to load existing VMs when catapult-node starts?
  - Use an sqlite3 db and consider it the only source of truth
  - Scan the /var/vms
  - Combine both options above - but adding a vms from /var/vms to the db
might be challenging as it would require learning about the VM specifications
from the API (there's no "virsh dumpxml") as far as I know

# Handling host state
- Upon starting we should check all the present hosts state
  - If a host is UP, we need to try and health check it using grpc
    - if it failes to respond it might mean the node process is down (should we turn it into a systemd service?), move it into DOWN state
  - We should not accept any request until we determined the host states

- We should be able to de/activate a host - an ansible script to start/stop
  the catapult-node process

# Storage
- catapult-storage will leverage cinderlib to CRUD ceph volumes
  - How does it communicate? Direct gRPC? Do we need 0mq in the middle?
  	- Maybe a "mutlicast" operation, send a "create volume" request and the first available
    	  host will attempt to fulfil it
	- Suppose a host failed to complete the request in a reasonable time, the manager can
	  retransmit the request and another host will pick it up. If the original host did not
	  realize it was retransmitted, it would be fenced and the host will rollback.
