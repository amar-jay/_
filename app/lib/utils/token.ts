import { AccessTokenOptions, VideoGrant, AccessToken } from 'livekit-server-sdk';
import { env } from '@/env.mjs';
import { useEffect, useState } from 'react';
import { Token } from '@/types/livekit';

const ROOM_PATTERN = /\w{4}-\w{4}/;

const createToken = (tokenOpts: AccessTokenOptions, grant: VideoGrant): string => {
	const token = new AccessToken(env.LIVEKIT_API_KEY, env.LIVEKIT_API_SECRET, tokenOpts);
	token.ttl = 60 * 60 * 24 * 7; // 1 week
	token.addGrant(grant);
	return token.toJwt();
}

interface handleTokenOptions {
	identity: string;
	room: string;
	name: string;
	metadata?: string;
}

/**  
 * to create a LiveKit token for a user. 
 * #### Note: this is a server function
*/
export const handleToken = ({ identity, room, name, metadata}: handleTokenOptions): Token => {
	if (!identity || typeof identity !== 'string') {
		throw Error('Identity is required')
	}
	if (!room || typeof room !== 'string') {
		throw Error('Room is required')
	}
	if (!name || typeof name !== 'string') {
		throw Error('Name is required')
	}

	const token = createToken({
		identity,
		name,
		metadata,
	}, {
		room: room,
		roomJoin: true,
		canPublish: true,
		canPublishData: true,
		canSubscribe: true,
		canUpdateOwnMetadata: true, // optional
	});


	return {
		token,
		identity
	}

}

export const useToken = ({ identity, room, name, metadata}: handleTokenOptions):Token => {
	const [ token, setToken ] = useState<Token>()
	useEffect(() => {
		try {
		const token = handleToken({
			identity,
			room,
			name
		})
		console.log(token)

	} catch (err) {
		console.log(err)
	}
	}, [identity, room, name])

	if (!token) {
		console.error('No token found')
		return {
			token: '',
			identity: ''
		}
	}

	return token
}
