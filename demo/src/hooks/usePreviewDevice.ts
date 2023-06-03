import { useMediaDevices } from "@livekit/components-react"
import {
	LocalAudioTrack,
	LocalVideoTrack,
} from "livekit-client"
import { useState } from "react"

const usePreviewDevice = <T extends LocalVideoTrack | LocalAudioTrack>(
	enabled: boolean,
	deviceId: string,
	kind: MediaDeviceKind 
) => {
	if (!deviceId)
		throw new Error("deviceId is required")

	const devices = useMediaDevices({ kind })
	const [error, setError] = useState<Error | null>(null)
	const [localDeviceId, setLocalDeviceId] = useState<string>(deviceId)
	const [localTracks, setLocalTracks] = useState<T>()
	const [selectedDevice, setSelectedDevice] = useState<MediaDeviceInfo | null>(null)

	return <></>

}