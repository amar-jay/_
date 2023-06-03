import { useEffect, useState } from "react";

export const handleLivekitServerUrl = (region: string ): string => {
	let target_key = 'LIVEKIT_URL';
	if (!region) {
		throw Error('Region is required')
		// throw Error('Multiple regions not supported')
	}
	target_key = target_key + '_' + region.toUpperCase()
	const url = process.env[target_key]
	if (!url) {
		throw Error(`No server url found for region ${region}`)
	}

	return url;
}

export const useLivekitServerUrl = (region: string) => {
	const [ url, setUrl ] = useState<string>('')

	useEffect(() => {
		try {
			const url = handleLivekitServerUrl(region)
			setUrl(url)
		} catch (err) {
			console.error(err)
		}
	}, [])

	return url
}