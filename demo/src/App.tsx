"use client";

import { useEffect, useState } from "react";
import "./App.css";
import Index from "./pages/Index";
import { CopyLink } from "./pages/CopyLink";
import { ToastProvider } from "./components/toast";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { Demo } from "./pages/Demo";
import { ErrorPage } from "./pages/ErrorPage";
import { Room } from "./pages/Room";

// import pages
// TODO: import pages dynamically with TYPE ANNOTATIONS
// const routes = await import.meta.globEager("./pages/**/*.tsx").then((modules) => {
//   return Object.entries(modules).map(([path, page]: [path: string, page: any]) => ({
//     path: path.replace("./pages", "").replace(".tsx", ""),
//     element: page.default,
//   }))
// });

// import components



function App() {
  const [status, setStatus] = useState<
    "connected" | "not-connected" | string | undefined
  >();
  

  useEffect(() => {
    // connect(device).then(() => {
    //   console.log("Connected to server");
      setStatus("connected");
    // });
  }, []);
  const router = createBrowserRouter([
    {
      path: "/",
      element: (() => <Index status={status} />)(),

    },
    {
      path: "/copy-link",
      element: (() => <CopyLink status={status} />)(),
    },
    {
      path: "/demo",
      element: <Demo/>,
    },
    {
      path: "/room",
      element: (() => <Room />)(),
    },
    {
      path: "*",
      element: (() => <ErrorPage message={"Page not found"} />)(),
    },

    {
      path: "/error",
      element: (() => <ErrorPage />)(),
    }
  ]);



  return (
    <ToastProvider>
      <RouterProvider router={router} />
    </ToastProvider>
  )
}

export default App;
