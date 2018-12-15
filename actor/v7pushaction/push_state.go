package v7pushaction

import (
	"fmt"

	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"code.cloudfoundry.org/cli/types"
	log "github.com/sirupsen/logrus"
)

type PushState struct {
	Application v7action.Application
	SpaceGUID   string
	OrgGUID     string
	Overrides   FlagOverrides
	Manifest    []byte

	Archive            bool
	BitsPath           string
	AllResources       []sharedaction.Resource
	MatchedResources   []sharedaction.Resource
	UnmatchedResources []sharedaction.Resource
}

type FlagOverrides struct {
	Buildpacks        []string
	DockerImage       string
	DockerPassword    string
	DockerUsername    string
	HealthCheckType   constant.HealthCheckType
	Memory            types.NullUint64
	ProvidedAppPath   string
	SkipRouteCreation bool
	StartCommand      types.FilteredString
	NoStart           bool
}

func (state PushState) String() string {
	return fmt.Sprintf(
		"Application: %#v - Space GUID: %s, Org GUID: %s, Archive: %t, Bits Path: %s",
		state.Application,
		state.SpaceGUID,
		state.OrgGUID,
		state.Archive,
		state.BitsPath,
	)
}

func (actor Actor) Conceptualize(appName string, spaceGUID string, orgGUID string, currentDir string, flagOverrides FlagOverrides, manifest []byte) ([]PushState, Warnings, error) {
	var (
		application v7action.Application
		warnings    v7action.Warnings
		err         error
	)

	log.WithFields(log.Fields{"appName": appName, "spaceGUID": spaceGUID}).Info("Application Look Up")
	application, warnings, err = actor.V7Actor.GetApplicationByNameAndSpace(appName, spaceGUID)
	if _, ok := err.(actionerror.ApplicationNotFoundError); ok {
		log.Debug("App not found")
		application = v7action.Application{
			Name: appName,
		}
	} else if err != nil {
		log.Errorln("Looking up application:", err)
		return nil, Warnings(warnings), err
	}

	if len(flagOverrides.Buildpacks) != 0 {
		application.LifecycleType = constant.AppLifecycleTypeBuildpack
		application.LifecycleBuildpacks = flagOverrides.Buildpacks
	}

	if len(flagOverrides.DockerImage) != 0 {
		application.LifecycleType = constant.AppLifecycleTypeDocker
	}

	bitsPath := currentDir
	if flagOverrides.ProvidedAppPath != "" {
		bitsPath = flagOverrides.ProvidedAppPath
	}

	resources, err := actor.SharedActor.GatherDirectoryResources(bitsPath)

	desiredState := []PushState{
		{
			Application: application,
			SpaceGUID:   spaceGUID,
			OrgGUID:     orgGUID,
			Overrides:   flagOverrides,
			Manifest:    manifest,

			BitsPath:     bitsPath,
			AllResources: resources,
		},
	}
	return desiredState, Warnings(warnings), err
}
