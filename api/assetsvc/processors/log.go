package processors

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"time"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/tinyci/ci-agents/grpc/services/asset"
	"github.com/tinyci/ci-agents/grpc/types"
	"google.golang.org/grpc/metadata"
)

const (
	defaultLogsRoot = "/var/tinyci/logs"
)

// AssetServer is the handler anchor for the GRPC network system.
type AssetServer struct{}

// PutLog writes the log to disk
func (as *AssetServer) PutLog(ap asset.Asset_PutLogServer) error {
	return submit(ap)
}

// GetLog spills the log back to connecting websocket.
func (as *AssetServer) GetLog(id *types.IntID, ag asset.Asset_GetLogServer) error {
	return attach(id.ID, ag)
}

func submit(ap asset.Asset_PutLogServer) (retErr error) {
	defer func() {
		if retErr != nil {
			md := metadata.New(nil)
			md.Append("errors", retErr.Error())
			ap.SetTrailer(md)
		}
	}()
	if err := os.MkdirAll(defaultLogsRoot, 0700); err != nil {
		return err
	}

	ls, err := ap.Recv()
	if err != nil {
		return err
	}

	file := path.Join(defaultLogsRoot, fmt.Sprintf("%d", ls.ID))

	if _, err := os.Stat(file); err == nil {
		return errors.New("log already exists")
	}

	if _, err := os.Stat(file + ".writing"); err == nil {
		return errors.New("log writing is currently in progress for this ID")
	}

	writing, err := os.Create(file + ".writing")
	if err != nil {
		return err
	}
	defer func() {
		writing.Close()
		os.Remove(writing.Name())
	}()

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		if _, err := f.Write(ls.Chunk); err != nil {
			return err
		}
		if ls, err = ap.Recv(); err != nil {
			return err
		}
	}
}

func write(ag asset.Asset_GetLogServer, buf []byte) error {
	return ag.Send(&asset.LogChunk{Chunk: buf})
}

func attach(id int64, ag asset.Asset_GetLogServer) error {
	defer write(ag, []byte(color.New(color.FgGreen).Sprintln("---- LOG COMPLETE ----")))

	file := path.Join(defaultLogsRoot, fmt.Sprintf("%d", id))

	log, err := os.Open(file) // #nosec
	if err != nil {
		return err
	}
	defer log.Close()

	buf := make([]byte, 256)
	var (
		last bool
	)

	writeBuf := []byte{}

	for {
		// these two conditionals kind of wrap the main body of the loop to allow
		// the secret sauce; thanks `perldoc -q tail`!
		n, readErr := log.Read(buf)
		if readErr == io.EOF {
			if _, err := os.Stat(file + ".writing"); err != nil {
				return nil // file is done writing
			}

			if _, err := log.Seek(0, io.SeekEnd); err != nil {
				last = true
			} else {
				time.Sleep(500 * time.Millisecond)
			}
		} else if readErr != nil {
			last = true
		}

		writeBuf = append(writeBuf, buf[:n]...)
		// websockets running in textencoding (Which xterm.js requires) require UTF8
		// clean strings to be passed on each write, otherwise it will break the
		// connection.
		// XXX this buffering can probably be abused somehow.
		if utf8.ValidString(string(writeBuf)) {
			if err := write(ag, writeBuf); err != nil {
				if err != io.EOF { // spam-free log experience
					return err
				}

				return nil
			}

			if last {
				if readErr != io.EOF {
					return readErr
				}
			}
			writeBuf = []byte{}
		}
	}
}
