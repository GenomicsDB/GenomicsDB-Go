/**
 *
 * The MIT License
 *
 * Copyright (c) 2023 dātma, inc™
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */

package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"runtime"

	_ "embed"
)

//go:embed install.sh
var script []byte

func InstallNativeGenomicsDB() {
	const GOOS string = runtime.GOOS

	installScript, err := os.CreateTemp("", "install.sh")
	if err != nil {
		log.Fatal("Could not read embedded install.sh script: ", err)
	}
	defer os.Remove(installScript.Name()) // clean-up when done

	file, err := os.OpenFile(installScript.Name(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Could not open temporary file for install script: ", err)
	}
	if _, err = file.Write(script); err != nil {
		log.Fatal("Could not write install script to temporary file: ", err)
	}
	if err = file.Close(); err != nil {
		log.Fatal("Could not close temporary file for install script: ", err)
	}
	if err = os.Chmod(installScript.Name(), 0755); err != nil {
		log.Fatal("Could not chmod install script: ", err)
	}

	var cmd *exec.Cmd
	ctx, _ := context.WithCancel(context.Background())
	if GOOS == "darwin" || GOOS == "linux" {
		cmd = exec.CommandContext(ctx, installScript.Name())
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		log.Fatal("Unsupported OS: ", GOOS)
	}

	if err = cmd.Start(); err != nil {
		log.Fatal("Could not start install script from exec: ", err)
	}
	if err = cmd.Wait(); err != nil {
		log.Fatal("Could not successfully wait for install script to return: ", err)
	}
}

func main() {
	log.Println("Installing Native GenomicsDB...")
	InstallNativeGenomicsDB()
	log.Println("Installing Native GenomicsDB DONE")
}
