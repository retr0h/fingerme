// Copyright (c) 2024 John Dewey

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package server

import (
	"bufio"
	_ "embed"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strconv"
)

//go:embed resources/plan.txt
var plan []byte

// New factory to create a new Server instance.
func New(
	logger *slog.Logger,
	port int,
) *Server {
	return &Server{
		logger: logger,
		port:   port,
	}
}

func (s *Server) finger(conn net.Conn) {
	defer func() { _ = conn.Close() }()

	reader := bufio.NewReader(conn)
	user, _, _ := reader.ReadLine()

	s.logger.Info("handling request",
		slog.Int("port", s.port),
		slog.String("user", string(user)),
	)

	// magic user
	if string(user) == "john" {
		_, _ = conn.Write(plan)
	} else {
		_, _ = conn.Write([]byte(`{"desc": "nothing"}`))
	}

	_ = conn.Close()
}

func (s *Server) fingering() int {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		fmt.Printf("failed to bind to port %d:n%s\n", s.port, err.Error())
		return 1
	}

	s.logger.Info("staring fingerd server",
		slog.Int("port", s.port),
	)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go s.finger(conn)
	}
}

func (s *Server) Serve() {
	os.Exit(s.fingering())
}
