package integration

import (
	"fmt"
	"os/exec"
	"path/filepath"

	. "gopkg.in/check.v1"
)

func (s *RunSuite) TestVolumeFromService(c *C) {
	p := s.RandomProject()
	cmd := exec.Command(s.command, "-f", "./assets/regression/60-volume_from.yml", "-p", p, "create")
	err := cmd.Run()
	c.Assert(err, IsNil)

	volumeFromContainer := fmt.Sprintf("%s_%s_1", p, "first")
	secondContainerName := p + "_second_1"

	cn := s.GetContainerByName(c, secondContainerName)
	c.Assert(cn, NotNil)

	c.Assert(len(cn.HostConfig.VolumesFrom), Equals, 1)
	c.Assert(cn.HostConfig.VolumesFrom[0], Equals, volumeFromContainer)
}

func (s *RunSuite) TestRelativeVolume(c *C) {
	p := s.ProjectFromText(c, "up", `
	server:
	  image: busybox
	  volumes:
	    - .:/path
	`)

	absPath, err := filepath.Abs(".")
	c.Assert(err, IsNil)
	serverName := fmt.Sprintf("%s_%s_1", p, "server")
	cn := s.GetContainerByName(c, serverName)

	c.Assert(cn, NotNil)
	c.Assert(len(cn.Mounts), DeepEquals, 1)
	c.Assert(cn.Mounts[0].Source, DeepEquals, absPath)
}
