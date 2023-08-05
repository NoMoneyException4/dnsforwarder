/*
Dns Forwarder, a dns server with temporary cache and hosts support.

Usage:

	sudo dnsforwader [options]

	Use sudo command to gain the root privileges for dns default port 53.

Options:

	-c Configuration file path (default: ./dnsforwarder.yml)
	-h Server host (default: 127.0.0.1)
	-p Server port (default: 53)
	-d Enable debug mode (default:false)
*/
package main
