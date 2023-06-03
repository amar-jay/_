import React from "react";
import { Html, Head, Main, NextScript } from "next/document";

import { Metadata } from "next"

import { siteConfig } from "@/config/site"
import { fontSans } from "@/lib/fonts"
import { cn } from "@/lib/utils"
import { SiteHeader } from "@/components/site-header"
import { TailwindIndicator } from "@/components/tailwind-indicator"
import { ThemeProvider } from "@/components/theme-provider"



interface RootLayoutProps {
  children: React.ReactNode
}

export default function Document() {
  return (
      <Html lang="en" suppressHydrationWarning>
	  <Head/>
        <body
          className={cn(
            "bg-background min-h-screen font-sans antialiased",
            fontSans.variable
          )}
        >
          <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
            {/* <div className="relative flex  min-h-screen flex-col bg-black"> */}
              <SiteHeader />
              {/* <div className="flex-1 bg-black"> */}
				<Main />
				{/* </div> */}
            {/* </div> */}
            <TailwindIndicator />
          </ThemeProvider>

		<NextScript />
	  </body>
	</Html>
  );
}