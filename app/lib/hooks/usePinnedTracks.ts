// This is from [livekit-examples/kitt](https://github.com/livekit-examples/kitt/blob/main/meet/hooks/usePinnedTracks.ts)

import { TrackReference } from '@livekit/components-core';
import React from 'react';
import { LayoutContextType, useEnsureLayoutContext } from '@livekit/components-react';


const usePinnedTracks = (layoutCtx: LayoutContextType):TrackReference[] => {
	layoutCtx = useEnsureLayoutContext(layoutCtx);

	return React.useMemo(() => {

		throw new Error('Not implemented yet');
		return [];
	}, [layoutCtx]);	
}