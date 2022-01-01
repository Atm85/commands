package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"strconv"
)

type Teleport struct {
	X string
	Y string
	Z string

	// TODO - implement teleporting between worlds
	World string `optional:""`
}

func (c Teleport) Run(source cmd.Source, output *cmd.Output) {

	pos := mgl64.Vec3{}

	// convert string values to float64
	var x, y, z float64
	var err error

	x, err = strconv.ParseFloat(c.X, 32)
	if err != nil {
		if c.X == "~" {
			x = source.Position().X()
		} else {
			output.Errorf("Cannot parse float: %v", err.Error())
		}
	}

	y, err = strconv.ParseFloat(c.Y, 32)
	if err != nil {
		if c.Y == "~" {
			y = source.Position().Y()
		} else {
			output.Errorf("Cannot parse float: %v", err.Error())
		}
	}

	z, err = strconv.ParseFloat(c.Z, 32)
	if err != nil {
		if c.Z == "~" {
			z = source.Position().Z()
		} else {
			output.Errorf("Cannot parse float: %v", err.Error())
		}
	}

	pos[0], pos[1], pos[2] = x, y, z
	source.(*player.Player).Teleport(pos)
	source.(*player.Player).Messagef("You teleported to: %v %v %v", int(pos.X()), int(pos.Y()), int(pos.Z()))
}
