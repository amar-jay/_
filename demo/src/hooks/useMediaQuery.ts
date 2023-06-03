/**
 * This is from 
 * [juliencrn/usehooks-ts](https://github.com/juliencrn/usehooks-ts)
 */

import { useCallback, useEffect, useState } from "react";

export const useMediaQuery = (query: string):boolean => {
	const getMatches = (query: string) => {
		// prevents SSR issues
		if (typeof window === 'undefined') return false;
		return window.matchMedia(query).matches;
	}

	const [matches, setMatches] = useState(getMatches(query));

	const handleMatchChange = useCallback(() => () => {
		setMatches(getMatches(query));
	}, [query]);

	useEffect(() => {
		const matchMedia = window.matchMedia(query);

		handleMatchChange();
		if (!matchMedia.addEventListener) {
			matchMedia.addEventListener('change', handleMatchChange);
		} else {
			matchMedia.addListener(handleMatchChange);
		}

		return () => {
			if (matchMedia.removeEventListener)
				matchMedia.removeEventListener('change', handleMatchChange);
			else
				matchMedia.removeListener(handleMatchChange);

		}
	}, [query, handleMatchChange]);

	return matches
}