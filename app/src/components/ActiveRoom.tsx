import React, { useCallback } from 'react'
import { Button } from './ui/button';

interface ActiveRoomProps {
	onExit : () => void,
	roomId: string,
}
function ActiveRoom ({onExit}: ActiveRoomProps) {

	const exitRoom = useCallback(() => onExit(), [onExit]);
	return (
		<div className="text-5xl text-slate-400">
			ActiveRoom
			<Button onClick={exitRoom}/>
		</div>
	)
}

export default ActiveRoom