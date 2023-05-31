"use client"
import ActiveRoom from '@/components/ActiveRoom';
import React from 'react'

interface RoomsProps {
  params: {
  room:  `room-${string}`
  }
}
function Rooms({params}: RoomsProps) {
  const roomId = params.room.split("-").slice(1,).join();
  return (
    <main>
      <h1 className='text-blue-500 text-5xl underline'> Hey </h1>
      {
        roomId && !Array.isArray(roomId) &&  (
          <ActiveRoom
            roomId={roomId}
            onExit={()=> {}}
            />
        )
      }
    </main>
	)
}

export default Rooms