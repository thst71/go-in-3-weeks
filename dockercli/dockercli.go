package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const version = "V1.0.1"

const ecode_noconfig = 1
const ecode_noprofile = 2
const ecode_badcmdline = 3
const ecode_failed = 4

const defaultProfileName = "default"

func main() {

	pErr("running %s on %s/%s\n", version, runtime.GOOS, runtime.GOARCH)

	appOptions, err := readOptions()

	if err != nil {
		pErr("Commandline: dockercli [--tags] [--profile <profilename>] <dockerfile>\nError: %s\n", err)
		os.Exit(ecode_badcmdline)
	}

	if appconfig, err := readConfig(); err == nil {
		profileName := appconfig.defaultProfile
		if appOptions.profile != defaultProfileName {
			profileName = appOptions.profile
		}

		if profile, ok := appconfig.profiles[profileName]; !ok {
			pErr("Unknown config profile %s, cannot continue\n", profileName)
			os.Exit(ecode_noprofile)
		} else {
			command := "tags/list"
			getTagsUrl := fmt.Sprintf(profile.registry+"%s/%s/%s", "/v2", appOptions.image, command)
			if getTagsReq, err := http.NewRequest("GET", getTagsUrl, bytes.NewBuffer([]byte{})); err == nil {
				if profile.username != "" {
					getTagsReq.SetBasicAuth(profile.username, profile.password)
				}
				client := &http.Client{}
				if getTagsResp, err := client.Do(getTagsReq); err == nil {
					defer getTagsResp.Body.Close()
					getTagsBody, _ := ioutil.ReadAll(getTagsResp.Body)
					pErr("Status: %s\nBody:\n", getTagsResp.Status)
					fmt.Println(string(getTagsBody))
				} else {
					pErr("Could not load tags from registry %s : %s\n", profile.registry, err)
					os.Exit(ecode_failed)
				}
			}
		}

	} else {
		pErr("Unable to load configfile %s\n", err)
		os.Exit(ecode_noconfig)
	}

}

type opmode int

const (
	op_tags opmode = iota
)

type appoptions struct {
	mode    opmode
	profile string
	image   string
}

func readOptions() (appoptions, error) {
	// fetchTags := *flag.Bool("tags", true, "fetch tags of dockerfile")
	useProfile := *flag.String("profile", defaultProfileName, "configprofile, defaults to default profile")

	flag.Parse()

	if flag.NArg() == 0 {
		return appoptions{}, errors.New("no docker image")
	}

	return appoptions{op_tags, useProfile, flag.Arg(0)}, nil
}

type config struct {
	defaultProfile string
	profiles       map[string]configProfile
}

type configProfile struct {
	id       string
	username string
	password string
	registry string
}

func (c *config) add(cp configProfile) {
	if c.profiles == nil {
		c.profiles = make(map[string]configProfile)
	}
	c.profiles[cp.id] = cp
}

func readConfig() (config, error) {
	home, _ := os.UserHomeDir()
	configData, err := ioutil.ReadFile(fmt.Sprintf("%s/.dockercli/config", home))
	if err != nil {
		return config{}, err
	}

	configString := string(configData)

	cfg := &config{}

	var currProfile *configProfile

	for _, rawline := range strings.Split(configString, "\n") {
		line := strings.TrimSpace(rawline)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Fields(line)

		if fields[0] == "defaultprofile" {
			if len(fields) < 2 {
				return config{}, errors.New("missing defaultprofile name in " + line)
			}

			cfg.defaultProfile = fields[1]
		} else if fields[0] == "profile" {
			if len(fields) < 2 {
				return config{}, errors.New("missing profile id in " + line)
			}

			pErr("reading %s", line)

			if currProfile != nil {
				cfg.add(*currProfile)
			}
			currProfile = &configProfile{id: fields[1]}
		} else if fields[0] == "username" {
			if currProfile == nil {
				return config{}, errors.New("username outside profile definition " + line)
			}
			if len(fields) < 2 {
				return config{}, errors.New("missing username in " + line)
			}

			currProfile.username = fields[1]
		} else if fields[0] == "password" {
			if currProfile == nil {
				return config{}, errors.New("password outside profile definition " + line)
			}
			if len(fields) < 2 {
				return config{}, errors.New("missing password in " + line)
			}

			currProfile.password = fields[1]
		} else if fields[0] == "registry" {
			if currProfile == nil {
				return config{}, errors.New("registry outside profile definition " + line)
			}
			if len(fields) < 2 {
				return config{}, errors.New("missing registry in " + line)
			}

			currProfile.registry = fields[1]
		}

	}

	if currProfile != nil {
		cfg.add(*currProfile)
	}

	return *cfg, nil
}

func pErr(format string, params ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, params)
}
