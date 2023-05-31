"use client"
import { userStore } from "@/lib/store/prejoin";
import { useMediaDevices } from "@livekit/components-react";
import { log } from "console";
import {
	  LocalAudioTrack,
	  LocalVideoTrack,
	  LocalTrack,
	  createLocalVideoTrack,
	  VideoPresets,
	  createLocalAudioTrack,
} from "livekit-client";
import React, { useCallback } from 'react';
import { useStore } from "zustand/react";

interface PreJoinProps {
	onError: Function,
	onValidate: () => void,
}

const DEBUG = true;
const usePreviewDevice = 
<T extends LocalVideoTrack | LocalAudioTrack> 
(enabled: boolean, deviceId: string, kind: 'audioinput' | 'videoinput') => {
	const [device, setDevice] = React.useState<MediaDeviceInfo>();
	const [deviceError, setDeviceError] = React.useState<Error>();
	const [localTrack, setLocalTrack] = React.useState<T>();
	const [localDeviceId, setLocalDeviceId] = React.useState<string>(deviceId);
	const mediaDevice = useMediaDevices({ kind });
	const prevDeviceId = React.useRef<string>(localDeviceId);

	React.useEffect(() => {

	}, [enabled, deviceId, kind]);

	
	// to keep track of the previous device id and avoid re-rendering when it changes
  const createTrack = useCallback(async (deviceId: string, kind: 'videoinput' | 'audioinput') => {
    try {
      const track =
        kind === 'videoinput'
          ? await createLocalVideoTrack({
              deviceId: deviceId,
              resolution: VideoPresets.h720.resolution,
            })
          : await createLocalAudioTrack({ deviceId });

      const newDeviceId = await track.getDeviceId();
      if (newDeviceId && deviceId !== newDeviceId) {
        prevDeviceId.current = newDeviceId;
        setLocalDeviceId(newDeviceId);
      }
      setLocalTrack(track as T);
    } catch (e) {
      if (e instanceof Error) {
        setDeviceError(e);
      }
    }
  }, [])

  // switch to new device if needed
  const switchDevice = useCallback(async (track: T, deviceId: string) => {
	track.restartTrack({ deviceId });
	prevDeviceId.current = deviceId;
  }, []);

  React.useEffect(() => {
	if (enabled && !deviceError && !localTrack) {
		createTrack(deviceId, kind);
		console.log('creating track', deviceId, kind);
	  return;	
	}

  }, [localTrack, enabled, deviceId, deviceError, kind, createTrack]);

  // switch camera if needed
	React.useEffect(() => {
		if (!enabled) {
		if (localTrack) {
			console.log(`muting ${kind} track`);
			localTrack.mute().then(() => console.log(localTrack.mediaStreamTrack));
		}
		return;
		}
		if (
		localTrack &&
		device?.deviceId &&
		prevDeviceId.current !== device?.deviceId
		) {
		console.log(`switching ${kind} device from`, prevDeviceId.current, device.deviceId);
		switchDevice(localTrack, device.deviceId);
		} else {
		console.log(`unmuting local ${kind} track`);
		localTrack?.unmute();
		}

		return () => {
		if (localTrack) {
			console.log(`stopping local ${kind} track`);
			localTrack.stop();
			localTrack.mute();
		}
		};
	}, [localTrack, device, enabled, kind, switchDevice]);

	React.useEffect(() => {
    setDevice(mediaDevice.find((dev) => dev.deviceId === localDeviceId));
  }, [localDeviceId, mediaDevice]);
  return {
	device,
	deviceError,
	localTrack,
  }

}

export default function PreJoin() {
	const [ 
		user,  
		toggleVideo,
		toggleAudio,
		setUserName,
		setLang,
		setUserChoices,
		setAudioDeviceId,
		setVideoDeviceId,
	] = userStore((user) => [
		user.user,
		user.toggleVideo,
		user.toggleAudio,
		user.setUserName,
		user.setLang,
		user.setNewUser,
		user.setAudioDeviceId,
		user.setVideoDeviceId,
	] )
	const onError = useCallback((err: Error) => {
		console.error('error while setting up rejoin', err);
	}, []);

	const onValidate = () => {
	}

	const audio = usePreviewDevice(user.audioEnabled, user.audioDeviceId, 'audioinput');
	const video = usePreviewDevice(user.videoEnabled, user.videoDeviceId, 'videoinput');
	const videoEl = React.useRef<HTMLVideoElement>(null);

	React.useEffect(() => {
		if (videoEl.current) {
			video.localTrack?.attach(videoEl.current);
		}

		return () => {
			video.localTrack?.detach();
		}
	}, [videoEl, video.localTrack]);

	  React.useEffect(() => {
    if (audio.deviceError) {
      onError(audio.deviceError);
    }
  }, [audio.deviceError, onError]);

  React.useEffect(() => {
    if (video.deviceError) {
      onError(video.deviceError);
    }
  }, [video.deviceError, onError]);


	return (

		<></>
	);
}