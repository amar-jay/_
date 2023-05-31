import { useRouter } from 'next/router'
import {NextRequest, NextResponse} from 'next/server'
import { generateRoomID } from '@/lib/generate';

export default function Home(request: NextRequest) {
  const url = "/rooms/" + generateRoomID()
  return NextResponse.rewrite(new URL('/rooms', request.url));
  
}
