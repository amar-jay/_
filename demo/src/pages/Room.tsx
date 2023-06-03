import React, { useEffect } from 'react'
import { ErrorPage } from "./ErrorPage";
import { Conference } from '../components/ui/Conference';
import { LayoutContextProvider } from '@livekit/components-react';
import { useNavigate } from 'react-router-dom';
import { MediaDeviceMenu, LiveKitRoom, AudioVisualizer, CarouselView } from '@livekit/components-react';

interface RoomProps {
	room_id?: string;
}
export function Room(_props: RoomProps) {
	const router = useNavigate()
	const room_id = new URL(window.location.href).searchParams.get("room_id"); // get room id from url
	if (!room_id) {
		router("/error?message=" + "Room id not found")
	}

	useEffect(() => {
		fetch("https://themanan.me" + room_id, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
			},
		}).then((res) => {
			if (res.status !== 200) {
				throw new Error(
					"Room not found",
				);
			}
		}).catch((e) => {
			// router("/error?message=" + e.message)
		})
		return () => {
			// cleanup
		}
	}, [room_id, router]);
	return (
		<div className="">
			<LiveKitRoom  
				serverUrl='wss://livekit.example.com'
				token={"xxx"}
				video={true}
				audio={true}
			>
			<LayoutContextProvider>
				Room: {room_id}
				<Conference />
	
			</LayoutContextProvider>
			</LiveKitRoom>
			<MediaDeviceMenu
				// kind="audioinput"
			/>

		</div>
	)

}
