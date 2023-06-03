import React from 'react'
interface ErrorPageProps {
	message?: string;
}
export function ErrorPage({message}: ErrorPageProps){
	if (!message) {
		// get message from url
		const url = new URL(window.location.href).searchParams.get("message");
		if (url) {
			message = url;
		} else {
			message = "Unknown error";
		}
	}

	return (
	<div>
		<h1>404</h1>
		<p>{message}</p>
	</div>
  );
}
