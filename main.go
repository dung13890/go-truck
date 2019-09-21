package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os/exec"
    "sync"

    yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
    WebServers
}

type WebServers struct {
    Hosts map[string]Server `yaml:"hosts"`
    Tasks []Task
}

type Server struct {
    Address string `yaml:"address"`
    User    string `yaml:"user"`
    Dir     string `yaml:"dir"`
}

type Task struct {
    Name string   `yaml:"name"`
    Cmd  string   `yaml:"cmd"`
    Args []string `yaml:"args"`
}

func (c *Configuration) getConf() *Configuration {

    yamlFile, err1 := ioutil.ReadFile("config.yml")
    if err1 != nil {
        log.Printf("yamlFile.Get err   #%v ", err1)
    }

    err2 := yaml.Unmarshal(yamlFile, &c)
    if err2 != nil {
        log.Fatalf("error: %v", err2)
    }

    return c
}

func Connection(s Server) {
    cmd := exec.Command(
        "ssh",
        "-t",
        fmt.Sprintf("%s@%s", s.User, s.Address),
        fmt.Sprintf("'cd %s ; ls -al'", s.Dir),
    )
    stdout, err := cmd.Output()
    if err != nil {
        println(err.Error())
        return
    }

    print(string(stdout))
}

func main() {
    c := Configuration{}
    c.getConf()
    var wg sync.WaitGroup
    for _, server := range c.WebServers.Hosts {
        wg.Add(1)
        go Connection(server)
    }
    wg.Wait()
}
