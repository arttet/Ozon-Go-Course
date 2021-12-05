
package main

import (
	"log"
	"net/http"

	"github.com/spf13/afero"
	"github.com/yandex/pandora/core/aggregator/netsample"
	"google.golang.org/grpc"

	"github.com/yandex/pandora/cli"
	phttp "github.com/yandex/pandora/components/phttp/import"
	"github.com/yandex/pandora/core"
	coreimport "github.com/yandex/pandora/core/import"
	"github.com/yandex/pandora/core/register"

	desc "github.com/ozonmp/omp-edu-lessons/week-6/lecture-14/1-load-testing/api"
)

type Ammo struct {
	Val        int32
}

type Sample struct {
	URL              string
	ShootTimeSeconds float64
}

type GunConfig struct {
	Target string `validate:"required"` // Configuration will fail, without target defined
}

type Gun struct {
	// Configured on construction.
	conf GunConfig
	client grpc.ClientConn

	// Configured on Bind, before shooting.
	aggr core.Aggregator // May be your custom Aggregator.
	core.GunDeps
}

func NewGun(conf GunConfig) *Gun {
	return &Gun{conf: conf}
}

func (g *Gun) Bind(aggr core.Aggregator, deps core.GunDeps) error {
	conn, err := grpc.Dial(
		g.conf.Target,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithUserAgent("pandora load test"))
	if err != nil {
		log.Fatalf("FATAL: %s", err)
	}
	g.client = *conn
	g.aggr = aggr
	g.GunDeps = deps
	return nil

}

func (g *Gun) Shoot(ammo core.Ammo) {
	customAmmo := ammo.(*Ammo) // Shoot will panic on unexpected ammo type. Panic cancels shooting.
	g.shoot(customAmmo)
}

func (g *Gun) shoot(ammo *Ammo) {
	conn := g.client
	client := desc.NewTestAPIClient(&conn)

	sample := netsample.Acquire("test")

	req := http.Request{}
	out, err := client.TestAPIHandler(req.Context(), &desc.TestAPIHandlerRequest{
		Val: ammo.Val,
	})

	code := 0
	if err != nil {
		code = 0
	}

	if out != nil {
		code = 200
	}

	defer func() {
		sample.SetProtoCode(code)
		g.aggr.Report(sample)
	}()

}

func main() {
	// Standard imports.
	fs := afero.NewOsFs()
	coreimport.Import(fs)

	// May not be imported, if you don't need http guns and etc.
	phttp.Import(fs)

	// Custom imports. Integrate your custom types into configuration system.
	coreimport.RegisterCustomJSONProvider("my-custom-provider-name", func() core.Ammo { return &Ammo{} })

	register.Gun("my-custom-gun-name", NewGun, func() GunConfig {
		return GunConfig{
			Target: "default target",
		}
	})
	register.Gun("my-custom/no-default", NewGun)

	cli.Run()
}