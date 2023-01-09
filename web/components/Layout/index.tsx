import { useRouter } from 'next/router'
import { ReactNode } from 'react'
import Breadcrumb from '../Breadcrumb'
import Navbar from '../Navbar'

type LayoutProps = {
  children: ReactNode
}

export default function Layout({ children }: LayoutProps) {
  const router = useRouter()

  if (router.pathname === '/_error') {
    return <>{children}</>
  }

  return (
    <>
      <Navbar />
      <div className='container max-w-screen-lg mx-auto px-6 mb-8'>
        {router.pathname !== '/' && <Breadcrumb />}
        {children}
      </div>
    </>
  )
}
