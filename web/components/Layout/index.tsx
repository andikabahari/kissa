import { useRouter } from 'next/router'
import { ReactNode } from 'react'
import Navbar from '../Navbar'

type LayoutProps = {
  children: ReactNode
}

export default function Layout({ children }: LayoutProps) {
  const router = useRouter()

  return (
    <>
      {router.pathname !== '/_error' ? (
        <div className='container max-w-screen-lg mx-auto'>
          <Navbar />
          {children}
        </div>
      ) : (
        <>{children}</>
      )}
    </>
  )
}
