import Link from "next/link"

import { siteConfig } from "@/config/site"
import { Button, buttonVariants } from "@/components/ui/button"
import { redirect } from 'next/navigation'
import { Metadata } from "next/types"

export const metadata: Metadata = {
  title: {
    default: siteConfig.name,
    template: `%s - ${siteConfig.name}`,
  },
  description: siteConfig.description,
  themeColor: [
    { media: "(prefers-color-scheme: light)", color: "white" },
    { media: "(prefers-color-scheme: dark)", color: "black" },
  ],
  icons: {
    icon: "/favicon.ico",
    shortcut: "/favicon-16x16.png",
    apple: "/apple-touch-icon.png",
  },
}
export default function IndexPage() {
  redirect('/home')
  // return (
  //   <section className="container grid items-center gap-6 pb-8 pt-6 md:py-10">
  //     <div className="flex max-w-[980px] flex-col items-start gap-2">
  //       <h1 className="text-3xl font-extrabold leading-tight tracking-tighter md:text-4xl">
  //         Beautifully designed components <br className="hidden sm:inline" />
  //         built with Radix UI and Tailwind CSS.
  //       </h1>
  //       <p className="text-muted-foreground max-w-[700px] text-lg">
  //         Accessible and customizable components that you can copy and paste
  //         into your apps. Free. Open Source. And Next.js 13 Ready.
  //       </p>
  //     </div>
  //     <div className="flex gap-4">
  //       <Link
  //         href={'/pages'}
  //         target="_blank"
  //         rel="noreferrer"
  //         className={buttonVariants()}
  //       >
  //         Documentation
  //       </Link>
  //       <Button
  //       variant={'outline'}
  //         // onClick={() => redirect('/pages')}
  //         className={buttonVariants({ variant: "outline" })}
  //       >
  //         GitHub
  //       </Button>
  //     </div>
  //   </section>
  // )
}
