import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import '@livekit/components-styles';
// import '@livekit/components-styles/themes/huddle';
// import '@livekit/components-styles/prefabs';

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
