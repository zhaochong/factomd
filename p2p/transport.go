// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package p2p

import (
	"encoding/gob"
	"fmt"
	"hash/crc32"
	"io"
	"net"
	"syscall"
	"time"
)

// The Transpoort interface allows the choice of different network protocols for lower level
// TCP communicaiton.  All of the functionality of connecting an disconnecting and communicating
// with peers is abstracted here.

// A transport needs to wrap a TCP connection and provide the functions below.

type Transport interface {
	Init(nameTo, nameFrom string) IPeer // Name of peer
	GetNameTo() string                  // Return the name of the peer
	GetNameFrom() string                // Return the name of the peer
	Send(IMsg) error                    // Send a message to this peer
	Recieve() (IMsg, error)             // Recieve a message from this peer; nil if no message is ready.
	Len() int                           // Returns the number of messages waiting to be read
	Equals(IPeer) bool                  // Is this connection equal to parm connection
	Weight() int                        // How many nodes does this peer represent?
}

CurrentNetworkTransport =