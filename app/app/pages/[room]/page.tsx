import { env } from "process"

interface RoomParams {
	room: string
}
export default function Room({params}: {params: RoomParams}) {

	return (
    <section className="container grid items-center gap-6 pb-8 pt-6 md:py-10">
	<h1>
	Room: {params.room}
	{env.NODE_ENV || "None "}
	</h1>
		</section>
	)
}