/*
Copyright 2021 Aru Moon

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

/*
	Test module package
	Execute to test:
		go test -v ./...
*/

package natsuki

import (
	"os"
	"testing"

	. "github.com/bwmarrin/discordgo" // Calling discordgo without qualifiaction
)

var cl *Session

func TestMain(t *testing.T) {
	cl := Run(os.Getenv("NATSUKI_TOKEN")) // os.Getenv("NATSUKI_TOKEN") can be replaced to your own method to authorize, it should be string

	cl.Close()
	return
}
