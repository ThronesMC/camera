package instructions

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

var (
	FreeCameraPreset             = free
	FirstPersonCameraPreset      = firstPerson
	ThirdPersonCameraPreset      = thirdPerson
	ThirdPersonFrontCameraPreset = thirdPersonFront
	FollowOrbitCameraPreset      = followOrbit
	FixedBoomCameraPreset        = fixedBoom

	Presets = presets
)

var (
	presets = []protocol.CameraPreset{FreeCameraPreset, FirstPersonCameraPreset, ThirdPersonCameraPreset, ThirdPersonFrontCameraPreset, FollowOrbitCameraPreset, FixedBoomCameraPreset}
	indexes = map[protocol.CameraPreset]uint32{free: 0, firstPerson: 1, thirdPerson: 2, thirdPersonFront: 3, followOrbit: 4, fixedBoom: 5}

	free = protocol.CameraPreset{
		Name: "minecraft:free",
		//MinYawLimit:   protocol.Optional[float32]{},
		//MaxYawLimit:   protocol.Optional[float32]{},
		AudioListener: protocol.Option(byte(protocol.AudioListenerPlayer)),
	}
	firstPerson = protocol.CameraPreset{
		Name: "minecraft:first_person",
		//MinYawLimit:   protocol.Optional[float32]{},
		//MaxYawLimit:   protocol.Optional[float32]{},
		AudioListener: protocol.Option(byte(protocol.AudioListenerPlayer)),
	}
	thirdPerson = protocol.CameraPreset{
		Name: "minecraft:third_person",
		//MinYawLimit:   protocol.Optional[float32]{},
		//MaxYawLimit:   protocol.Optional[float32]{},
		AudioListener: protocol.Option(byte(protocol.AudioListenerPlayer)),
	}
	thirdPersonFront = protocol.CameraPreset{
		Name: "minecraft:third_person_front",
		//MinYawLimit:   protocol.Optional[float32]{},
		//MaxYawLimit:   protocol.Optional[float32]{},
		AudioListener: protocol.Option(byte(protocol.AudioListenerPlayer)),
	}
	followOrbit = protocol.CameraPreset{
		Name: "minecraft:follow_orbit",
		//MinYawLimit:   protocol.Optional[float32]{},
		//MaxYawLimit:   protocol.Optional[float32]{},
		AudioListener: protocol.Option(byte(protocol.AudioListenerPlayer)),
	}
	fixedBoom = protocol.CameraPreset{
		Name: "minecraft:fixed_boom",
		//MinYawLimit:   protocol.Optional[float32]{},
		//MaxYawLimit:   protocol.Optional[float32]{},
		AudioListener: protocol.Option(byte(protocol.AudioListenerPlayer)),
	}
)
