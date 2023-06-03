import { getRooms } from "@/lib/rooms";

Room.getStaticPaths = async () => {
	const rooms = await getRooms();
	const paths = rooms.map((room) => ({
		params: { room: room.id },
	}));

	return { paths, fallback: false };
}

interface RoomProps {
	room: string;
}
export default function Room({ room }: RoomProps) {

}