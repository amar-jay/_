import { NextResponse } from "next/server";

export default function Create() {
	NextResponse.redirect("/rooms/" + Math.random().toString(36).substring(7));
	return null;
}